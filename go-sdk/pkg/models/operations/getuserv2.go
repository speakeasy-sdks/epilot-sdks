package operations

import (
	"openapi/pkg/models/shared"
)

type GetUserV2PathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetUserV2Request struct {
	PathParams GetUserV2PathParams
}

type GetUserV2Response struct {
	ContentType string
	StatusCode  int64
	UserV2      *shared.UserV2
}
