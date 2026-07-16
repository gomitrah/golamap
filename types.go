package golamap

type NearBySearch struct {
	Layers       string
	Location     string
	Types        string
	Radius       string
	Strictbounds string
	WithCentroid string
	Limit        string
}

type MapImage struct {
	Stylename   string
	Imagewidth  string
	Imageheight string
	Imageformat string
	Path        string
	Markers     []string
}

type MapImageBounded struct {
	Stylename   string
	Minxstr     string
	Minystr     string
	Maxxstr     string
	Maxystr     string
	Imagewidth  string
	Imageheight string
	Imageformat string
	Markers     []string
	Path        string
}

type MapImageCenter struct {
	Stylename   string
	Longitude   string
	Latitude    string
	Zoomlevel   string
	Imagewidth  string
	Imageheight string
	Imageformat string
	Markers     []string
	Path        string
}

type TextSearch struct {
	Input    string
	Location string
	Radius   string
	Types    string
	Size     string
}

type Directions struct {
	GeocodedWaypoints []GeocodedWaypoint `json:"geocoded_waypoints"`
	Routes            []Route            `json:"routes"`
	Status            string             `json:"status"`
	SourceFrom        string             `json:"source_from"`
}

type GeocodedWaypoint struct {
	GeocoderStatus string   `json:"geocoder_status"`
	PlaceID        string   `json:"place_id"`
	Types          []string `json:"types"`
}

// Step represents a single step in the route.
type Step struct {
	Distance         int      `json:"distance"`
	ReadableDistance string   `json:"readable_distance"`
	Duration         int      `json:"duration"`
	ReadableDuration string   `json:"readable_duration"`
	StartLocation    Location `json:"start_location"`
	EndLocation      Location `json:"end_location"`
	Instructions     string   `json:"instructions"`
	Maneuver         string   `json:"maneuver"`
	BearingBefore    int      `json:"bearing_before"`
	BearingAfter     int      `json:"bearing_after"`
}

// Location represents geographical coordinates.
type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Leg struct {
	Distance         int      `json:"distance"`
	ReadableDistance string   `json:"readable_distance"`
	Duration         int      `json:"duration"`
	ReadableDuration string   `json:"readable_duration"`
	EndAddress       string   `json:"end_address"`
	EndLocation      Location `json:"end_location"`
	StartAddress     string   `json:"start_address"`
	StartLocation    Location `json:"start_location"`
	Steps            []Step   `json:"steps"`
}

type Route struct {
	Bounds           Bounds   `json:"bounds"`
	Copyrights       string   `json:"copyrights"`
	Legs             []Leg    `json:"legs"`
	OverviewPolyline string   `json:"overview_polyline"`
	TravelAdvisory   string   `json:"travel_advisory"`
	Summary          string   `json:"summary"`
	Warnings         []string `json:"warnings"`
	WaypointOrder    []int    `json:"waypoint_order"`
}

// Bounds represents the geographical bounds of the route.
type Bounds struct {
	Southwest Location `json:"southwest"`
	Northeast Location `json:"northeast"`
}

type DistanceMatrix struct {
	Rows   []Row  `json:"rows"`
	Status string `json:"status"`
}

type Row struct {
	Elements []Element `json:"elements"`
}

type Element struct {
	Duration int    `json:"duration"`
	Distance int    `json:"distance"`
	Polyline string `json:"polyline"`
	Status   string `json:"status"`
}

type AutoComplete struct {
	Predictions  []Prediction `json:"predictions"`
	InfoMessages []string     `json:"info_messages"`
	ErrorMessage string       `json:"error_message"`
	Status       string       `json:"status"`
}

type Prediction struct {
	Reference            string                           `json:"reference"`
	Types                []string                         `json:"types"`
	MatchedSubstrings    []MatchedSubstring               `json:"matched_substrings"`
	DistanceMeters       int                              `json:"distance_meters"`
	Terms                []Term                           `json:"terms"`
	StructuredFormatting AutoCompleteStructuredFormatting `json:"structured_formatting"`
	Description          string                           `json:"description"`
	Geometry             Geometry                         `json:"geometry"`
	PlaceID              string                           `json:"place_id"`
}

type MatchedSubstring struct {
	Offset int `json:"offset"`
	Length int `json:"length"`
}

type Term struct {
	Offset int    `json:"offset"`
	Value  string `json:"value"`
}

type AutoCompleteStructuredFormatting struct {
	MainText                  string             `json:"main_text"`
	SecondaryText             string             `json:"secondary_text"`
	MainTextMatchedSubstrings []MatchedSubstring `json:"main_text_matched_substrings"`
}

type Geometry struct {
	Location Location `json:"location"`
}

type AddressComponent struct {
	LongName  string   `json:"long_name"`
	ShortName string   `json:"short_name"`
	Types     []string `json:"types"`
}

type Viewport struct {
	Northeast Location `json:"northeast"`
	Southwest Location `json:"southwest"`
}

type PlusCode struct {
	CompoundCode string `json:"compound_code"`
	GlobalCode   string `json:"global_code"`
}

type ForwardGecode struct {
	Status           string            `json:"status"`
	GeocodingResults []GeocodingResult `json:"geocodingResults"`
}

type GeocodingResult struct {
	FormattedAddress  string             `json:"formatted_address"`
	Types             []string           `json:"types"`
	Name              string             `json:"name"`
	Geometry          GecodeGeometry     `json:"geometry"`
	AddressComponents []AddressComponent `json:"address_components"`
	PlusCode          PlusCode           `json:"plus_code"`
	PlaceID           string             `json:"place_id"`
	Layer             []string           `json:"layer"`
}

type GecodeGeometry struct {
	Location     Location `json:"location"`
	Viewport     Viewport `json:"viewport"`
	LocationType string   `json:"location_type"`
}

type ReverseGecode struct {
	ErrorMessage string             `json:"error_message"`
	InfoMessages []string           `json:"info_messages"`
	Results      []ReverseGeoResult `json:"results"`
	PlusCode     PlusCode           `json:"plus_code"`
	Status       string             `json:"status"`
}

type ReverseGeoResult struct {
	FormattedAddress  string             `json:"formatted_address"`
	Types             []string           `json:"types"`
	Name              string             `json:"name"`
	Geometry          ReverseGeometry    `json:"geometry"`
	AddressComponents []AddressComponent `json:"address_components"`
	PlusCode          PlusCode           `json:"plus_code"`
	PlaceID           string             `json:"place_id"`
	Layer             []string           `json:"layer"`
}

type ReverseGeometry struct {
	Location     Location `json:"location"`
	Viewport     Viewport `json:"viewport"`
	LocationType string   `json:"location_type"`
}
