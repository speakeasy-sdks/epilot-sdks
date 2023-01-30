package operations

import (
	"epilot-workflows-definition/sdk/pkg/models/shared"
)

type GetDefinitionPathParams struct {
	DefinitionID string `pathParam:"style=simple,explode=false,name=definitionId"`
}

type GetDefinitionRequest struct {
	PathParams GetDefinitionPathParams
}

type GetDefinitionResponse struct {
	ContentType            string
	DefinitionNotFoundResp *interface{}
	ErrorResp              *shared.ErrorResp
	StatusCode             int64
	WorkflowDefinition     *shared.WorkflowDefinition
}
