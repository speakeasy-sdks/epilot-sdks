package operations

import (
	"epilot-workflows-definition/sdk/pkg/models/shared"
)

type UpdateDefinitionPathParams struct {
	DefinitionID string `pathParam:"style=simple,explode=false,name=definitionId"`
}

type UpdateDefinitionRequest struct {
	PathParams UpdateDefinitionPathParams
	Request    shared.WorkflowDefinition `request:"mediaType=application/json"`
}

type UpdateDefinitionResponse struct {
	ContentType        string
	ErrorResp          *shared.ErrorResp
	StatusCode         int64
	WorkflowDefinition *shared.WorkflowDefinition
}
