package operations

import (
	"openapi/pkg/models/shared"
)

type GetUserLoginParametersV2PathParams struct {
	Username string `pathParam:"style=simple,explode=false,name=username"`
}

type GetUserLoginParametersV2200ApplicationJSON struct {
	LoginParameters []shared.LoginParameters `json:"login_parameters,omitempty"`
}

type GetUserLoginParametersV2Request struct {
	PathParams GetUserLoginParametersV2PathParams
}

type GetUserLoginParametersV2Response struct {
	ContentType                                      string
	StatusCode                                       int64
	GetUserLoginParametersV2200ApplicationJSONObject *GetUserLoginParametersV2200ApplicationJSON
}
