package operations

import (
	"epilot-workflows-definition/sdk/pkg/models/shared"
)

type ChangeReasonStatusPathParams struct {
	ReasonID string `pathParam:"style=simple,explode=false,name=reasonId"`
}

type ChangeReasonStatusRequest struct {
	PathParams ChangeReasonStatusPathParams
	Request    *shared.ChangeReasonStatusReq `request:"mediaType=application/json"`
}

type ChangeReasonStatusResponse struct {
	ContentType string
	ErrorResp   *shared.ErrorResp
	StatusCode  int64
}
