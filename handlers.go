package golamap

import (
	"errors"
	"fmt"
	"net/url"
)

func (o *OLAMap) buildQueryURL(format string, values ...string) string {
	args := make([]interface{}, len(values))
	for i, value := range values {
		args[i] = url.QueryEscape(value)
	}
	return fmt.Sprintf(format, args...)
}

// Get directions
func (o *OLAMap) GetDirections(origin, destination string) (interface{}, error) {
	if origin == "" || destination == "" {
		return nil, errors.New("Missing required query parameters: 'origin' and/or 'destination'")
	}

	oauthToken := o.Token
	if oauthToken == "" {
		return nil, errors.New("Invalid OAuth token")
	}

	url := fmt.Sprintf(DirectionsURL, origin, destination)

	var apiResponse Directions

	// Make external request
	err := o.sendRequest("POST", url, &apiResponse)
	if err != nil {
		return nil, errors.New("failed to send request to Olamaps API")
	}

	return apiResponse, nil
}

// PlaceAutoComplete
func (o *OLAMap) PlaceAutoComplete(input string) (interface{}, error) {
	if input == "" {
		return nil, errors.New("Missing required query parameters: 'input'")
	}

	oauthToken := o.Token
	if oauthToken == "" {
		return nil, errors.New("Invalid OAuth token")
	}

	url := o.buildQueryURL(PlaceAutoCompleteURL, input)

	var apiResponse AutoComplete

	// Make the external request
	err := o.sendRequest("GET", url, &apiResponse)
	if err != nil {
		return nil, errors.New("failed to send request to Olamaps API")
	}

	return apiResponse, nil
}

// GeoCode
func (o *OLAMap) GeoCode(address, bounds, language string) (interface{}, error) {
	if address == "" {
		return nil, errors.New("Missing required query parameters: 'address'")
	}

	oauthToken := o.Token
	if oauthToken == "" {
		return nil, errors.New("Invalid OAuth token")
	}

	// Construct the URL for the Olamaps API request
	url := o.buildQueryURL(GeoCodeURL, address, bounds, language)

	var apiResponse ForwardGecode

	// Make the external request
	err := o.sendRequest("GET", url, &apiResponse)
	if err != nil {
		return nil, errors.New("failed to send request to Olamaps API")
	}

	return apiResponse, nil
}

// ReverseGeocode
func (o *OLAMap) ReverseGeocode(latlng string) (interface{}, error) {
	if latlng == "" {
		return nil, errors.New("Missing required query parameters: 'latlng'")
	}

	oauthToken := o.Token
	if oauthToken == "" {
		return nil, errors.New("Invalid OAuth token")
	}

	// Construct the URL for the API request
	urlWithParams := o.buildQueryURL(ReverseGeocodeURL, latlng)

	var apiResponse ReverseGecode

	// Make the external request
	err := o.sendRequest("GET", urlWithParams, &apiResponse)
	if err != nil {
		return nil, errors.New("failed to send request to Olamaps API")
	}

	return apiResponse, nil
}

// GetDistanceMatrix
func (o *OLAMap) GetDistanceMatrix(origins, destinations string) (interface{}, error) {
	if origins == "" || destinations == "" {
		return nil, errors.New("Missing required query parameters: 'origin' and/or 'destination'")
	}

	oauthToken := o.Token
	if oauthToken == "" {
		return nil, errors.New("Invalid OAuth token")
	}

	// Construct the URL for the Olamaps API request
	url := fmt.Sprintf(DistanceMatrixURL,
		url.QueryEscape(origins), url.QueryEscape(destinations))

	var apiResponse DistanceMatrix

	// Make the external request
	err := o.sendRequest("GET", url, &apiResponse)
	if err != nil {
		return nil, errors.New("failed to send request to Olamaps API")
	}

	return apiResponse, nil
}
