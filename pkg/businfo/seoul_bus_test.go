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
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))

	seoulStore = businfo.NewSeoulBusStore(testServer.URL, "test")
}

func TestSeoulBusStore_GetNearbyStations(t *testing.T) {
	stations, err := seoulStore.GetNearbyStations(context.Background(), 127.0277, 37.4979)
	assert.Nil(t, err)

	assert.Equal(t, 42, len(stations))
	assert.Equal(t, "강남역", stations[0].Name)
	assert.Equal(t, "22339", stations[0].Code)
	assert.Equal(t, "121000262", stations[0].ID)
	assert.Equal(t, 37.4975009943, stations[0].Position.Lat)
	assert.Equal(t, 127.0268505284, stations[0].Position.Lng)
}
