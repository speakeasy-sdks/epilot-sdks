package operations

import (
	"epilot-workflows-definition/sdk/pkg/models/shared"
)

type CreateDefinitionRequest struct {
	Request shared.WorkflowDefinition `request:"mediaType=application/json"`
}

type CreateDefinitionResponse struct {
	ContentType        string
	ErrorResp          *shared.ErrorResp
	StatusCode         int64
	WorkflowDefinition *shared.WorkflowDefinition
}
