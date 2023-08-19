package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

// Parameters:
//   - lat (float, required)
//   - lng (float, required)
func ListStationsByPosition(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	lat, latOk := queryFloat64(req, "lat")
	lng, lngOk := queryFloat64(req, "lng")
	if !latOk || !lngOk {
		responseError(w, http.StatusBadRequest, "Required parameters are missing")
		return
	}

	result, err := busInfoStore.GetNearbyStations(req.Context(), lat, lng)
	if err != nil {
		log.Error(err)
		responseError(w, http.StatusInternalServerError, "Unexpected error")
		return
	}

	responseJSON(w, result)
}

// Parameters:
//   - region = "seoul"
//   - stationId (string, required)
func ListArrivalsByStation(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	stationId := ps.ByName("stationId")
	if stationId == "" {
		responseError(w, http.StatusBadRequest, "Required parameters are missing")
		return
	}

	result, err := busInfoStore.GetStationArrivals(req.Context(), stationId)
	if err != nil {
		log.Error(err)
		responseError(w, http.StatusInternalServerError, "Unexpected error")
		return
	}

	responseJSON(w, result)
}
