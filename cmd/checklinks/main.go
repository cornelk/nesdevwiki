// Command checklinks verifies the internal links, images and #fragments of a built Hugo site.
//
// Usage:
//
//	go run ./cmd/checklinks -public public
//
// Run it after the Pagefind index has been built, as the search page references its files.
package main

import (
	"flag"
	"fmt"
	"io/fs"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

// pageInfo holds what is needed from a page to validate links to it.
type pageInfo struct {
	url   string
	ids   map[string]struct{}
	links []link
}

type link struct {
	raw string
	tag string
}

func main() {
	public := flag.String("public", "public", "directory of the built site")
	anchors := flag.Bool("anchors", true, "also check that #fragments exist on the target page")
	flag.Parse()

	pages, err := loadPages(*public)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(2)
	}

	broken := check(*public, pages, *anchors)
	for _, b := range broken {
		fmt.Println(b)
	}
	fmt.Printf("%d pages checked, %d broken references\n", len(pages), len(broken))
	if len(broken) > 0 {
		os.Exit(1)
	}
}

// loadPages parses every HTML file below the public directory.
func loadPages(public string) (map[string]*pageInfo, error) {
	var files []string
	err := filepath.WalkDir(public, func(p string, d fs.DirEntry, err error) error {
		if err == nil && !d.IsDir() && strings.HasSuffix(p, ".html") {
			files = append(files, p)
		}
		return err
	})
	if err != nil {
		return nil, err
	}

	var (
		mu    sync.Mutex
		wg    sync.WaitGroup
		first error
		pages = make(map[string]*pageInfo, len(files))
		sem   = make(chan struct{}, 8)
	)
	for _, file := range files {
		wg.Add(1)
		sem <- struct{}{}
		go func() {
			defer wg.Done()
			defer func() { <-sem }()
			info, err := parseFile(public, file)
			mu.Lock()
			defer mu.Unlock()
			if err != nil && first == nil {
				first = err
			}
			if info != nil {
				pages[info.url] = info
			}
		}()
	}
	wg.Wait()
	return pages, first
}

func parseFile(public, file string) (*pageInfo, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	doc, err := html.Parse(f)
	if err != nil {
		return nil, fmt.Errorf("parsing %s: %w", file, err)
	}

	rel, _ := filepath.Rel(public, file)
	pageURL := "/" + filepath.ToSlash(rel)
	if strings.HasSuffix(pageURL, "/index.html") {
		pageURL = strings.TrimSuffix(pageURL, "index.html")
	}
	info := &pageInfo{url: pageURL, ids: make(map[string]struct{})}

	var walk func(*html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				switch {
				case a.Key == "id", a.Key == "name" && n.Data == "a":
					info.ids[a.Val] = struct{}{}
				case a.Key == "href" && (n.Data == "a" || n.Data == "link"), a.Key == "src":
					info.links = append(info.links, link{raw: a.Val, tag: n.Data})
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(doc)
	return info, nil
}

// check resolves every link of every page and returns a description of each broken one.
func check(public string, pages map[string]*pageInfo, anchors bool) []string {
	var broken []string
	for _, p := range pages {
		for _, l := range p.links {
			if msg := checkLink(public, pages, p, l, anchors); msg != "" {
				broken = append(broken, msg)
			}
		}
	}
	sort.Strings(broken)
	return broken
}

func checkLink(public string, pages map[string]*pageInfo, from *pageInfo, l link, anchors bool) string {
	u, err := url.Parse(l.raw)
	if err != nil {
		return fmt.Sprintf("%s: invalid URL %q", from.url, l.raw)
	}
	if u.Scheme != "" || u.Host != "" {
		return ""
	}

	target := from.url
	switch {
	case strings.HasPrefix(u.Path, "/"):
		target = u.Path
	case u.Path != "":
		target = path.Join(from.url, u.Path)
		if strings.HasSuffix(u.Path, "/") {
			target += "/"
		}
	}

	fragment := u.Fragment
	if p, ok := pages[target]; ok {
		if anchors && fragment != "" && !hasAnchor(p, fragment) {
			return fmt.Sprintf("%s: missing anchor %q on %s", from.url, fragment, target)
		}
		return ""
	}

	// Assets and other files.
	full := filepath.Join(public, filepath.FromSlash(target))
	if st, err := os.Stat(full); err == nil && !st.IsDir() {
		return ""
	}
	if st, err := os.Stat(filepath.Join(full, "index.html")); err == nil && !st.IsDir() {
		return ""
	}
	return fmt.Sprintf("%s: broken %s %q -> %s", from.url, l.tag, l.raw, target)
}

func hasAnchor(p *pageInfo, fragment string) bool {
	if fragment == "top" {
		return true
	}
	if _, ok := p.ids[fragment]; ok {
		return true
	}
	return false
}
