/*
 * The Blue Alliance API v3
 *
 * # Overview    Information and statistics about FIRST Robotics Competition teams and events.   # Authentication   All endpoints require an Auth Key to be passed in the header `X-TBA-Auth-Key`. If you do not have an auth key yet, you can obtain one from your [Account Page](/account).    A `User-Agent` header may need to be set to prevent a 403 Unauthorized error.
 *
 * API version: 3.04.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type EventRankingRankings struct {

	// Number of times disqualified.
	Dq int32 `json:"dq"`

	// Number of matches played by this team.
	MatchesPlayed int32 `json:"matches_played"`

	// The average match score during qualifications. Year specific. May be null if not relevant for a given year.
	QualAverage int32 `json:"qual_average,omitempty"`

	// The team's rank at the event as provided by FIRST.
	Rank int32 `json:"rank"`

	Record *WltRecord `json:"record"`

	// Additional special data on the team's performance calculated by TBA.
	ExtraStats []float32 `json:"extra_stats,omitempty"`

	// Additional year-specific information, may be null. See parent `sort_order_info` for details.
	SortOrders []float32 `json:"sort_orders,omitempty"`

	// The team with this rank.
	TeamKey string `json:"team_key"`
}
