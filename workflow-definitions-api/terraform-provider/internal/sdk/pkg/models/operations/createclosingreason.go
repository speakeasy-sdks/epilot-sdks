package operations

import (
	"epilot-workflows-definition/sdk/pkg/models/shared"
)

type CreateClosingReasonRequest struct {
	Request shared.ClosingReason `request:"mediaType=application/json"`
}

type CreateClosingReasonResponse struct {
	ClosingReason *shared.ClosingReason
	ContentType   string
	StatusCode    int64
}
