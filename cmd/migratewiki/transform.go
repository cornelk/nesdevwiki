package main

import (
	"bytes"
	"fmt"
	"net/url"
	"path"
	"regexp"
	"strings"
	"sync"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

const imagesURL = "/wiki-images/"

// anchorable lists the elements that can hold an empty anchor child without breaking the HTML structure.
var anchorable = map[string]bool{
	"p": true, "div": true, "span": true, "sup": true, "sub": true, "dt": true, "dd": true,
	"b": true, "i": true, "code": true, "tt": true, "li": true, "td": true, "th": true,
}

var (
	// urlScheme matches absolute URLs like https:// or mailto:. Colons inside fragments are not schemes.
	urlScheme    = regexp.MustCompile(`^[A-Za-z][A-Za-z0-9+.-]*:`)
	liveWikiLink = regexp.MustCompile(`^https?://(?:www\.)?nesdev\.org/wiki/([^#?]*)(#.*)?$`)
	tagGap       = regexp.MustCompile(`>\s+<`)
)

// site holds everything needed to resolve links between pages and to images.
type site struct {
	pages      map[string]*page // by dump file name
	byCanon    map[string]*page // by lower-case canonical name
	images     map[string]string
	fileAssets map[string]string // dump file name of a File: page -> image name

	mu             sync.Mutex
	missingLinks   map[string]int
	missingImages  map[string]int
	externalLinks  int
	droppedLinks   int
	rewrittenLinks int
}

// transformer rewrites the article of a single page.
type transformer struct {
	site     *site
	page     *page
	infobox  []string
	headings int
}

// transform cleans the article DOM and returns its HTML together with the infobox blocks that must
// bypass Markdown conversion.
func (s *site) transform(p *page) (string, []string) {
	t := &transformer{site: s, page: p}
	root := p.Article

	t.strip(root)
	t.tableHeaders(root)
	t.headingsAndAnchors(root)
	t.rewriteLinks(root)
	t.rewriteImages(root)
	t.extractInfoboxes(root)

	var buf bytes.Buffer
	for c := root.FirstChild; c != nil; c = c.NextSibling {
		_ = html.Render(&buf, c)
	}
	return buf.String(), t.infobox
}

// strip removes comments, scripts, the table of contents (Hugo builds its own), edit links, and the
// generated member lists of category pages.
func (t *transformer) strip(root *html.Node) {
	for n := range descendants(root) {
		switch {
		case n.Type == html.CommentNode:
			remove(n)
		case n.Type != html.ElementNode:
		case n.Data == "script" || n.Data == "style":
			remove(n)
		case n.Data == "div" && (attr(n, "id") == "toc" || hasClass(n, "toc")):
			remove(n)
		case hasClass(n, "mw-category-generated"):
			// Static member lists of category pages, Hugo generates them from the page front matter.
			remove(n)
		case hasClass(n, "mw-editsection") || hasClass(n, "printfooter") || hasClass(n, "noprint"):
			remove(n)
		}
	}
}

// tableHeaders moves a leading row of header cells into a thead. MediaWiki leaves it in the tbody, so pandoc
// would turn it into a data row below an empty header row.
func (t *transformer) tableHeaders(root *html.Node) {
	for _, table := range findAll(root, func(n *html.Node) bool { return isElement(n, "table") }) {
		var body *html.Node
		hasHead := false
		for c := table.FirstChild; c != nil; c = c.NextSibling {
			switch {
			case isElement(c, "thead"):
				hasHead = true
			case body == nil && isElement(c, "tbody"):
				body = c
			}
		}
		if hasHead || body == nil {
			continue
		}

		var row *html.Node
		for c := body.FirstChild; c != nil && row == nil; c = c.NextSibling {
			if isElement(c, "tr") {
				row = c
			}
		}
		if row == nil || !onlyHeaderCells(row) {
			continue
		}

		head := &html.Node{Type: html.ElementNode, DataAtom: atom.Thead, Data: "thead"}
		body.RemoveChild(row)
		head.AppendChild(row)
		table.InsertBefore(head, body)
	}
}

// onlyHeaderCells reports whether a table row consists of th cells only.
func onlyHeaderCells(row *html.Node) bool {
	cells := 0
	for c := row.FirstChild; c != nil; c = c.NextSibling {
		if c.Type != html.ElementNode {
			continue
		}
		if !isElement(c, "th") {
			return false
		}
		cells++
	}
	return cells > 0
}

// headingsAndAnchors keeps MediaWiki anchor ids intact: headings get an explicit {#id} attribute and
// other elements with an id get an empty anchor, so existing #fragment links keep working.
func (t *transformer) headingsAndAnchors(root *html.Node) {
	for _, n := range findAll(root, func(n *html.Node) bool { return n.Type == html.ElementNode }) {
		switch {
		case n.Data == "div" && hasClass(n, "mw-heading"):
			unwrap(n)
		case isHeading(n):
			t.heading(n)
		case n.Data == "span" && hasClass(n, "mw-headline"):
			unwrap(n)
		case attr(n, "id") != "" && anchorable[n.Data]:
			anchor := &html.Node{Type: html.ElementNode, DataAtom: atom.A, Data: "a",
				Attr: []html.Attribute{{Key: "id", Val: anchorID(attr(n, "id"))}}}
			prependChild(n, anchor)
			removeAttr(n, "id")
		}
	}
}

// anchorID makes an anchor id usable in a Markdown {#id} attribute, which only accepts letters, digits and
// _ . : -. Other bytes are escaped the way MediaWiki does it (".28" for "("). Applying the same function
// to ids and to link fragments keeps both sides identical, whichever form the export used.
func anchorID(id string) string {
	var b strings.Builder
	for i := 0; i < len(id); i++ {
		c := id[i]
		switch {
		case c >= 'a' && c <= 'z', c >= 'A' && c <= 'Z', c >= '0' && c <= '9', c == '_', c == '.', c == ':', c == '-':
			b.WriteByte(c)
		default:
			fmt.Fprintf(&b, ".%02X", c)
		}
	}
	return b.String()
}

// fragmentID converts the fragment of a link, which may be percent-encoded, into the id of its anchor.
func fragmentID(fragment string) string {
	if unescaped, err := url.PathUnescape(fragment); err == nil {
		fragment = unescaped
	}
	return anchorID(fragment)
}

func isHeading(n *html.Node) bool {
	return len(n.Data) == 2 && n.Data[0] == 'h' && n.Data[1] >= '1' && n.Data[1] <= '6'
}

func (t *transformer) heading(n *html.Node) {
	id := attr(n, "id")
	if id == "" {
		if span := find(n, func(c *html.Node) bool { return isElement(c, "span") && attr(c, "id") != "" }); span != nil {
			id = attr(span, "id")
		}
	}
	if id == "" {
		return
	}
	n.AppendChild(&html.Node{Type: html.TextNode, Data: " {#" + anchorID(id) + "}"})
	removeAttr(n, "id")
	t.headings++
}

func (t *transformer) rewriteLinks(root *html.Node) {
	for _, a := range findAll(root, func(n *html.Node) bool { return isElement(n, "a") }) {
		href := attr(a, "href")
		if href == "" {
			if attr(a, "id") == "" {
				unwrap(a)
			}
			continue
		}
		dest, keep := t.resolveLink(href)
		if !keep {
			unwrap(a)
			continue
		}
		setAttr(a, "href", dest)
		for _, key := range []string{"class", "title", "rel", "style"} {
			removeAttr(a, key)
		}
	}
}

// resolveLink returns the new destination of an href, or false when the link has to be dropped.
func (t *transformer) resolveLink(href string) (string, bool) {
	if fragment, ok := strings.CutPrefix(href, "#"); ok {
		if fragment == "" {
			return href, true
		}
		return "#" + fragmentID(fragment), true
	}
	if m := liveWikiLink.FindStringSubmatch(href); m != nil {
		name, _ := url.PathUnescape(m[1])
		if p := t.site.byCanon[strings.ToLower(name)]; p != nil {
			return t.pageLink(p, strings.TrimPrefix(m[2], "#"))
		}
		return href, true
	}
	if urlScheme.MatchString(href) || strings.HasPrefix(href, "//") {
		t.site.count(&t.site.externalLinks)
		return href, true
	}

	target, fragment, _ := strings.Cut(href, "#")
	if strings.HasPrefix(target, "../wiki-images/") {
		name, _ := url.PathUnescape(strings.TrimPrefix(target, "../wiki-images/"))
		return t.assetLink(name)
	}
	if !strings.HasSuffix(target, ".xhtml") {
		return href, true
	}

	file, err := url.PathUnescape(strings.TrimSuffix(target, ".xhtml"))
	if err != nil {
		file = strings.TrimSuffix(target, ".xhtml")
	}
	if name, ok := t.site.fileAssets[file]; ok {
		return t.assetLink(name)
	}
	target0, frag, ok := t.site.resolve(file, fragment)
	if !ok {
		if p := t.site.pages[file]; p != nil && p.Kind == kindDropped {
			t.site.count(&t.site.droppedLinks)
		} else {
			t.site.miss(t.site.missingLinks, file)
		}
		return "", false
	}
	return t.pageLink(target0, frag)
}

func (t *transformer) pageLink(target *page, fragment string) (string, bool) {
	t.site.count(&t.site.rewrittenLinks)
	dest := relURL(t.page.URL, target.URL)
	if fragment != "" {
		dest += "#" + fragmentID(fragment)
	}
	return dest, true
}

func (t *transformer) assetLink(name string) (string, bool) {
	actual, ok := t.site.images[strings.ToLower(name)]
	if !ok {
		t.site.miss(t.site.missingImages, name)
		return "", false
	}
	return relURL(t.page.URL, imagesURL+actual), true
}

func (t *transformer) rewriteImages(root *html.Node) {
	for _, img := range findAll(root, func(n *html.Node) bool { return isElement(n, "img") }) {
		name := imageName(attr(img, "src"))
		dest, ok := t.assetLink(name)
		if !ok || name == "" {
			if p := img.Parent; p != nil && isElement(p, "a") {
				unwrap(p)
			}
			remove(img)
			continue
		}
		setAttr(img, "src", dest)
		for _, key := range []string{"srcset", "decoding", "class", "style", "data-file-width", "data-file-height"} {
			removeAttr(img, key)
		}
	}
}

// imageName extracts the file name from the image paths used by the export: ../wiki-images/Name.png
// and MediaWiki upload URLs such as /w/images/default/thumb/c/ce/Name.png/32px-Name.png.
func imageName(src string) string {
	const marker = "/images/default/"
	switch {
	case strings.HasPrefix(src, "../wiki-images/"):
		name, _ := url.PathUnescape(strings.TrimPrefix(src, "../wiki-images/"))
		return name
	case strings.Contains(src, marker):
		_, rest, _ := strings.Cut(src, marker)
		parts := strings.Split(rest, "/")
		index := 2
		if parts[0] == "thumb" {
			index = 3
		}
		if index < len(parts) {
			name, _ := url.PathUnescape(parts[index])
			return name
		}
	}
	return ""
}

// extractInfoboxes moves infobox blocks out of the DOM. Markdown cannot express their layout, so they
// are re-inserted as raw HTML after conversion.
func (t *transformer) extractInfoboxes(root *html.Node) {
	for _, box := range findAll(root, func(n *html.Node) bool { return isElement(n, "div") && hasClass(n, "infobox") }) {
		var buf bytes.Buffer
		_ = html.Render(&buf, box)
		compact := tagGap.ReplaceAllString(strings.Join(strings.Fields(buf.String()), " "), "><")

		token := &html.Node{Type: html.ElementNode, Data: "p", DataAtom: atom.P}
		token.AppendChild(&html.Node{Type: html.TextNode, Data: infoboxToken(len(t.infobox))})
		box.Parent.InsertBefore(token, box)
		remove(box)
		t.infobox = append(t.infobox, compact)
	}
}

func infoboxToken(i int) string {
	return fmt.Sprintf("ZZINFOBOX%dZZ", i)
}

// resolve follows redirect stubs to the page holding the content.
func (s *site) resolve(file, fragment string) (*page, string, bool) {
	for range 8 {
		p := s.pages[file]
		if p == nil || p.Kind == kindDropped {
			return nil, "", false
		}
		if p.Redirect == nil {
			return p, fragment, true
		}
		if fragment == "" {
			fragment = p.Redirect.Fragment
		}
		file = p.Redirect.File
	}
	return nil, "", false
}

func (s *site) miss(m map[string]int, key string) {
	s.mu.Lock()
	m[key]++
	s.mu.Unlock()
}

func (s *site) count(n *int) {
	s.mu.Lock()
	*n++
	s.mu.Unlock()
}

// relURL builds a relative link from one site URL to another, so the site works under any base path.
func relURL(from, to string) string {
	fromDir := splitPath(from)
	toDir := splitPath(to)
	isFile := !strings.HasSuffix(to, "/")

	limit := len(toDir)
	if isFile {
		limit--
	}
	common := 0
	for common < len(fromDir) && common < limit && fromDir[common] == toDir[common] {
		common++
	}

	rest := make([]string, 0, len(toDir)-common)
	for _, seg := range toDir[common:] {
		rest = append(rest, url.PathEscape(seg))
	}
	out := strings.Repeat("../", len(fromDir)-common) + strings.Join(rest, "/")
	if len(rest) > 0 && !isFile {
		out += "/"
	}
	if out == "" {
		return "./"
	}
	if len(fromDir) == common && strings.Contains(strings.SplitN(out, "/", 2)[0], ":") {
		out = "./" + out
	}
	return out
}

func splitPath(p string) []string {
	p = strings.Trim(path.Clean("/"+p), "/")
	if p == "" {
		return nil
	}
	return strings.Split(p, "/")
}
