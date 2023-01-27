package shared

type UserActivationPayload struct {
	DisplayName *string `json:"display_name,omitempty"`
	Password    *string `json:"password,omitempty"`
}
