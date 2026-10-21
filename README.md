# NESdev Wiki

Static rebuild of the [NESdev Wiki](https://www.nesdev.org/wiki/) on [Hugo](https://gohugo.io/), published on
Cloudflare Pages at <https://nesdev.cornel.contact/>. The content was imported once from the 2026-10 backup of the
old wiki and is now plain Markdown in `content/`.

## Layout

| Path | Content |
| --- | --- |
| `content/wiki/` | Articles, one Markdown file per page. The front matter `url` keeps the old `/wiki/<name>/` address. |
| `content/category/` | Category pages. Their member lists are generated from the `wikiCategories` of the articles. |
| `content/_index.md`, `content/search.md` | Home page and search page. |
| `layouts/`, `assets/css/wiki.css` | Wiki templates (category and A-Z listings, search) and the style. The colors are the ones of the old wiki's MonoBook skin, listed at the top of `wiki.css`. |
| `data/menu/` | The sidebar menu of the theme. |
| `static/brand.svg` | Logo, the green NesDev logo of the old wiki (`NESdev-logo-flat.svg` from the backup). |
| `themes/hugo-geekdoc/` | [Geekdoc](https://github.com/thegeeklab/hugo-geekdoc) v4.1.4, vendored from its release tarball. |
| `static/wiki-images/` | All images and files of the old wiki. |
| `static/_redirects` | Redirects of the old wiki's redirect pages, for Cloudflare Pages. |
| `cmd/migratewiki/` | The one-time import of the backup, see below. |
| `cmd/checklinks/` | Link and anchor checker for the built site. |
| `docs/migration-issues.md` | Known problems of the imported content. |

## Working on the wiki

Requirements: Hugo (0.160 or newer, the theme needs it), Go, Node.js (for Pagefind search).

```bash
make serve     # dev server, http://localhost:12346 (search does not work here)
make preview   # production build and local server on http://localhost:8080 with search
make check     # build and verify links, images and anchors
make test lint
```

To edit a page, change its Markdown file in `content/wiki/`, or use the "Edit this page on GitHub" link at the bottom of
each page (`wikiEditURL` in `hugo.toml`). Link to other pages with a relative URL like
`[mapper](../Mapper/)`. Category membership is set in the front matter:

```yaml
wikiCategories:
  - name: "ASIC mappers"
    url: "/category/ASIC_mappers/"
```

Tables and infoboxes are HTML inside the Markdown, which is allowed by `markup.goldmark.renderer.unsafe`.

## Search

Search uses [Pagefind](https://pagefind.app/) instead of the built-in search of the theme, which would put the
content of all pages into one file. It indexes the articles below `/wiki/` (`make index`) and writes the files to
`public/pagefind/`, so it works in `make preview` and on Cloudflare, not in `hugo server`.

## Deployment

Cloudflare Pages builds the GitHub repository on every push.

| Setting | Value |
| --- | --- |
| Build command | `hugo --minify && npx --yes pagefind@1.4.0 --site public --glob "wiki/**/*.html" --root-selector "#main-content"` |
| Build output directory | `public` |
| Environment variable | `HUGO_VERSION=0.168.0` |
| Custom domain | `nesdev.cornel.contact` |

## The import of the old wiki

The backup (`202610-NesDevWiki.zip`, not committed) is a static XHTML export. `cmd/migratewiki` converted it with
pandoc: it rewrote links and anchors, extracted categories, turned redirect pages into `static/_redirects`, copied
the images, and dropped the Talk, User, File, Template and MediaWiki namespaces.

The import was a one-time step and the result is edited by hand now. The command refuses to run when
`content/wiki` exists, to protect those edits. To repeat it in a fresh checkout:

```bash
unzip 202610-NesDevWiki.zip -d /tmp/nesdev-dump
make migrate SRC=/tmp/nesdev-dump   # needs pandoc
```

The known problems of the imported content (links that were already broken on the old wiki, dropped pages,
changed URLs) are listed in [docs/migration-issues.md](docs/migration-issues.md).

## License

The wiki declares its content public domain, see the
[general disclaimer](content/wiki/Nesdev_wiki_General_disclaimer.md). The Geekdoc theme is MIT licensed
(`themes/hugo-geekdoc/LICENSE`, its icons have their own licenses, see its README).
