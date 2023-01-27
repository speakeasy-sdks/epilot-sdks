package operations

import (
	"openapi/pkg/models/shared"
)

type GetMeV2Response struct {
	ContentType string
	StatusCode  int64
	UserV2      *shared.UserV2
}
