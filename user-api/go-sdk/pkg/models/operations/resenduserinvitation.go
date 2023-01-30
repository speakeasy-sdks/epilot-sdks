package operations

import (
	"openapi/pkg/models/shared"
)

type ResendUserInvitationRequestBodyLanguageEnum string

const (
	ResendUserInvitationRequestBodyLanguageEnumEn ResendUserInvitationRequestBodyLanguageEnum = "en"
	ResendUserInvitationRequestBodyLanguageEnumDe ResendUserInvitationRequestBodyLanguageEnum = "de"
)

type ResendUserInvitationRequestBody struct {
	Email    *string                                      `json:"email,omitempty"`
	Language *ResendUserInvitationRequestBodyLanguageEnum `json:"language,omitempty"`
}

type ResendUserInvitationRequest struct {
	Request *ResendUserInvitationRequestBody `request:"mediaType=application/json"`
}

type ResendUserInvitationResponse struct {
	ContentType string
	StatusCode  int64
	UserV2      *shared.UserV2
}
