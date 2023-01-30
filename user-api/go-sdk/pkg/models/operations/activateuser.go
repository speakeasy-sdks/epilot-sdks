package operations

import (
	"openapi/pkg/models/shared"
)

type ActivateUserQueryParams struct {
	Token string `queryParam:"style=form,explode=true,name=token"`
}

type ActivateUserRequest struct {
	QueryParams ActivateUserQueryParams
	Request     *shared.UserActivationPayload `request:"mediaType=application/json"`
}

type ActivateUserResponse struct {
	ContentType string
	StatusCode  int64
}
