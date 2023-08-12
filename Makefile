.PHONY: test
test:
	go test -cover ./...

.PHONY: server.run
server.run:
	go run ./cmd/server

.PHONY: server.build
build:
	go build ./cmd/server
