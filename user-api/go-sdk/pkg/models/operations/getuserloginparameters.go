package operations

import (
	"openapi/pkg/models/shared"
)

type GetUserLoginParametersPathParams struct {
	Username string `pathParam:"style=simple,explode=false,name=username"`
}

type GetUserLoginParameters200ApplicationJSON struct {
	LoginParameters []shared.LoginParameters `json:"login_parameters,omitempty"`
}

type GetUserLoginParametersRequest struct {
	PathParams GetUserLoginParametersPathParams
}

type GetUserLoginParametersResponse struct {
	ContentType                                    string
	StatusCode                                     int64
	GetUserLoginParameters200ApplicationJSONObject *GetUserLoginParameters200ApplicationJSON
}
