package main

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

// kind classifies a wiki page by namespace.
type kind int

const (
	kindMain kind = iota
	kindCategory
	kindDropped
)

// page is one exported wiki page.
type page struct {
	File       string // dump file name without extension
	Title      string
	Canonical  string // decoded page name as used in the original URL
	Kind       kind
	URL        string // site-absolute URL path with leading and trailing slash
	Article    *html.Node
	Categories []categoryRef
	Redirect   *redirect // set for redirect stubs

	// Aliases are the original URLs of redirect stubs pointing at this page.
	Aliases []string
}

// categoryRef is a category a page belongs to.
type categoryRef struct {
	File string
	Name string
}

// redirect is the target of a redirect stub.
type redirect struct {
	File     string
	Fragment string
}

// namespaces maps title prefixes to the kind of page they hold. Prefixes missing here are part of
// the page name (for example "Visual6502wiki:" pages live in the main namespace).
var namespaces = map[string]kind{
	"Category":         kindCategory,
	"User":             kindDropped,
	"User talk":        kindDropped,
	"Talk":             kindDropped,
	"Category talk":    kindDropped,
	"Help talk":        kindDropped,
	"Nesdev wiki talk": kindDropped,
	"File":             kindDropped,
	"File talk":        kindDropped,
	"Template":         kindDropped,
	"Template talk":    kindDropped,
	"MediaWiki":        kindDropped,
	"MediaWiki talk":   kindDropped,
	"Special":          kindDropped,
}

// homeFile is the portal page that becomes the site's home page. The wiki's Main_Page only holds the
// default MediaWiki installation text.
const homeFile = "Nesdev_Wiki"

// projectPages are the pages of the "Nesdev wiki:" namespace worth keeping. The others describe MediaWiki
// permission groups and spam filters that do not exist in the static wiki.
var projectPages = map[string]struct{}{
	"Nesdev_wiki_About":                    {},
	"Nesdev_wiki_General_disclaimer":       {},
	"Nesdev_wiki_Privacy_policy":           {},
	"Nesdev_wiki_Manual_of_Style_RFC_2119": {},
}

var (
	spaces       = regexp.MustCompile(`\s+`)
	unsafePathCh = regexp.MustCompile(`[^\pL\pN_.,'()!~*+=@$&/-]`)
)

// loadPage parses one exported XHTML page.
func loadPage(dir, file string) (*page, error) {
	f, err := os.Open(filepath.Join(dir, file+".xhtml"))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	doc, err := html.Parse(f)
	if err != nil {
		return nil, fmt.Errorf("parsing %s: %w", file, err)
	}

	p := &page{File: file}
	if h1 := find(doc, func(n *html.Node) bool { return isElement(n, "h1") }); h1 != nil {
		p.Title = strings.TrimSpace(text(h1))
	}
	if link := find(doc, func(n *html.Node) bool {
		return isElement(n, "link") && attr(n, "rel") == "canonical"
	}); link != nil {
		p.Canonical = canonicalName(attr(link, "href"))
	}
	if p.Title == "" || p.Canonical == "" {
		return nil, fmt.Errorf("%s has no title or canonical link", file)
	}

	p.Article = find(doc, func(n *html.Node) bool { return isElement(n, "div") && hasClass(n, "article") })
	if p.Article == nil {
		return nil, fmt.Errorf("%s has no article", file)
	}

	p.Kind = kindOfTitle(p.Title)
	// Main_Page only holds the default MediaWiki text, Special_AllPages is a static page list that the site
	// generates itself (its title has no namespace prefix, so the namespace check does not catch it).
	if file == "Main_Page" || file == "Special_AllPages" {
		p.Kind = kindDropped
	}
	if strings.HasPrefix(p.Title, "Nesdev wiki:") {
		if _, keep := projectPages[file]; !keep {
			p.Kind = kindDropped
		}
	}
	p.Categories = extractCategories(p.Article)
	p.URL = pageURL(p)
	return p, nil
}

