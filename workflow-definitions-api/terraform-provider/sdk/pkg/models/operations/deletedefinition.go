package operations

import (
	"epilot-workflows-definition/sdk/pkg/models/shared"
)

type DeleteDefinitionPathParams struct {
	DefinitionID string `pathParam:"style=simple,explode=false,name=definitionId"`
}

type DeleteDefinitionRequest struct {
	PathParams DeleteDefinitionPathParams
}

type DeleteDefinitionResponse struct {
	ContentType string
	ErrorResp   *shared.ErrorResp
	StatusCode  int64
}
