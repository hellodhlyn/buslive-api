package businfo_test

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/hellodhlyn/buslive-api/pkg/businfo"
	"github.com/stretchr/testify/assert"
)

var seoulStore businfo.Store

func init() {
	serveFile := func(w http.ResponseWriter, filename string) {
		file, _ := os.Open(filename)
		w.Header().Set("Content-Type", "application/xml")
		io.Copy(w, file)
	}

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path == "/api/rest/stationinfo/getStationByPos" {
			serveFile(w, "../../test/testdata/seoul_bus_getStationByPos.xml")
		} else if req.URL.Path == "/api/rest/stationinfo/getStationByUid" {
			serveFile(w, "../../test/testdata/seoul_bus_getStationByUid.xml")
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))

	seoulStore = businfo.NewSeoulBusStore(testServer.URL, "test")
}

func TestSeoulBusStore_GetRegion(t *testing.T) {
	assert.Equal(t, "seoul", seoulStore.GetRegion())
}

func TestSeoulBusStore_GetNearbyStations(t *testing.T) {
	stations, err := seoulStore.GetNearbyStations(context.Background(), 127.0277, 37.4979)
	assert.Nil(t, err)

	assert.Equal(t, 42, len(stations))
	assert.Equal(t, "강남역", stations[0].Name)
	assert.Equal(t, "seoul", stations[0].Region)
	assert.Equal(t, "22339", stations[0].Code)
	assert.Equal(t, "22339", stations[0].ID)
	assert.Equal(t, 37.4975009943, stations[0].Position.Lat)
	assert.Equal(t, 127.0268505284, stations[0].Position.Lng)
}

func TestSeoulBusStore_GetStationArrivals(t *testing.T) {
	arrivals, err := seoulStore.GetStationArrivals(context.Background(), "22339")
	assert.Nil(t, err)

	assert.Equal(t, 3, len(arrivals))
	assert.Equal(t, "340", arrivals[0].RouteName)
	assert.Equal(t, "수협서초지점", arrivals[0].NextStationName)
	assert.Equal(t, 2, len(arrivals[0].Positions))
	assert.Equal(t, 277, arrivals[0].Positions[0].RemainingSeconds)
	assert.Equal(t, 2, arrivals[0].Positions[0].RemainingStops)
	assert.Equal(t, 1039, arrivals[0].Positions[1].RemainingSeconds)
	assert.Equal(t, 8, arrivals[0].Positions[1].RemainingStops)

	assert.Equal(t, "N61", arrivals[1].RouteName)
	assert.Equal(t, "강남역.역삼세무서", arrivals[1].NextStationName)
	assert.Equal(t, 0, len(arrivals[1].Positions))
}
