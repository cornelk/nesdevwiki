// Command migratewiki converts an XHTML export of the NESdev wiki into Hugo content.
//
// Pages of the main, Help, Category and selected project namespaces become Markdown files below
// content/. Redirect stubs become Cloudflare Pages redirects. Images are grouped below assets/.
// Talk, User, File, Template and MediaWiki pages are not migrated.
//
// This was a one-time import: afterwards the generated Markdown is the source of truth and is edited
// directly. To protect those edits the command refuses to run when the target already holds content.
// Delete content/wiki, content/category, content/_index.md and assets/wiki-images to import again.
//
// Usage:
//
//	go run ./cmd/migratewiki -src /path/to/extracted/dump -out .
//
// The source directory contains the wikipages and wiki-images folders.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func main() {
	src := flag.String("src", "", "directory containing the wikipages and wiki-images folders")
	out := flag.String("out", ".", "root of the Hugo site")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	if *src == "" {
		logger.Error("Missing -src")
		os.Exit(2)
	}
	if err := run(logger, *src, *out); err != nil {
		logger.Error("Migration failed", "error", err)
		os.Exit(1)
	}
}

func run(logger *slog.Logger, src, out string) error {
	for _, generated := range []string{"content/wiki", "content/category", "content/_index.md", "assets/wiki-images"} {
		path := filepath.Join(out, filepath.FromSlash(generated))
		if _, err := os.Stat(path); err == nil {
			return fmt.Errorf("%s already exists, refusing to overwrite imported content", path)
		}
	}

	s, err := loadSite(logger, src)
	if err != nil {
		return err
	}

	emitted, err := s.convertAll(logger, out)
	if err != nil {
		return err
	}
	if err := s.writeCategoryStubs(out, emitted); err != nil {
		return err
	}
	if err := s.writeRedirects(out); err != nil {
		return err
	}
	if err := copyImages(filepath.Join(src, "wiki-images"), filepath.Join(out, "assets", "wiki-images")); err != nil {
		return err
	}

	s.report(logger, emitted)
	return nil
}

// loadSite parses every page and builds the lookup tables for link resolution.
func loadSite(logger *slog.Logger, src string) (*site, error) {
	entries, err := os.ReadDir(filepath.Join(src, "wikipages"))
	if err != nil {
		return nil, err
	}

	s := &site{
		pages:         make(map[string]*page),
		byCanon:       make(map[string]*page),
		images:        make(map[string]string),
		fileAssets:    make(map[string]string),
		missingLinks:  make(map[string]int),
		missingImages: make(map[string]int),
	}

	images, err := os.ReadDir(filepath.Join(src, "wiki-images"))
	if err != nil {
		return nil, err
	}
	for _, e := range images {
		s.images[strings.ToLower(e.Name())] = e.Name()
	}

	for _, e := range entries {
		file, ok := strings.CutSuffix(e.Name(), ".xhtml")
		if !ok {
			continue
		}
		p, err := loadPage(filepath.Join(src, "wikipages"), file)
		if err != nil {
			return nil, err
		}
		s.pages[file] = p
		if name, isFile := strings.CutPrefix(p.Canonical, "File:"); isFile {
			s.fileAssets[file] = name
		}
	}

	for _, p := range s.pages {
		if p.Kind != kindDropped {
			p.Redirect = redirectOf(p)
		}
	}
	for _, p := range s.pages {
		if p.Kind == kindDropped {
			continue
		}
		s.byCanon[strings.ToLower(p.Canonical)] = p
	}

	// Stubs whose target does not exist are not pages and have no content to keep.
	for _, p := range s.pages {
		if p.Redirect == nil {
			continue
		}
		target, _, ok := s.resolve(p.Redirect.File, "")
		if !ok {
			logger.Warn("Redirect target missing", "page", p.File, "target", p.Redirect.File)
			p.Kind = kindDropped
			p.Redirect = nil
			continue
		}
		target.Aliases = append(target.Aliases, p.URL)
	}
	return s, nil
}

// convertAll converts all content pages in parallel and returns the names of the written files.
func (s *site) convertAll(logger *slog.Logger, out string) (map[string]*page, error) {
	var todo []*page
	for _, p := range s.pages {
		if p.Kind != kindDropped && p.Redirect == nil {
			todo = append(todo, p)
		}
	}
	sort.Slice(todo, func(i, j int) bool { return todo[i].File < todo[j].File })

	if err := checkCollisions(todo); err != nil {
		return nil, err
	}

	var (
		wg    sync.WaitGroup
		mu    sync.Mutex
		first error
		jobs  = make(chan *page)
	)
	for range runtime.NumCPU() {
		wg.Go(func() {
			for p := range jobs {
				if err := s.convert(out, p); err != nil {
					mu.Lock()
					if first == nil {
						first = fmt.Errorf("converting %s: %w", p.File, err)
					}
					mu.Unlock()
				}
			}
		})
	}
	for _, p := range todo {
		jobs <- p
	}
	close(jobs)
	wg.Wait()

	emitted := make(map[string]*page, len(todo))
	for _, p := range todo {
		emitted[p.File] = p
	}
	return emitted, first
}

