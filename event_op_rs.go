/*
 * The Blue Alliance API v3
 *
 * # Overview    Information and statistics about FIRST Robotics Competition teams and events.   # Authentication   All endpoints require an Auth Key to be passed in the header `X-TBA-Auth-Key`. If you do not have an auth key yet, you can obtain one from your [Account Page](/account).    A `User-Agent` header may need to be set to prevent a 403 Unauthorized error.
 *
 * API version: 3.03.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// OPR, DPR, and CCWM for teams at the event.
type EventOpRs struct {

	// A key-value pair with team key (eg `frc254`) as key and OPR as value.
	Oprs map[string]float32 `json:"oprs,omitempty"`

	// A key-value pair with team key (eg `frc254`) as key and DPR as value.
	Dprs map[string]float32 `json:"dprs,omitempty"`

	// A key-value pair with team key (eg `frc254`) as key and CCWM as value.
	Ccwms map[string]float32 `json:"ccwms,omitempty"`
}
