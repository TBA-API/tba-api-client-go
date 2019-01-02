/*
 * The Blue Alliance API v3
 *
 * # Overview    Information and statistics about FIRST Robotics Competition teams and events.   # Authentication   All endpoints require an Auth Key to be passed in the header `X-TBA-Auth-Key`. If you do not have an auth key yet, you can obtain one from your [Account Page](/account).    A `User-Agent` header may need to be set to prevent a 403 Unauthorized error.
 *
 * API version: 3.04.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// The `Media` object contains a reference for most any media associated with a team or event on TBA.
type Media struct {

	// TBA identifier for this media.
	Key string `json:"key"`

	// String type of the media element.
	Type_ string `json:"type"`

	// The key used to identify this media on the media site.
	ForeignKey string `json:"foreign_key,omitempty"`

	// If required, a JSON dict of additional media information.
	Details *interface{} `json:"details,omitempty"`

	// True if the media is of high quality.
	Preferred bool `json:"preferred,omitempty"`

	// Direct URL to the media.
	DirectUrl string `json:"direct_url,omitempty"`

	// The URL that leads to the full web page for the media, if one exists.
	ViewUrl string `json:"view_url,omitempty"`
}
