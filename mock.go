package golamap

import (
	"errors"
	"strings"
)

var (
	MockDirectionsURL     = "routing/v1/directions"
	MockDistanceMatrixURL = "routing/v1/distanceMatrix"
	MockAutoCompleteURL   = "places/v1/autocomplete"
	MockGecodeURL         = "places/v1/geocode"
	MockReverseGecodeURL  = "places/v1/reverse-geocode"
)

type MockInterface interface {
	SendOlaMapRequest(method, url, requestID, oauthToken string, responseObj interface{}) error
}
type MockStruct struct {
	MockBody   string
	StatusCode int
}

func (mock *MockStruct) SendOlaMapRequest(method, url, requestID, oauthToken string, responseObj interface{}) error {
	switch {
	case strings.Contains(url, MockDirectionsURL):
		mock.StatusCode = 200
		mock.MockBody = DirectionResponse
	case strings.Contains(url, MockDistanceMatrixURL):
		mock.StatusCode = 200
		mock.MockBody = DistanceMatrixResponse
	case strings.Contains(url, MockGecodeURL):
		mock.StatusCode = 200
		mock.MockBody = GeoCodeResponse
	case strings.Contains(url, MockReverseGecodeURL):
		mock.StatusCode = 200
		mock.MockBody = ReverseGeocodeResponse
	case strings.Contains(url, MockAutoCompleteURL):
		mock.StatusCode = 200
		mock.MockBody = AutoCompleteResponse
	default:
		mock.StatusCode = 400
		mock.MockBody = ""
		return errors.New("Invalid request")
	}
	return nil
}
