package businfo

import (
	"context"
	"encoding/xml"
	"errors"
	"io"
	"net/http"
	"strconv"

	"golang.org/x/exp/slices"
)

// SeoulBusStore 는 서울시 버스 정보를 제공하는 API를 사용하는 Store입니다.
// https://www.data.go.kr/tcs/dss/selectApiDataDetailView.do?publicDataPk=15000303
type SeoulBusStore struct {
	host   string
	apiKey string

	httpClient *http.Client
}

func NewSeoulBusStore(host, apiKey string) *SeoulBusStore {
	return &SeoulBusStore{
		host:       host,
		apiKey:     apiKey,
		httpClient: http.DefaultClient,
	}
}

type responseData[T interface{}] struct {
	Header struct {
		Code    string `xml:"headerCd"`
		Message string `xml:"headerMsg"`
	} `xml:"msgHeader"`
	Body struct {
		XMLName xml.Name `xml:"msgBody"`
		Data    T        `xml:"itemList"`
	}
}

type stationData struct {
	StationName string `xml:"stationNm"`
	StationID   string `xml:"stationId"`
	StationCode string `xml:"arsId"`
	Lat         string `xml:"gpsY"`
	Lng         string `xml:"gpsX"`
}

// GetNearbyStations 는 지정된 좌표 주변의 정류장 목록을 반환합니다.
func (s *SeoulBusStore) GetNearbyStations(ctx context.Context, lat, lng float64) ([]Station, error) {
	path := "/api/rest/stationinfo/getStationByPos"
	params := map[string]string{
		"tmX":    strconv.FormatFloat(lng, 'f', 6, 64),
		"tmY":    strconv.FormatFloat(lat, 'f', 6, 64),
		"radius": "500",
	}

	var data responseData[[]stationData]
	err := s.requestGet(ctx, path, params, &data)
	if err != nil {
		return nil, err
	} else if data.Header.Code != "0" {
		return nil, errors.New(data.Header.Message)
	}

	result := make([]Station, len(data.Body.Data))
	for i, station := range data.Body.Data {
		lat, _ := strconv.ParseFloat(station.Lat, 64)
		lng, _ := strconv.ParseFloat(station.Lng, 64)
		result[i] = Station{
			Name:     station.StationName,
			ID:       station.StationID,
			Code:     station.StationCode,
			Position: StationPosition{Lat: lat, Lng: lng},
		}
	}

	// 경유 정류장은 응답에 포함하지 않음
	result = slices.DeleteFunc(result, func(i Station) bool {
		return i.Code == "" || i.Code == "0"
	})
	return result, nil
}

func (s *SeoulBusStore) requestGet(ctx context.Context, path string, params map[string]string, data interface{}) error {
	req, _ := http.NewRequest(http.MethodGet, s.host+path, nil)
	req = req.WithContext(ctx)

	q := req.URL.Query()
	q.Add("ServiceKey", s.apiKey)
	for k, v := range params {
		q.Add(k, v)
	}

	res, err := s.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return errors.New("request failed: " + res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	return xml.Unmarshal(body, data)
}