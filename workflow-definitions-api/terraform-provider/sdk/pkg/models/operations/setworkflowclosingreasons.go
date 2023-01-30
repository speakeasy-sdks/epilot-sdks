package operations

import (
	"epilot-workflows-definition/sdk/pkg/models/shared"
)

type SetWorkflowClosingReasonsPathParams struct {
	DefinitionID string `pathParam:"style=simple,explode=false,name=definitionId"`
}

type SetWorkflowClosingReasonsRequest struct {
	PathParams SetWorkflowClosingReasonsPathParams
	Request    shared.ClosingReasonsIds `request:"mediaType=application/json"`
}

type SetWorkflowClosingReasonsResponse struct {
	ContentType string
	StatusCode  int64
}