// checkCollisions makes sure no two pages share a URL or a file name on case-insensitive file systems.
func checkCollisions(pages []*page) error {
	urls := make(map[string]string)
	files := make(map[string]string)
	for _, p := range pages {
		if other, ok := urls[p.URL]; ok {
			return fmt.Errorf("pages %s and %s share URL %s", other, p.File, p.URL)
		}
		urls[p.URL] = p.File
		lower := strings.ToLower(p.File)
		if other, ok := files[lower]; ok {
			return fmt.Errorf("pages %s and %s differ only in case", other, p.File)
		}
		files[lower] = p.File
	}
	return nil
}

func (s *site) convert(out string, p *page) error {
	desc := description(p)
	body, infoboxes := s.transform(p)
	markdown, err := toMarkdown(body)
	if err != nil {
		return err
	}
	for i, box := range infoboxes {
		markdown = strings.Replace(markdown, infoboxToken(i), box, 1)
	}
	if strings.Contains(markdown, "{{<") || strings.Contains(markdown, "{{%") {
		markdown = strings.NewReplacer("{{<", "{ {<", "{{%", "{ {%").Replace(markdown)
	}

	front := s.frontMatter(p, desc)
	dir, name := contentPath(p)
	full := filepath.Join(out, "content", dir)
	if err := os.MkdirAll(full, 0o755); err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(full, name), []byte(front+"\n"+markdown), 0o644)
}

// contentPath returns the content directory and file name of a page.
func contentPath(p *page) (string, string) {
	switch {
	case p.File == homeFile:
		return "", "_index.md"
	case p.Kind == kindCategory:
		return "category", p.File + ".md"
	default:
		return filepath.Join("wiki", wikiBucket(p.File)), p.File + ".md"
	}
}

func wikiBucket(name string) string {
	name = strings.TrimLeft(name, "_")
	for _, prefix := range []string{"NES_2_0_Mapper_", "INES_Mapper_"} {
		suffix := strings.TrimPrefix(name, prefix)
		if suffix == name {
			continue
		}
		number, err := strconv.Atoi(suffix)
		if err != nil {
			continue
		}
		start := number / 100 * 100
		return filepath.Join(prefix[:len(prefix)-1], fmt.Sprintf("%03d-%03d", start, start+99))
	}
	return sourceBucket(name)
}

func sourceBucket(name string) string {
	if name == "" {
		return "Other"
	}
	first := name[0]
	switch {
	case first >= '0' && first <= '9':
		return "0-9"
	case first >= 'a' && first <= 'z':
		return string(first - ('a' - 'A'))
	case first >= 'A' && first <= 'Z':
		return string(first)
	default:
		return "Other"
	}
}

func (s *site) frontMatter(p *page, desc string) string {
	var b strings.Builder
	b.WriteString("---\n")
	title := p.Title
	if p.Kind == kindCategory {
		title = strings.TrimPrefix(title, "Category:")
	}
	b.WriteString("title: " + quote(title) + "\n")
	b.WriteString("url: " + quote(p.URL) + "\n")
	if desc != "" {
		b.WriteString("description: " + quote(desc) + "\n")
	}
	b.WriteString("wikiPage: " + quote(p.File) + "\n")
	if p.Kind == kindCategory {
		b.WriteString("type: category\n")
	}
	if p.File == homeFile {
		b.WriteString("type: home\n")
	}
	if len(p.Categories) > 0 {
		b.WriteString("wikiCategories:\n")
		for _, c := range p.Categories {
			b.WriteString("  - name: " + quote(c.Name) + "\n")
			b.WriteString("    url: " + quote(s.categoryURL(c)) + "\n")
		}
	}
	b.WriteString("---\n")
	return b.String()
}

func (s *site) categoryURL(c categoryRef) string {
	if p := s.pages[c.File]; p != nil && p.Kind == kindCategory {
		return p.URL
	}
	return wikiURL("Category:" + c.Name)
}

// quote renders a YAML double-quoted scalar; JSON strings are valid YAML.
func quote(s string) string {
	b, _ := json.Marshal(s)
	return string(b)
}

// toMarkdown converts article HTML with pandoc.
func toMarkdown(htmlText string) (string, error) {
	cmd := exec.Command("pandoc", "--from=html-native_divs-native_spans", "--to=gfm", "--wrap=none")
	cmd.Stdin = strings.NewReader(htmlText)
	var stderr strings.Builder
	cmd.Stderr = &stderr
	stdout, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("running pandoc: %w: %s", err, stderr.String())
	}
	return tidy(string(stdout)), nil
}

