package operations

import (
	"epilot-workflows-definition/sdk/pkg/models/shared"
)

type GetAllClosingReasonsQueryParams struct {
	IncludeInactive *bool `queryParam:"style=form,explode=true,name=includeInactive"`
}

type GetAllClosingReasonsRequest struct {
	QueryParams GetAllClosingReasonsQueryParams
}

type GetAllClosingReasonsResponse struct {
	ClosingReasons *shared.ClosingReasons
	ContentType    string
	StatusCode     int64
}
