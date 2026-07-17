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

func TestArrayOfData(t *testing.T) {
	t.Run("Invalid datasetname", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.HttpService = &MockStruct{}
		_, err := olaMap.ArrayOfData("")
		expectedErr := fmt.Errorf("Missing required query parameters: 'datasetname'")
		assert.Exactly(t, err, expectedErr)

	})
	t.Run("Invalid OAuth token", func(t *testing.T) {
		olaMap := &OLAMap{}
		_, err := olaMap.ArrayOfData("mock-datasetname")
		expectedErr := fmt.Errorf("Invalid OAuth token")
		assert.Exactly(t, err, expectedErr)

	})
	t.Run("success", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.Token = "mockToken"
		mocking := &MockStruct{}
		olaMap.HttpService = mocking
		_, err := olaMap.ArrayOfData("mock-datasetname")
		assert.Nil(t, nil, err)
		if mocking.StatusCode != 200 {
			t.Error("expected : ", 200, "got : ", mocking.StatusCode)
		}
		assert.Equal(t, ArrayOfDataResponse, mocking.MockBody)

	})
}

func TestGetStyleDetails(t *testing.T) {
	t.Run("Invalid stylename", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.HttpService = &MockStruct{}
		_, err := olaMap.GetStyleDetails("")
		expectedErr := fmt.Errorf("Missing required query parameters: 'stylename'")
		assert.Exactly(t, err, expectedErr)

	})
	t.Run("Invalid OAuth token", func(t *testing.T) {
		olaMap := &OLAMap{}
		_, err := olaMap.GetStyleDetails("mock-stylename")
		expectedErr := fmt.Errorf("Invalid OAuth token")
		assert.Exactly(t, err, expectedErr)

	})
	t.Run("success", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.Token = "mockToken"
		mocking := &MockStruct{}
		olaMap.HttpService = mocking
		_, err := olaMap.GetStyleDetails("mock-stylename")
		assert.Nil(t, nil, err)
		if mocking.StatusCode != 200 {
			t.Error("expected : ", 200, "got : ", mocking.StatusCode)
		}
		assert.Equal(t, StyleDetailResponse, mocking.MockBody)

	})
}

func TestGetMapStyle(t *testing.T) {
	t.Run("Invalid OAuth token", func(t *testing.T) {
		olaMap := &OLAMap{}
		_, err := olaMap.GetMapStyle()
		expectedErr := fmt.Errorf("Invalid OAuth token")
		assert.Exactly(t, err, expectedErr)

	})

	t.Run("success", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.Token = "mockToken"
		mocking := &MockStruct{}
		olaMap.HttpService = mocking
		_, err := olaMap.GetMapStyle()
		assert.Nil(t, nil, err)
		if mocking.StatusCode != 200 {
			t.Error("expected : ", 200, "got : ", mocking.StatusCode)
		}
		assert.Equal(t, MapStyleResponse, mocking.MockBody)
	})
}

func TestGetPlaceDetail(t *testing.T) {
	t.Run("Invalid 'placeid'", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.HttpService = &MockStruct{}
		_, err := olaMap.GetPlaceDetail("")
		expectedErr := fmt.Errorf("Missing required query parameters: 'placeid'")
		assert.Exactly(t, err, expectedErr)

	})
	t.Run("Invalid OAuth token", func(t *testing.T) {
		olaMap := &OLAMap{}
		_, err := olaMap.GetPlaceDetail("mock-placeid")
		expectedErr := fmt.Errorf("Invalid OAuth token")
		assert.Exactly(t, err, expectedErr)

	})
	t.Run("success", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.Token = "mockToken"
		mocking := &MockStruct{}
		olaMap.HttpService = mocking
		_, err := olaMap.GetPlaceDetail("mock-placeid")
		assert.Nil(t, nil, err)
		if mocking.StatusCode != 200 {
			t.Error("expected : ", 200, "got : ", mocking.StatusCode)
		}
		assert.Equal(t, PlaceDetailResponse, mocking.MockBody)

	})
}

func TestGetNearBySearch(t *testing.T) {
	t.Run("Invalid 'layers' and 'location'", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.HttpService = &MockStruct{}
		_, err := olaMap.GetNearBySearch(NearBySearch{Layers: "", Location: ""})
		expectedErr := fmt.Errorf("Missing required query parameters: 'layers' and/or 'location'")
		assert.Exactly(t, err, expectedErr)

	})
	t.Run("Invalid OAuth token", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.HttpService = &MockStruct{}
		fmt.Printf("Mocking = %+v", olaMap)
		_, err := olaMap.GetNearBySearch(NearBySearch{Layers: "Mock-Layers", Location: "Mock-Location"})
		expectedErr := fmt.Errorf("Invalid OAuth token")
		assert.Exactly(t, err, expectedErr)

	})
	t.Run("success", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.Token = "mockToken"
		mocking := &MockStruct{}
		olaMap.HttpService = mocking
		_, err := olaMap.GetNearBySearch(NearBySearch{Layers: "Mock-Layers", Location: "Mock-Location"})
		assert.Nil(t, nil, err)
		if mocking.StatusCode != 200 {
			t.Error("expected : ", 200, "got : ", mocking.StatusCode)
		}
		assert.Equal(t, NearBySearchResponses, mocking.MockBody)
	})
}

func TestGetTextSearch(t *testing.T) {
	t.Run("Invalid input", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.HttpService = &MockStruct{}
		_, err := olaMap.GetTextSearch(TextSearch{Input: ""})
		expectedErr := fmt.Errorf("Missing required query parameters: 'input'")
		assert.Exactly(t, err, expectedErr)

	})
	t.Run("Invalid OAuth token", func(t *testing.T) {
		olaMap := &OLAMap{}
		_, err := olaMap.GetTextSearch(TextSearch{Input: "mock-input"})
		expectedErr := fmt.Errorf("Invalid OAuth token")
		assert.Exactly(t, err, expectedErr)

	})
	t.Run("success", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.Token = "mockToken"
		mocking := &MockStruct{}
		olaMap.HttpService = mocking
		_, err := olaMap.GetTextSearch(TextSearch{Input: "mock-input"})
		assert.Nil(t, nil, err)
		if mocking.StatusCode != 200 {
			t.Error("expected : ", 200, "got : ", mocking.StatusCode)
		}
		assert.Equal(t, TextSearchResponse, mocking.MockBody)
	})
}
