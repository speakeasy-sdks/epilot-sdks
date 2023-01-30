package operations

import (
	"epilot-workflows-definition/sdk/pkg/models/shared"
)

type GetMaxAllowedLimitResponse struct {
	ContentType     string
	ErrorResp       *shared.ErrorResp
	MaxAllowedLimit *shared.MaxAllowedLimit
	StatusCode      int64
}
