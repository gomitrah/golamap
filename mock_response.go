package golamap

const (
	DirectionResponse = `
	{
	  "geocoded_waypoints": [
		{
		  "geocoder_status": "OK",
		  "place_id": "tdqcey303",
		  "types": []
		}
	  ],
	  "routes": [
		{
		  "bounds": "string",
		  "copyrights": "OLA Map data ©2024",
		  "legs": [
			{
			  "distance": 3836,
			  "readable_distance": "3 km 800 metres",
			  "duration": 698,
			  "readable_duration": "14 mins",
			  "end_address": "12.93397277.611342",
			  "end_location": {
				"lat": 12.933972,
				"lng": 77.611342
			  },
			  "start_address": "12.90934277.621689",
			  "start_location": {
				"lat": 12.909342,
				"lng": 77.62168
			  },
			  "steps": [
				{
				  "distance": 476,
				  "readable_distance": "0.5 km",
				  "duration": 99,
				  "readable_duration": "0 hours 2 min",
				  "start_location": {
					"lat": 12.90934,
					"lng": 77.62169
				  },
				  "end_location": {
					"lat": 12.908710000000001,
					"lng": 77.62158
				  },
				  "instructions": "Head west on NH48",
				  "maneuver": "turn-right",
				  "bearing_before": 0,
				  "bearing_after": 259
				}
			  ]
			}
		  ],
		  "travel_advisory": "0,1,0 | 1,3,15 | 3,4,10",
		  "summary": "",
		  "warnings": [],
		  "waypoint_order": [
			0,
			3,
			2,
			1
		  ]
		}
	  ],
	  "status": "OK",
	  "source_from": "Ola Maps"
	}`

	DistanceMatrixResponse = `{
		"rows": [
		  {
			"elements": [
			  {
				"duration": 719,
				"distance": 3224,
				"polyline": "su|mAyjvxM[GMCPmA?E@CJgADGDAxBTDJ[lCe@nDc@lDALAX^?d@WnAk@hBcAHEb@UVOJSHa@@CL_@Pc@FUDSJc@Lg@dBiGzFaTBIFUpBeGZaA|@iCpCiJ@CpBqGf@{AbByFx@_CFW@kBg@iFASUiBGu@y@kJAECMASKaAASCSSaCOaBUsCG}@E}@?YD]",
				"status": "OK"
			  },
			  {
				"duration": 1257,
				"distance": 5863,
				"status": "OK"
			  }
			]
		  }
		],
	  }`

	GeoCodeResponse = `{
		"status": "ok",
		"geocodingResults": [
		  {
			"formatted_address": "Mumbai, Mumbai Suburban, Mumbai Suburban District, Maharashtra, India",
			"types": [],
			"name": "Mumbai",
			"geometry": {
			  "viewport": {
				"southwest": {
				  "lng": 72.87232696702962,
				  "lat": 19.114096358602232
				},
				"northeast": {
				  "lng": 72.87498903297039,
				  "lat": 19.116611641397768
				}
			  },
			  "location": {
				"lng": 72.873658,
				"lat": 19.115354
			  },
			  "location_type": "geometric_center"
			},
			"address_components": [
			  {
				"types": [
				  "country"
				],
				"short_name": "India",
				"long_name": "India"
			  },
			  {
				"types": [
				  "administrative_area_level_1"
				],
				"short_name": "Maharashtra",
				"long_name": "Maharashtra"
			  },
			  {
				"types": [
				  "administrative_area_level_2"
				],
				"short_name": "Mumbai Suburban District",
				"long_name": "Mumbai Suburban District"
			  },
			  {
				"types": [
				  "administrative_area_level_3"
				],
				"short_name": "Mumbai Suburban",
				"long_name": "Mumbai Suburban"
			  },
			  {
				"types": [
				  "locality"
				],
				"short_name": "Mumbai",
				"long_name": "Mumbai"
			  },
			  {
				"types": [
				  "postal_code"
				],
				"short_name": "400059",
				"long_name": "400059"
			  },
			  {
				"types": [
				  "street_address"
				],
				"short_name": "Road Number 16",
				"long_name": "Road Number 16"
			  }
			],
			"plus_code": {
			  "compound_code": "NA",
			  "global_code": "NA"
			},
			"place_id": "ola-platform:1",
			"layer": [
			  "locality"
			]
		  }
		]
	  }`

	ReverseGeocodeResponse = `{
		"error_message": "",
		"info_messages": [],
		"results": [
		  {
			"formatted_address": "West Oidar Services, Central Tax, Bangalore, 560085, Phase 3, Banashankari, Bengaluru, Karnataka",
			"types": [
			  "point_of_interest",
			  "establishment"
			],
			"name": "West Oidar Services, Central Tax, Bangalore",
			"geometry": {
			  "viewport": {
				"southwest": {
				  "lng": 77.55141967245793,
				  "lat": 12.922552358602232
				},
				"northeast": {
				  "lng": 77.55400032754208,
				  "lat": 12.925067641397767
				}
			  },
			  "location": {
				"lng": 77.55271,
				"lat": 12.92381
			  },
			  "location_type": "rooftop"
			},
			"address_components": [
			  {
				"types": [
				  [
					"country"
				  ]
				],
				"short_name": "India",
				"long_name": "India"
			  }
			],
			"plus_code": {
			  "compound_code": "NA",
			  "global_code": "NA"
			},
			"place_id": "ola-platform:90be49b9e86d386a7076e25f71b402ca",
			"layer": [
			  "venue"
			]
		  }
		],
		"plus_code": {
		  "compound_code": "NA",
		  "global_code": "NA"
		},
		"status": "ok"
	  }`

	AutoCompleteResponse = `{
		"predictions": [
		  {
			"reference": "ChIJZWJEdf4crjsRjkEpoelwbCk",
			"types": [
			  "airport",
			  "point_of_interest",
			  "establishment"
			],
			"matched_substrings": [
			  {
				"offset": 0,
				"length": 5
			  }
			],
			"distance_meters": 31312,
			"terms": [
			  {
				"offset": 0,
				"value": "Kempegowda International Airport Bengaluru (BLR)"
			  }
			],
			"structured_formatting": {
			  "main_text_matched_substrings": [
				{
				  "offset": 0,
				  "length": 5
				}
			  ],
			  "secondary_text": "KIAL Rd, Devanahalli, Bengaluru, Karnataka, India",
			  "main_text": "Kempegowda International Airport Bengaluru (BLR)"
			},
			"description": "Kempegowda International Airport Bengaluru (BLR), KIAL Rd, Devanahalli, Bengaluru, Karnataka, India",
			"geometry": {
			  "location": {
				"lng": 77.61194,
				"lat": 13.05158
			  }
			},
			"place_id": "ola-platform:a1e99e3bd07b545284643f7bde183120"
		  }
		],
		"info_messages": [],
		"error_message": "",
		"status": "ok"
	  `
)
