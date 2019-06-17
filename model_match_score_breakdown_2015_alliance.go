/*
 * The Blue Alliance API v3
 *
 * # Overview    Information and statistics about FIRST Robotics Competition teams and events.   # Authentication   All endpoints require an Auth Key to be passed in the header `X-TBA-Auth-Key`. If you do not have an auth key yet, you can obtain one from your [Account Page](/account).    A `User-Agent` header may need to be set to prevent a 403 Unauthorized error.
 *
 * API version: 3.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type MatchScoreBreakdown2015Alliance struct {
	AutoPoints int32 `json:"auto_points,omitempty"`
	TeleopPoints int32 `json:"teleop_points,omitempty"`
	ContainerPoints int32 `json:"container_points,omitempty"`
	TotePoints int32 `json:"tote_points,omitempty"`
	LitterPoints int32 `json:"litter_points,omitempty"`
	FoulPoints int32 `json:"foul_points,omitempty"`
	AdjustPoints int32 `json:"adjust_points,omitempty"`
	TotalPoints int32 `json:"total_points,omitempty"`
	FoulCount int32 `json:"foul_count,omitempty"`
	ToteCountFar int32 `json:"tote_count_far,omitempty"`
	ToteCountNear int32 `json:"tote_count_near,omitempty"`
	ToteSet bool `json:"tote_set,omitempty"`
	ToteStack bool `json:"tote_stack,omitempty"`
	ContainerCountLevel1 int32 `json:"container_count_level1,omitempty"`
	ContainerCountLevel2 int32 `json:"container_count_level2,omitempty"`
	ContainerCountLevel3 int32 `json:"container_count_level3,omitempty"`
	ContainerCountLevel4 int32 `json:"container_count_level4,omitempty"`
	ContainerCountLevel5 int32 `json:"container_count_level5,omitempty"`
	ContainerCountLevel6 int32 `json:"container_count_level6,omitempty"`
	ContainerSet bool `json:"container_set,omitempty"`
	LitterCountContainer int32 `json:"litter_count_container,omitempty"`
	LitterCountLandfill int32 `json:"litter_count_landfill,omitempty"`
	LitterCountUnprocessed int32 `json:"litter_count_unprocessed,omitempty"`
	RobotSet bool `json:"robot_set,omitempty"`
}
