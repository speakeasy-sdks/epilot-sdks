package operations

import (
	"openapi/pkg/models/shared"
)

type UpdateUserV2PathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateUserV2Request struct {
	PathParams UpdateUserV2PathParams
	Request    *shared.UserV2 `request:"mediaType=application/json"`
}

type UpdateUserV2Response struct {
	ContentType string
	StatusCode  int64
	UserV2      *shared.UserV2
}
