# Go OLA Map Wrapper

This Go package provides a wrapper for the OLA Map API, allowing you to interact with various map-related functionalities such as getting directions, geocoding, and more. 

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Available Methods](#available-methods)

## Features
- Easy-to-use Golang wrapper for OLA Map.
- Provides methods to interact with OLA Map and its data structures.
- Supports searching and querying geographical data using OLA codes.
- Lightweight and performant implementation suitable for production use.
- Customizable: Add additional utilities to enhance functionality.

## Installation

1. To install the OLA Map wrapper, use the following command:
   ```bash
   go get github.com/golang-mitrah/golamap
   ```

## Usage

Below is an example of how to use the OLA Map Wrapper in your project:

```go
package main

import (
    "fmt"
    "github.com/golang-mitrah/golamap"
)

func main() {
    // Initialize the OLA Map with a request ID
    olaMap := golamap.Initialize("your-unique-request-id")

    // Configure access token
    err := olaMap.ConfigureAccessToken("your-client-id", "your-client-secret")
    if err != nil {
        fmt.Println("Error configuring access token:", err)
        return
    }

    // Get directions
    directions, err := olaMap.GetDirections("origin", "destination")
    if err != nil {
        fmt.Println("Error getting directions:", err)
        return
    }
    fmt.Println("Directions:", directions)
}
```

## Available Methods

The following methods are available for use with the `OLAMap` struct:

- **`Initialize(requestID string) *OLAMap`**: Initializes a new OLA Map instance with a unique request ID.
- **`ConfigureAccessToken(clientID, clientSecret string) error`**: Configures the OLA access token using client credentials.
- **`GetDirections(origin, destination string) (Directions, error)`**: Retrieves directions from the origin to the destination.
- **`PlaceAutoComplete(input string) (Places, error)`**: Provides place suggestions based on the input.
- **`GeoCode(address, bounds, language string) (GeoData, error)`**: Converts an address into geographic coordinates.
- **`ReverseGeocode(latlng string) (Address, error)`**: Converts geographic coordinates back into an address.
- **`GetDistanceMatrix(origins, destinations string) (DistanceMatrix, error)`**: Calculates distances between multiple origins and destinations.
