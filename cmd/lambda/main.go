package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/hellodhlyn/buslive-api/internal/httpserver"
)

func main() {
	handler := httpserver.NewHandler()
	lambda.Start(httpadapter.New(handler).ProxyWithContext)
}
