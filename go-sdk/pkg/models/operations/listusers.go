package operations

import (
	"openapi/pkg/models/shared"
)

type ListUsersQueryParams struct {
	Limit  *float64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *float64 `queryParam:"style=form,explode=true,name=offset"`
	OrgIds []string `queryParam:"style=form,explode=false,name=org_ids"`
	Query  *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListUsers200ApplicationJSON struct {
	Users []shared.User `json:"users,omitempty"`
}

type ListUsersRequest struct {
	QueryParams ListUsersQueryParams
}

type ListUsersResponse struct {
	ContentType                       string
	StatusCode                        int64
	ListUsers200ApplicationJSONObject *ListUsers200ApplicationJSON
}
