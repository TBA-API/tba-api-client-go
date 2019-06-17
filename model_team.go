/*
 * The Blue Alliance API v3
 *
 * # Overview    Information and statistics about FIRST Robotics Competition teams and events.   # Authentication   All endpoints require an Auth Key to be passed in the header `X-TBA-Auth-Key`. If you do not have an auth key yet, you can obtain one from your [Account Page](/account).    A `User-Agent` header may need to be set to prevent a 403 Unauthorized error.
 *
 * API version: 3.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type Team struct {
	// TBA team key with the format `frcXXXX` with `XXXX` representing the team number.
	Key string `json:"key"`
	// Official team number issued by FIRST.
	TeamNumber int32 `json:"team_number"`
	// Team nickname provided by FIRST.
	Nickname string `json:"nickname,omitempty"`
	// Official long name registered with FIRST.
	Name string `json:"name"`
	// City of team derived from parsing the address registered with FIRST.
	City string `json:"city,omitempty"`
	// State of team derived from parsing the address registered with FIRST.
	StateProv string `json:"state_prov,omitempty"`
	// Country of team derived from parsing the address registered with FIRST.
	Country string `json:"country,omitempty"`
	// Will be NULL, for future development.
	Address string `json:"address,omitempty"`
	// Postal code from the team address.
	PostalCode string `json:"postal_code,omitempty"`
	// Will be NULL, for future development.
	GmapsPlaceId string `json:"gmaps_place_id,omitempty"`
	// Will be NULL, for future development.
	GmapsUrl string `json:"gmaps_url,omitempty"`
	// Will be NULL, for future development.
	Lat float64 `json:"lat,omitempty"`
	// Will be NULL, for future development.
	Lng float64 `json:"lng,omitempty"`
	// Will be NULL, for future development.
	LocationName string `json:"location_name,omitempty"`
	// Official website associated with the team.
	Website string `json:"website,omitempty"`
	// First year the team officially competed.
	RookieYear int32 `json:"rookie_year"`
	// Team's motto as provided by FIRST. This field is deprecated and will return null - will be removed at end-of-season in 2019.
	Motto string `json:"motto,omitempty"`
	// Location of the team's home championship each year as a key-value pair. The year (as a string) is the key, and the city is the value.
	HomeChampionship map[string]interface{} `json:"home_championship,omitempty"`
}
