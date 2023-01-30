package operations

import (
	"openapi/pkg/models/shared"
)

type GetUserPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetUserRequest struct {
	PathParams GetUserPathParams
}

type GetUserResponse struct {
	ContentType string
	StatusCode  int64
	User        *shared.User
}
