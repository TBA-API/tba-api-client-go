/*
 * The Blue Alliance API v3
 *
 * # Overview    Information and statistics about FIRST Robotics Competition teams and events.   # Authentication   All endpoints require an Auth Key to be passed in the header `X-TBA-Auth-Key`. If you do not have an auth key yet, you can obtain one from your [Account Page](/account).    A `User-Agent` header may need to be set to prevent a 403 Unauthorized error.
 *
 * API version: 3.03.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ApiStatusAppVersion struct {

	// Internal use - Minimum application version required to correctly connect and process data.
	MinAppVersion int32 `json:"min_app_version"`

	// Internal use - Latest application version available.
	LatestAppVersion int32 `json:"latest_app_version"`
}
