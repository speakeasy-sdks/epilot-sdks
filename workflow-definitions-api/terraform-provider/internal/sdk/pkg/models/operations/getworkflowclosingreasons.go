package operations

import (
	"epilot-workflows-definition/sdk/pkg/models/shared"
)

type GetWorkflowClosingReasonsPathParams struct {
	DefinitionID string `pathParam:"style=simple,explode=false,name=definitionId"`
}

type GetWorkflowClosingReasonsRequest struct {
	PathParams GetWorkflowClosingReasonsPathParams
}

type GetWorkflowClosingReasonsResponse struct {
	ClosingReasonsIds *shared.ClosingReasonsIds
	ContentType       string
	StatusCode        int64
}
