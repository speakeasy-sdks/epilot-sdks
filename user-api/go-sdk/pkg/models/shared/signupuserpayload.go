package shared

type SignupUserPayloadLanguageEnum string

const (
	SignupUserPayloadLanguageEnumEn SignupUserPayloadLanguageEnum = "en"
	SignupUserPayloadLanguageEnumDe SignupUserPayloadLanguageEnum = "de"
)

type SignupUserPayload struct {
	Language           *SignupUserPayloadLanguageEnum `json:"language,omitempty"`
	OrganizationDetail map[string]interface{}         `json:"organization_detail,omitempty"`
	UserDetail         *UserDetail                    `json:"user_detail,omitempty"`
}
