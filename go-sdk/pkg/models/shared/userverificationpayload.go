package shared

type UserVerificationPayload struct {
	Password *string `json:"password,omitempty"`
}
