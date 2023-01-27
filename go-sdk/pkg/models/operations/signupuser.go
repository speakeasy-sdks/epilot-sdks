package operations

import (
	"openapi/pkg/models/shared"
)

type SignUpUserQueryParams struct {
	Token *string `queryParam:"style=form,explode=true,name=token"`
}

type SignUpUser200ApplicationJSON struct {
	Organization *shared.Organization `json:"organization,omitempty"`
	User         *shared.User         `json:"user,omitempty"`
}

type SignUpUserRequest struct {
	QueryParams SignUpUserQueryParams
	Request     *shared.SignupUserPayload `request:"mediaType=application/json"`
}

type SignUpUserResponse struct {
	ContentType                        string
	StatusCode                         int64
	SignUpUser200ApplicationJSONObject *SignUpUser200ApplicationJSON
}
