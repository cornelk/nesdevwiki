.PHONY: help serve build index preview check test lint clean migrate

# Hugo cache directory that the snap package of Hugo is allowed to use.
CACHE_DIR ?= $(HOME)/snap/hugo/common/cache
# Pagefind version, keep in sync with the Cloudflare Pages build command in README.md.
PAGEFIND ?= pagefind@1.4.0
# Extracted backup directory for "make migrate" (contains wikipages/ and wiki-images/).
SRC ?=

help:
	@echo "Available targets:"
	@echo "  serve    - Start Hugo dev server on port 12346 (no search index)"
	@echo "  build    - Build the site into public/ and create the Pagefind search index"
	@echo "  index    - Create the Pagefind search index for an existing public/"
	@echo "  preview  - Build, then serve public/ on port 8080 with working search"
	@echo "  check    - Check internal links, images and #anchors of the built site"
	@echo "  test     - Run Go tests"
	@echo "  lint     - Run gofmt and go vet"
	@echo "  migrate  - One-time import of the wiki backup (usage: make migrate SRC=/path/to/dump)"
	@echo "  clean    - Remove build output and Hugo cache"

# Start development server
serve:
	hugo server -D --bind 0.0.0.0 -p 12346 --cacheDir $(CACHE_DIR)

# Build for production
build:
	hugo --minify --cacheDir $(CACHE_DIR)
	$(MAKE) index

# Create the search index
index:
	npx --yes $(PAGEFIND) --site public --glob "wiki/**/*.html" --root-selector "#main-content"

# Build and serve the result, the search only works on built output
preview: build
	python3 -m http.server -d public 8080

# Check internal links, images and #anchors. Anchors that never existed on the old wiki are
# listed in docs/migration-issues.md, so this reports them without failing the build.
check: build
	go run ./cmd/checklinks -public public -anchors=false
	-go run ./cmd/checklinks -public public | grep "missing anchor" | wc -l | xargs -I{} echo "{} known broken #anchors, see docs/migration-issues.md"

test:
	go test ./cmd/...

lint:
	@test -z "$$(gofmt -l cmd)" || (gofmt -l cmd; exit 1)
	go vet ./cmd/...

# The import is a one-time operation: the generated Markdown is edited by hand afterwards and the
# command refuses to run when content/wiki already exists.
migrate:
	@test -n "$(SRC)" || (echo "usage: make migrate SRC=/path/to/extracted/dump"; exit 2)
	go run ./cmd/migratewiki -src $(SRC) -out .

clean:
	rm -rf public/ resources/ .hugo_build.lock
