/*
 * The Blue Alliance API v3
 *
 * # Overview    Information and statistics about FIRST Robotics Competition teams and events.   # Authentication   All endpoints require an Auth Key to be passed in the header `X-TBA-Auth-Key`. If you do not have an auth key yet, you can obtain one from your [Account Page](/account).    A `User-Agent` header may need to be set to prevent a 403 Unauthorized error.
 *
 * API version: 3.04.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Event struct {

	// TBA event key with the format yyyy[EVENT_CODE], where yyyy is the year, and EVENT_CODE is the event code of the event.
	Key string `json:"key"`

	// Official name of event on record either provided by FIRST or organizers of offseason event.
	Name string `json:"name"`

	// Event short code, as provided by FIRST.
	EventCode string `json:"event_code"`

	// Event Type, as defined here: https://github.com/the-blue-alliance/the-blue-alliance/blob/master/consts/event_type.py#L2
	EventType int32 `json:"event_type"`

	// The district this event is in, may be null.
	District *DistrictList `json:"district,omitempty"`

	// City, town, village, etc. the event is located in.
	City string `json:"city,omitempty"`

	// State or Province the event is located in.
	StateProv string `json:"state_prov,omitempty"`

	// Country the event is located in.
	Country string `json:"country,omitempty"`

	// Event start date in `yyyy-mm-dd` format.
	StartDate string `json:"start_date"`

	// Event end date in `yyyy-mm-dd` format.
	EndDate string `json:"end_date"`

	// Year the event data is for.
	Year int32 `json:"year"`

	// Same as `name` but doesn't include event specifiers, such as 'Regional' or 'District'. May be null.
	ShortName string `json:"short_name,omitempty"`

	// Event Type, eg Regional, District, or Offseason.
	EventTypeString string `json:"event_type_string"`

	// Week of the event relative to the first official season event, zero-indexed. Only valid for Regionals, Districts, and District Championships. Null otherwise. (Eg. A season with a week 0 'preseason' event does not count, and week 1 events will show 0 here. Seasons with a week 0.5 regional event will show week 0 for those event(s) and week 1 for week 1 events and so on.)
	Week int32 `json:"week,omitempty"`

	// Address of the event's venue, if available.
	Address string `json:"address,omitempty"`

	// Postal code from the event address.
	PostalCode string `json:"postal_code,omitempty"`

	// Google Maps Place ID for the event address.
	GmapsPlaceId string `json:"gmaps_place_id,omitempty"`

	// Link to address location on Google Maps.
	GmapsUrl string `json:"gmaps_url,omitempty"`

	// Latitude for the event address.
	Lat float64 `json:"lat,omitempty"`

	// Longitude for the event address.
	Lng float64 `json:"lng,omitempty"`

	// Name of the location at the address for the event, eg. Blue Alliance High School.
	LocationName string `json:"location_name,omitempty"`

	// Timezone name.
	Timezone string `json:"timezone,omitempty"`

	// The event's website, if any.
	Website string `json:"website,omitempty"`

	// The FIRST internal Event ID, used to link to the event on the FRC webpage.
	FirstEventId string `json:"first_event_id,omitempty"`

	// Public facing event code used by FIRST (on frc-events.firstinspires.org, for example)
	FirstEventCode string `json:"first_event_code,omitempty"`

	Webcasts []Webcast `json:"webcasts,omitempty"`

	// An array of event keys for the divisions at this event.
	DivisionKeys []string `json:"division_keys,omitempty"`

	// The TBA Event key that represents the event's parent. Used to link back to the event from a division event. It is also the inverse relation of `divison_keys`.
	ParentEventKey string `json:"parent_event_key,omitempty"`

	// Playoff Type, as defined here: https://github.com/the-blue-alliance/the-blue-alliance/blob/master/consts/playoff_type.py#L4, or null.
	PlayoffType int32 `json:"playoff_type,omitempty"`

	// String representation of the `playoff_type`, or null.
	PlayoffTypeString string `json:"playoff_type_string,omitempty"`
}
