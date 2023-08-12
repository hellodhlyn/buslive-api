package businfo

import "context"

type Station struct {
	Name string `json:"name"` // 강남역
	ID   string `json:"id"`   // 121000262
	Code string `json:"code"` // 22339

	Position StationPosition `json:"position"`
}

type StationPosition struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Store interface {
	GetNearbyStations(ctx context.Context, lat, lng float64) ([]Station, error)
}
