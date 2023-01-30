package shared

type UserProperties struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type User struct {
	DisplayName       *string                `json:"display_name,omitempty"`
	Email             string                 `json:"email"`
	ID                string                 `json:"id"`
	ImageURI          map[string]interface{} `json:"image_uri,omitempty"`
	Name              string                 `json:"name"`
	OrganizationID    string                 `json:"organization_id"`
	PreferredLanguage string                 `json:"preferred_language"`
	Properties        []UserProperties       `json:"properties"`
	Roles             []string               `json:"roles"`
	Signature         *string                `json:"signature,omitempty"`
}
