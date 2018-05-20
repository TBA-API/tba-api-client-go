/*
 * The Blue Alliance API v3
 *
 * # Overview    Information and statistics about FIRST Robotics Competition teams and events.   # Authentication   All endpoints require an Auth Key to be passed in the header `X-TBA-Auth-Key`. If you do not have an auth key yet, you can obtain one from your [Account Page](/account).    A `User-Agent` header may need to be set to prevent a 403 Unauthorized error.
 *
 * API version: 3.03.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// An `Award_Recipient` object represents the team and/or person who received an award at an event.
type AwardRecipient struct {

	// The TBA team key for the team that was given the award. May be null.
	TeamKey string `json:"team_key,omitempty"`

	// The name of the individual given the award. May be null.
	Awardee string `json:"awardee,omitempty"`
}
