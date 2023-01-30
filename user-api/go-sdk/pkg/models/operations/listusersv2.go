package operations

import (
	"openapi/pkg/models/shared"
)

type ListUsersV2QueryParams struct {
	Limit  *float64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *float64 `queryParam:"style=form,explode=true,name=offset"`
	Query  *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListUsersV2200ApplicationJSON struct {
	Results []shared.UserV2 `json:"results,omitempty"`
}

type ListUsersV2Request struct {
	QueryParams ListUsersV2QueryParams
}

type ListUsersV2Response struct {
	ContentType                         string
	StatusCode                          int64
	ListUsersV2200ApplicationJSONObject *ListUsersV2200ApplicationJSON
}
