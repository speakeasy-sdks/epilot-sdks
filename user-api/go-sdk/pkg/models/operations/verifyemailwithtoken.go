package operations

import (
	"openapi/pkg/models/shared"
)

type VerifyEmailWithTokenQueryParams struct {
	Token string `queryParam:"style=form,explode=true,name=token"`
}

type VerifyEmailWithTokenRequest struct {
	QueryParams VerifyEmailWithTokenQueryParams
	Request     *shared.UserVerificationPayload `request:"mediaType=application/json"`
}

type VerifyEmailWithTokenResponse struct {
	ContentType string
	StatusCode  int64
}
