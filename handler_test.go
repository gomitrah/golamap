package golamap

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDirections(t *testing.T) {
	t.Run("Invalid origin & destination", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.HttpService = &MockStruct{}
		_, err := olaMap.GetDirections("", "")
		expectedErr := fmt.Errorf("Missing required query parameters: 'origin' and/or 'destination'")
		assert.Exactly(t, err, expectedErr)

	})
	t.Run("Invalid OAuth token", func(t *testing.T) {
		olaMap := &OLAMap{}
		_, err := olaMap.GetDirections("mock-origin", "mock-destination")
		expectedErr := fmt.Errorf("Invalid OAuth token")
		assert.Exactly(t, err, expectedErr)
	})
	t.Run("success", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.Token = "mockToken"
		mocking := &MockStruct{}
		olaMap.HttpService = mocking
		_, err := olaMap.GetDirections("mock-origin", "mock-destination")
		assert.Nil(t, err)
		if mocking.StatusCode != 200 {
			t.Error("expected : ", 200, "got : ", mocking.StatusCode)
		}
		assert.Equal(t, DirectionResponse, mocking.MockBody)

	})
}

func TestPlaceAutoComplete(t *testing.T) {
	t.Run("Invalid input", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.HttpService = &MockStruct{}
		_, err := olaMap.PlaceAutoComplete("")
		expectedErr := fmt.Errorf("Missing required query parameters: 'input'")
		assert.Exactly(t, err, expectedErr)

	})
	t.Run("Invalid OAuth token", func(t *testing.T) {
		olaMap := &OLAMap{}
		_, err := olaMap.PlaceAutoComplete("mock-input")
		expectedErr := fmt.Errorf("Invalid OAuth token")
		assert.Exactly(t, err, expectedErr)

	})
	t.Run("success", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.Token = "mockToken"
		mocking := &MockStruct{}
		olaMap.HttpService = mocking
		_, err := olaMap.PlaceAutoComplete("mock-input")
		assert.Nil(t, err)
		if mocking.StatusCode != 200 {
			t.Error("expected : ", 200, "got : ", mocking.StatusCode)
		}
		assert.Equal(t, AutoCompleteResponse, mocking.MockBody)
	})
}

func TestGecode(t *testing.T) {
	t.Run("Invalid address", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.HttpService = &MockStruct{}
		_, err := olaMap.GeoCode("", "", "")
		expectedErr := fmt.Errorf("Missing required query parameters: 'address'")
		assert.Exactly(t, err, expectedErr)

	})
	t.Run("Invalid OAuth token", func(t *testing.T) {
		olaMap := &OLAMap{}
		_, err := olaMap.GeoCode("mock-address", "mock-bounds", "mock-language")
		expectedErr := fmt.Errorf("Invalid OAuth token")
		assert.Exactly(t, err, expectedErr)

	})
	t.Run("success", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.Token = "mockToken"
		mocking := &MockStruct{}
		olaMap.HttpService = mocking
		_, err := olaMap.GeoCode("mock-address", "mock-bounds", "mock-language")
		assert.Nil(t, err)
		if mocking.StatusCode != 200 {
			t.Error("expected : ", 200, "got : ", mocking.StatusCode)
		}
		assert.Equal(t, GeoCodeResponse, mocking.MockBody)

	})
}

func TestReverseGeocode(t *testing.T) {
	t.Run("Invalid latlng", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.HttpService = &MockStruct{}
		_, err := olaMap.ReverseGeocode("")
		expectedErr := fmt.Errorf("Missing required query parameters: 'latlng'")
		assert.Exactly(t, err, expectedErr)

	})
	t.Run("Invalid OAuth token", func(t *testing.T) {
		olaMap := &OLAMap{}
		_, err := olaMap.ReverseGeocode("mock-latlng")
		expectedErr := fmt.Errorf("Invalid OAuth token")
		assert.Exactly(t, err, expectedErr)

	})
	t.Run("success", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.Token = "mockToken"
		mocking := &MockStruct{}
		olaMap.HttpService = mocking
		_, err := olaMap.ReverseGeocode("mock-latlng")
		assert.Nil(t, err)
		if mocking.StatusCode != 200 {
			t.Error("expected : ", 200, "got : ", mocking.StatusCode)
		}
		assert.Equal(t, ReverseGeocodeResponse, mocking.MockBody)

	})
}

func TestGetDistanceMatrix(t *testing.T) {
	t.Run("Invalid origins & destinations", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.HttpService = &MockStruct{}
		_, err := olaMap.GetDistanceMatrix("", "")
		expectedErr := fmt.Errorf("Missing required query parameters: 'origin' and/or 'destination'")
		assert.Exactly(t, err, expectedErr)

	})
	t.Run("Invalid OAuth token", func(t *testing.T) {
		olaMap := &OLAMap{}
		_, err := olaMap.GetDistanceMatrix("mock-origins", "mock-destinations")
		expectedErr := fmt.Errorf("Invalid OAuth token")
		assert.Exactly(t, err, expectedErr)

	})
	t.Run("success", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.Token = "mockToken"
		mocking := &MockStruct{}
		olaMap.HttpService = mocking
		_, err := olaMap.GetDistanceMatrix("mock-origins", "mock-destinations")
		assert.Nil(t, err)
		if mocking.StatusCode != 200 {
			t.Error("expected : ", 200, "got : ", mocking.StatusCode)
		}
		assert.Equal(t, DistanceMatrixResponse, mocking.MockBody)
	})
}
