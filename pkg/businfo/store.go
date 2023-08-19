package businfo

import "context"

type Station struct {
	Name   string `json:"name"`   // 강남역
	Region string `json:"region"` // seoul
	ID     string `json:"id"`     // 121000262
	Code   string `json:"code"`   // 22339

	Position StationPosition `json:"position"`
}

type StationPosition struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type StationArrivals struct {
	RouteID         string                   `json:"routeId"`         // 100100055
	RouteName       string                   `json:"routeName"`       // 340
	NextStationName string                   `json:"nextStationName"` // 수협서초지점
	Positions       []StationArrivalPosition `json:"positions"`
	UpdatedAt       int64                    `json:"updatedAt"`
}

type StationArrivalPosition struct {
	RemainingSeconds int `json:"remainingSeconds"`
	RemainingStops   int `json:"remainingStops"`
}

type Store interface {
	GetRegion() string
	GetNearbyStations(ctx context.Context, lat, lng float64) ([]Station, error)
	GetStationArrivals(ctx context.Context, stationID string) ([]StationArrivals, error)
}
