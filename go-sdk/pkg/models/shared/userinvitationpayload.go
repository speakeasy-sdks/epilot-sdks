package shared

type UserInvitationPayloadLanguageEnum string

const (
	UserInvitationPayloadLanguageEnumEn UserInvitationPayloadLanguageEnum = "en"
	UserInvitationPayloadLanguageEnumDe UserInvitationPayloadLanguageEnum = "de"
)

type UserInvitationPayload struct {
	Email    *string                            `json:"email,omitempty"`
	Language *UserInvitationPayloadLanguageEnum `json:"language,omitempty"`
	Roles    []string                           `json:"roles,omitempty"`
}
