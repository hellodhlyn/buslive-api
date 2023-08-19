package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/hellodhlyn/buslive-api/pkg/businfo"
)

var (
	busInfoStore businfo.Store
)

func init() {
	busInfoStore = businfo.NewSeoulBusStore("http://ws.bus.go.kr", os.Getenv("SEOUL_BUS_API_KEY"))
}

func queryString(req *http.Request, key string) (string, bool) {
	values, ok := req.URL.Query()[key]
	if !ok || len(values) < 1 {
		return "", false
	}
	return values[0], true
}

func queryFloat64(req *http.Request, key string) (float64, bool) {
	value, ok := queryString(req, key)
	if !ok {
		return 0, false
	}

	result, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, false
	}
	return result, true
}

func responseJSON(w http.ResponseWriter, data interface{}) {
	body, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

func responseError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	responseJSON(w, map[string]string{
		"error": message,
	})
}
