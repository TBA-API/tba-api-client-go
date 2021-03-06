/*
 * The Blue Alliance API v3
 *
 * # Overview    Information and statistics about FIRST Robotics Competition teams and events.   # Authentication   All endpoints require an Auth Key to be passed in the header `X-TBA-Auth-Key`. If you do not have an auth key yet, you can obtain one from your [Account Page](/account).    A `User-Agent` header may need to be set to prevent a 403 Unauthorized error.
 *
 * API version: 3.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// MatchScoreBreakdown2019 See the 2019 FMS API documentation for a description of each value. https://frcevents2.docs.apiary.io/#reference/match-results/score-details
type MatchScoreBreakdown2019 struct {
	Blue MatchScoreBreakdown2019Alliance `json:"blue,omitempty"`
	Red MatchScoreBreakdown2019Alliance `json:"red,omitempty"`
}
