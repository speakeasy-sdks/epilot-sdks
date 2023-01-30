package operations

import (
	"openapi/pkg/models/shared"
)

type DeleteUserV2PathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type DeleteUserV2Request struct {
	PathParams DeleteUserV2PathParams
}

type DeleteUserV2Response struct {
	ContentType string
	StatusCode  int64
	User        *shared.User
}
