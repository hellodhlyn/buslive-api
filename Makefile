.PHONY: test
test:
	@go test -cover ./...

.PHONY: server.run
server.run:
	@go run ./cmd/server

.PHONY: server.build
server.build:
	@go build -o build/server ./cmd/server

.PHONY: lambda.build
lambda.build:
	@GOOS=linux GOARCH=amd64 go build -o build/lambda ./cmd/lambda

.PHONY: lambda.deploy
lambda.deploy:
	@terraform -chdir=deploy/terraform apply
