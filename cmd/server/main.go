package main

import (
	"net/http"

	"github.com/hellodhlyn/buslive-api/internal/httpserver"
	log "github.com/sirupsen/logrus"
)

func main() {
	handler := httpserver.NewHandler()

	log.Info("Starting HTTP server on :8080")
	http.ListenAndServe(":8080", handler)
}
