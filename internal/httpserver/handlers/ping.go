package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Ping(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	w.Write([]byte("pong"))
}