// canonicalName extracts the page name from a canonical URL like https://www.nesdev.org/wiki/Foo%3ABar.
func canonicalName(href string) string {
	_, name, ok := strings.Cut(href, "/wiki/")
	if !ok {
		return ""
	}
	decoded, err := url.PathUnescape(name)
	if err != nil {
		return name
	}
	return decoded
}

func kindOfTitle(title string) kind {
	prefix, _, ok := strings.Cut(title, ":")
	if !ok {
		return kindMain
	}
	if k, found := namespaces[prefix]; found {
		return k
	}
	return kindMain
}

// pageURL builds the site URL of a page, keeping the original /wiki/<name>/ layout.
func pageURL(p *page) string {
	if p.File == homeFile {
		return "/"
	}
	return wikiURL(p.Canonical)
}

// wikiURL turns a page name into a site URL path that is safe for files and links. Colons are not kept:
// Hugo cannot build pages with them in their URL and Cloudflare _redirects treats them as placeholders.
// Categories live below /category/, other namespaces keep their prefix joined by an underscore.
func wikiURL(name string) string {
	name = strings.ReplaceAll(name, " ", "_")
	if cat, ok := strings.CutPrefix(name, "Category:"); ok {
		return "/category/" + unsafePathCh.ReplaceAllString(cat, "_") + "/"
	}
	return "/wiki/" + unsafePathCh.ReplaceAllString(name, "_") + "/"
}

// extractCategories removes the trailing category list from the article and returns its entries.
func extractCategories(article *html.Node) []categoryRef {
	node := find(article, func(n *html.Node) bool { return isElement(n, "p") && hasClass(n, "categories") })
	if node == nil {
		return nil
	}

	var refs []categoryRef
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if !isElement(c, "a") {
			continue
		}
		refs = append(refs, categoryRef{
			File: strings.TrimSuffix(attr(c, "href"), ".xhtml"),
			Name: strings.TrimSpace(text(c)),
		})
	}
	node.Parent.RemoveChild(node)
	return refs
}

// redirectOf detects redirect stubs: pages consisting only of "See <link>.".
func redirectOf(p *page) *redirect {
	var links []*html.Node
	for n := range descendants(p.Article) {
		if n.Type != html.ElementNode {
			continue
		}
		switch n.Data {
		case "a":
			links = append(links, n)
		case "p", "div":
		default:
			return nil
		}
	}
	if len(links) != 1 {
		return nil
	}

	body := normalizeSpace(text(p.Article))
	label := normalizeSpace(text(links[0]))
	if body != label && body != "See "+label+"." && body != "See "+label {
		return nil
	}

	target, fragment, _ := strings.Cut(attr(links[0], "href"), "#")
	target = strings.TrimSuffix(target, ".xhtml")
	if target == "" || strings.Contains(target, "/") {
		return nil
	}
	unescaped, err := url.PathUnescape(target)
	if err != nil {
		unescaped = target
	}
	return &redirect{File: unescaped, Fragment: fragment}
}

func normalizeSpace(s string) string {
	return strings.TrimSpace(spaces.ReplaceAllString(s, " "))
}

const descriptionLimit = 200

// description returns the first prose paragraph of the article, shortened for meta tags and search.
func description(p *page) string {
	for n := range descendants(p.Article) {
		if !isElement(n, "p") || !insideWrapper(n) {
			continue
		}
		s := normalizeSpace(text(n))
		if len(s) < 40 {
			continue
		}
		if len(s) > descriptionLimit {
			s = s[:descriptionLimit]
			if i := strings.LastIndexByte(s, ' '); i > 0 {
				s = s[:i]
			}
			s += "…"
		}
		return s
	}
	return ""
}

// insideWrapper reports whether the only block ancestors of n are MediaWiki content wrappers.
func insideWrapper(n *html.Node) bool {
	for a := n.Parent; a != nil; a = a.Parent {
		if isElement(a, "div") && !hasClass(a, "mw-parser-output") && !hasClass(a, "mw-content-ltr") && !hasClass(a, "article") {
			return false
		}
		if isElement(a, "table") || isElement(a, "ul") || isElement(a, "ol") {
			return false
		}
	}
	return true
}
