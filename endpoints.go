package golamap

var (
	TokenURL             = "https://account.olamaps.io/realms/olamaps/protocol/openid-connect/token"
	DirectionsURL        = "https://api.olamaps.io/routing/v1/directions?origin=%s&destination=%s"
	PlaceAutoCompleteURL = "https://api.olamaps.io/places/v1/autocomplete?input=%s"
	GeoCodeURL           = "https://api.olamaps.io/places/v1/geocode?address=%s&bounds=%s&language=%s"
	ReverseGeocodeURL    = "https://api.olamaps.io/places/v1/reverse-geocode?latlng=%s"
	DistanceMatrixURL    = "https://api.olamaps.io/routing/v1/distanceMatrix?origins=%s&destinations=%s"
)
