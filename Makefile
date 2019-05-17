.PHONY: install
install: 
	yarn
	go get -u golang.org/x/tools/cmd/goimports
	go get -u golang.org/x/lint/golint

.PHONY: lint
lint: 
	GOOS=js GOARCH=wasm go vet ./src/background
	golint ./src/background

.PHONY: fmt
fmt:
	goimports -w ./src/background

.PHONY: test
test: lint fmt
	GOOS=js GOARCH=wasm go test -v -cover -exec="$(GOROOT)/misc/wasm/go_js_wasm_exec" ./src/background

.PHONY: build
build: lint fmt
	yarn webpack
	GOOS=js GOARCH=wasm go build -o formatter.wasm ./src/background
	mv formatter.wasm build && cp src/manifest.json src/icons/*.png build

.PHONY: clean
clean:
	rm -rf build/

.PHONY: destroy
destroy:
	rm -rf node_modules/