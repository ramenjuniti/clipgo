.PHONY: install
install: 
	yarn

.PHONY: test
test:
	GOOS=js GOARCH=wasm go test -v -exec="$(GOROOT)/misc/wasm/go_js_wasm_exec" ./...

.PHONY: build
build:
	yarn webpack
	GOOS=js GOARCH=wasm go build -o main.wasm clipgo/src/background
	mv main.wasm build && cp src/manifest.json src/icons/*.png build

.PHONY: clean
clean:
	rm -rf build/

.PHONY: destroy
destroy:
	rm -rf node_modules/