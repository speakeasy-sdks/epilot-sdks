package shared

type LoginParametersOauthResponseTypeEnum string

const (
	LoginParametersOauthResponseTypeEnumCode  LoginParametersOauthResponseTypeEnum = "code"
	LoginParametersOauthResponseTypeEnumToken LoginParametersOauthResponseTypeEnum = "token"
)

type LoginParameters struct {
	CognitoIdentityPoolID   *string                               `json:"cognito_identity_pool_id,omitempty"`
	CognitoOauthDomain      *string                               `json:"cognito_oauth_domain,omitempty"`
	CognitoOauthScopes      []string                              `json:"cognito_oauth_scopes,omitempty"`
	CognitoRegion           *string                               `json:"cognito_region,omitempty"`
	CognitoUserPoolClientID *string                               `json:"cognito_user_pool_client_id,omitempty"`
	CognitoUserPoolID       *string                               `json:"cognito_user_pool_id,omitempty"`
	OauthResponseType       *LoginParametersOauthResponseTypeEnum `json:"oauth_response_type,omitempty"`
	OrganizationID          *string                               `json:"organization_id,omitempty"`
	OrganizationName        *string                               `json:"organization_name,omitempty"`
	OrganizationType        *string                               `json:"organization_type,omitempty"`
}
