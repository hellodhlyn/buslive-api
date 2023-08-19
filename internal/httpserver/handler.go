package httpserver

import (
	"net/http"

	"github.com/hellodhlyn/buslive-api/internal/httpserver/handlers"
	"github.com/julienschmidt/httprouter"
)

func NewHandler() http.Handler {
	router := httprouter.New()
	router.GET("/api/ping", handlers.Ping)
	router.GET("/api/v1/stations_by_position", handlers.ListStationsByPosition)
	router.GET("/api/v1/stations/seoul/:stationId/arrivals", handlers.ListArrivalsByStation)
	return router
}
