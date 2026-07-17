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

// ArrayOfData
func (o *OLAMap) ArrayOfData(datasetName string) (interface{}, error) {
	if datasetName == "" {
		return nil, errors.New("Missing required query parameters: 'datasetname'")
	}

	oauthToken := o.Token
	if oauthToken == "" {
		return nil, errors.New("Invalid OAuth token")
	}

	// Construct the URL for the Olamaps API request
	apiURL := fmt.Sprintf(ArrayOfDataURL, url.QueryEscape(datasetName))

	var apiResponse ArrayOfData

	// Make the external request
	err := o.HttpService.SendOlaMapRequest("GET", apiURL, o.RequestId, oauthToken, &apiResponse)
	if err != nil {
		return nil, errors.New("failed to send request to Olamaps API")
	}

	return apiResponse, nil
}

// GetStyleDetails
func (o *OLAMap) GetStyleDetails(styleName string) (interface{}, error) {
	if styleName == "" {
		return nil, errors.New("Missing required query parameters: 'stylename'")
	}

	oauthToken := o.Token
	if oauthToken == "" {
		return nil, errors.New("Invalid OAuth token")
	}

	// Construct the URL for the Olamaps API request
	apiURL := fmt.Sprintf(StyleDetailsURL, url.QueryEscape(styleName))

	var apiResponse VectorStyleDetails

	// Make the external request
	err := o.HttpService.SendOlaMapRequest("GET", apiURL, o.RequestId, oauthToken, &apiResponse)
	if err != nil {
		return nil, errors.New("failed to send request to Olamaps API")
	}

	return apiResponse, nil
}

// GetMapStyle
func (o *OLAMap) GetMapStyle() (interface{}, error) {
	oauthToken := o.Token
	if oauthToken == "" {
		return nil, errors.New("Invalid OAuth token")
	}

	apiURL := MapStyleURL

	var apiResponse []VectorMapStyle

	// Make the external request
	err := o.HttpService.SendOlaMapRequest("GET", apiURL, o.RequestId, oauthToken, &apiResponse)
	if err != nil {
		return nil, errors.New("failed to send request to Olamaps API")
	}

	return apiResponse, nil
}

// GetPlaceDetail
func (o *OLAMap) GetPlaceDetail(placeID string) (interface{}, error) {
	if placeID == "" {
		return nil, errors.New("Missing required query parameters: 'placeid'")
	}

	oauthToken := o.Token
	if oauthToken == "" {
		return nil, errors.New("Invalid OAuth token")
	}

	apiURL := fmt.Sprintf(PlaceDetailURL, placeID)

	var apiResponse PlaceDetail

	// Make the external request
	err := o.HttpService.SendOlaMapRequest("GET", apiURL, o.RequestId, oauthToken, &apiResponse)
	if err != nil {
		return nil, errors.New("failed to send request to Olamaps API")
	}

	return apiResponse, nil
}

// GetNearBySearch
func (o *OLAMap) GetNearBySearch(nearBySearch NearBySearch) (interface{}, error) {
	if nearBySearch.Layers == "" || nearBySearch.Location == "" {
		return nil, errors.New("Missing required query parameters: 'layers' and/or 'location'")
	}

	oauthToken := o.Token
	if oauthToken == "" {
		return nil, errors.New("Invalid OAuth token")
	}

	apiURL := fmt.Sprintf(NearBySearchURL, nearBySearch.Layers, nearBySearch.Location, nearBySearch.Types, nearBySearch.Radius, nearBySearch.Strictbounds, nearBySearch.WithCentroid, nearBySearch.Limit)

	var apiResponse NearBySearchResponse

	// Make the external request
	err := o.HttpService.SendOlaMapRequest("GET", apiURL, o.RequestId, oauthToken, &apiResponse)
	if err != nil {
		return nil, errors.New("failed to send request to Olamaps API")
	}

	return apiResponse, nil
}

// GetTextSearch
func (o *OLAMap) GetTextSearch(textSearch TextSearch) (interface{}, error) {
	// Extract query parameters
	if textSearch.Input == "" {
		return nil, errors.New("Missing required query parameters: 'input'")
	}

	oauthToken := o.Token
	if oauthToken == "" {
		return nil, errors.New("Invalid OAuth token")
	}

	// Construct the API URL
	apiURL := fmt.Sprintf(TextSearchURL, url.QueryEscape(textSearch.Input), textSearch.Location, textSearch.Radius, textSearch.Types, textSearch.Size)

	var apiResponse TextBySearch

	// Make the external request
	err := o.HttpService.SendOlaMapRequest("GET", apiURL, o.RequestId, oauthToken, &apiResponse)
	if err != nil {
		return nil, errors.New("failed to send request to Olamaps API")
	}

	return apiResponse, nil
}
