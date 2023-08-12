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
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := busInfoStore.GetNearbyStations(req.Context(), lat, lng)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseJSON(w, result)
}
