package operations

import (
	"openapi/pkg/models/shared"
)

type InviteUserRequest struct {
	Request *shared.UserInvitationPayload `request:"mediaType=application/json"`
}

type InviteUserResponse struct {
	ContentType string
	StatusCode  int64
	UserV2      *shared.UserV2
}
