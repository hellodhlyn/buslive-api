.PHONY: dev
dev:
	wrangler dev

.PHONY: build
build:
	go run github.com/syumai/workers/cmd/workers-assets-gen@latest
	tinygo build -o ./build/app.wasm -target wasm -no-debug ./cmd/worker

.PHONY: test
test:
	go test -cover ./pkg/...

.PHONY: deploy
deploy:
	wrangler deploy
