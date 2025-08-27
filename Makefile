BINPATH = bin/app.exe

build: build-templ build-app build-css

build-app:
		go build -o $(BINPATH) cmd/main.go

.PHONY: build-templ
build-templ:
		templ generate

.PHONY: build-css
build-css:
		npm --prefix web run build:css

build-templ:
		templ generate

.PHONY: run
run: build
		$(BINPATH)

.PHONY: watch
watch:
		$(MAKE)  -j3 watch-app watch-templ watch-css

.PHONY: watch-app
watch-app:
		go run github.com/air-verse/air@latest \
		--build.cmd="$(MAKE) build-app" \
		--build.bin  "$(BINPATH)" \
		--build.include_ext "go" \
		--build.exclude_dir "bin,web" \

.PHONY: watch-templ
watch-templ:
		templ generate \
		--watch \
		--proxy="http://localhost:8080" \
		--open-browser=false \

.PHONY: watch-css
watch-css:
		npm --prefix web run watch:css

.PHONY: fmt
fmt:
		templ fmt internal/view