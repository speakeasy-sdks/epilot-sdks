package operations

import (
	"epilot-workflows-definition/sdk/pkg/models/shared"
)

type GetDefinitionsResponse struct {
	ContentType         string
	ErrorResp           *shared.ErrorResp
	StatusCode          int64
	WorkflowDefinitions []shared.WorkflowDefinition
}