// headingAttr matches the explicit {#id} attribute at the end of an ATX heading.
var headingAttr = regexp.MustCompile(`(?m)^(#{1,6} .*) \{#([^}\n]*)\}$`)

// unescapeHeadingIDs removes the backslashes pandoc adds in front of underscores, which break the attribute.
func unescapeHeadingIDs(markdown string) string {
	return headingAttr.ReplaceAllStringFunc(markdown, func(line string) string {
		i := strings.LastIndex(line, " {#")
		return line[:i] + strings.ReplaceAll(line[i:], `\_`, "_")
	})
}

// tidy collapses runs of blank lines and trims the result.
func tidy(markdown string) string {
	markdown = unescapeHeadingIDs(markdown)
	lines := strings.Split(markdown, "\n")
	out := make([]string, 0, len(lines))
	blank := 0
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			blank++
			if blank > 1 {
				continue
			}
			line = ""
		} else {
			blank = 0
		}
		out = append(out, line)
	}
	return strings.TrimSpace(strings.Join(out, "\n")) + "\n"
}

// writeCategoryStubs creates pages for categories that are used but have no page in the export.
func (s *site) writeCategoryStubs(out string, emitted map[string]*page) error {
	stubs := make(map[string]categoryRef)
	for _, p := range emitted {
		for _, c := range p.Categories {
			if e := emitted[c.File]; e == nil || e.Kind != kindCategory {
				stubs[c.File] = c
			}
		}
	}

	dir := filepath.Join(out, "content", "category")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	for file, c := range stubs {
		front := "---\ntitle: " + quote(c.Name) + "\nurl: " + quote(wikiURL("Category:"+c.Name)) +
			"\ntype: category\nwikiPage: " + quote(file) + "\n---\n"
		if err := os.WriteFile(filepath.Join(dir, file+".md"), []byte(front), 0o644); err != nil {
			return err
		}
	}
	return nil
}

// writeRedirects emits the Cloudflare Pages _redirects file for all redirect stubs.
func (s *site) writeRedirects(out string) error {
	var lines []string
	for _, p := range s.pages {
		if p.Redirect == nil {
			continue
		}
		target, fragment, ok := s.resolve(p.File, "")
		if !ok {
			continue
		}
		dest := target.URL
		if fragment != "" {
			dest += "#" + fragmentID(fragment)
		}
		if p.URL == dest {
			continue
		}
		lines = append(lines, fmt.Sprintf("%s %s 301", strings.TrimSuffix(p.URL, "/"), dest))
	}
	lines = append(lines, "/wiki/Main_Page / 301", "/wiki/"+homeFile+" / 301", "/wiki/Special_AllPages /wiki/ 301")
	sort.Strings(lines)

	dir := filepath.Join(out, "static")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	data := "# Generated by cmd/migratewiki: redirects of the old wiki's redirect pages.\n" +
		strings.Join(lines, "\n") + "\n"
	return os.WriteFile(filepath.Join(dir, "_redirects"), []byte(data), 0o644)
}

func copyImages(src, dst string) error {
	if err := os.MkdirAll(dst, 0o755); err != nil {
		return err
	}
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		imageDir := filepath.Join(dst, sourceBucket(e.Name()))
		if err := os.MkdirAll(imageDir, 0o755); err != nil {
			return err
		}
		if err := copyFile(filepath.Join(src, e.Name()), filepath.Join(imageDir, e.Name())); err != nil {
			return err
		}
	}
	return nil
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	if _, err := io.Copy(out, in); err != nil {
		out.Close()
		return err
	}
	return out.Close()
}

func (s *site) report(logger *slog.Logger, emitted map[string]*page) {
	var main, categories int
	for _, p := range emitted {
		if p.Kind == kindCategory {
			categories++
		} else {
			main++
		}
	}
	logger.Info("Migration finished", "pages", main, "category_pages", categories,
		"links_rewritten", s.rewrittenLinks, "links_to_dropped_namespaces", s.droppedLinks,
		"links_to_missing_pages", len(s.missingLinks),
		"images_missing", len(s.missingImages))

	for _, m := range []struct {
		name string
		data map[string]int
	}{{"Link to missing page", s.missingLinks}, {"Missing image", s.missingImages}} {
		keys := make([]string, 0, len(m.data))
		for k := range m.data {
			keys = append(keys, k)
		}
		sort.Slice(keys, func(i, j int) bool { return m.data[keys[i]] > m.data[keys[j]] })
		for _, k := range keys[:min(len(keys), 15)] {
			logger.Warn(m.name, "name", k, "count", m.data[k])
		}
	}
}
