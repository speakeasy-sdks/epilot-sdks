package shared

type UserV2Properties struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type UserV2StatusEnum string

const (
	UserV2StatusEnumActive      UserV2StatusEnum = "Active"
	UserV2StatusEnumPending     UserV2StatusEnum = "Pending"
	UserV2StatusEnumDeactivated UserV2StatusEnum = "Deactivated"
	UserV2StatusEnumDeleted     UserV2StatusEnum = "Deleted"
)

type UserV2 struct {
	CreatedAt          *string                `json:"created_at,omitempty"`
	Department         *string                `json:"department,omitempty"`
	DisplayName        *string                `json:"display_name,omitempty"`
	DraftEmail         *string                `json:"draft_email,omitempty"`
	Email              *string                `json:"email,omitempty"`
	ID                 *string                `json:"id,omitempty"`
	ImageURI           map[string]interface{} `json:"image_uri,omitempty"`
	IsSignatureEnabled *bool                  `json:"is_signature_enabled,omitempty"`
	MfaEnabled         *bool                  `json:"mfa_enabled,omitempty"`
	OrganizationID     *string                `json:"organization_id,omitempty"`
	Phone              *string                `json:"phone,omitempty"`
	PhoneVerified      *bool                  `json:"phone_verified,omitempty"`
	PreferredLanguage  *string                `json:"preferred_language,omitempty"`
	Properties         []UserV2Properties     `json:"properties,omitempty"`
	Signature          *string                `json:"signature,omitempty"`
	Status             *UserV2StatusEnum      `json:"status,omitempty"`
	Token              *string                `json:"token,omitempty"`
}
