package shared

type OrganizationAddress struct {
	City         *string `json:"city,omitempty"`
	Country      *string `json:"country,omitempty"`
	PostalCode   *string `json:"postal_code,omitempty"`
	Street       *string `json:"street,omitempty"`
	StreetNumber *string `json:"street_number,omitempty"`
}

type OrganizationTypeEnum string

const (
	OrganizationTypeEnumVendor  OrganizationTypeEnum = "Vendor"
	OrganizationTypeEnumPartner OrganizationTypeEnum = "Partner"
)

type Organization struct {
	Address          *OrganizationAddress  `json:"address,omitempty"`
	Email            *string               `json:"email,omitempty"`
	ID               *string               `json:"id,omitempty"`
	IsUnlicensedOrg  *bool                 `json:"is_unlicensed_org,omitempty"`
	LogoThumbnailURL *string               `json:"logo_thumbnail_url,omitempty"`
	LogoURL          *string               `json:"logo_url,omitempty"`
	Name             *string               `json:"name,omitempty"`
	Phone            *string               `json:"phone,omitempty"`
	PricingTier      *string               `json:"pricing_tier,omitempty"`
	Signature        *string               `json:"signature,omitempty"`
	Symbol           *string               `json:"symbol,omitempty"`
	Type             *OrganizationTypeEnum `json:"type,omitempty"`
	Website          *string               `json:"website,omitempty"`
}
