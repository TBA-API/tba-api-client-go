/*
 * The Blue Alliance API v3
 *
 * # Overview    Information and statistics about FIRST Robotics Competition teams and events.   # Authentication   All endpoints require an Auth Key to be passed in the header `X-TBA-Auth-Key`. If you do not have an auth key yet, you can obtain one from your [Account Page](/account).    A `User-Agent` header may need to be set to prevent a 403 Unauthorized error.
 *
 * API version: 3.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type EventSimple struct {
	// TBA event key with the format yyyy[EVENT_CODE], where yyyy is the year, and EVENT_CODE is the event code of the event.
	Key string `json:"key"`
	// Official name of event on record either provided by FIRST or organizers of offseason event.
	Name string `json:"name"`
	// Event short code, as provided by FIRST.
	EventCode string `json:"event_code"`
	// Event Type, as defined here: https://github.com/the-blue-alliance/the-blue-alliance/blob/master/consts/event_type.py#L2
	EventType int32 `json:"event_type"`
	District DistrictList `json:"district,omitempty"`
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
}
