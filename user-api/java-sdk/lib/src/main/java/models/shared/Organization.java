package .models.shared;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonInclude.Include;
public class Organization {
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("address")
    public OrganizationAddress address;
    public Organization withAddress(OrganizationAddress address) {
        this.address = address;
        return this;
    }
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("email")
    public String email;
    public Organization withEmail(String email) {
        this.email = email;
        return this;
    }
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("id")
    public String id;
    public Organization withId(String id) {
        this.id = id;
        return this;
    }
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("is_unlicensed_org")
    public Boolean isUnlicensedOrg;
    public Organization withIsUnlicensedOrg(Boolean isUnlicensedOrg) {
        this.isUnlicensedOrg = isUnlicensedOrg;
        return this;
    }
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("logo_thumbnail_url")
    public String logoThumbnailUrl;
    public Organization withLogoThumbnailUrl(String logoThumbnailUrl) {
        this.logoThumbnailUrl = logoThumbnailUrl;
        return this;
    }
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("logo_url")
    public String logoUrl;
    public Organization withLogoUrl(String logoUrl) {
        this.logoUrl = logoUrl;
        return this;
    }
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("name")
    public String name;
    public Organization withName(String name) {
        this.name = name;
        return this;
    }
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("phone")
    public String phone;
    public Organization withPhone(String phone) {
        this.phone = phone;
        return this;
    }
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("pricing_tier")
    public String pricingTier;
    public Organization withPricingTier(String pricingTier) {
        this.pricingTier = pricingTier;
        return this;
    }
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("signature")
    public String signature;
    public Organization withSignature(String signature) {
        this.signature = signature;
        return this;
    }
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("symbol")
    public String symbol;
    public Organization withSymbol(String symbol) {
        this.symbol = symbol;
        return this;
    }
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("type")
    public OrganizationTypeEnum type;
    public Organization withType(OrganizationTypeEnum type) {
        this.type = type;
        return this;
    }
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("website")
    public String website;
    public Organization withWebsite(String website) {
        this.website = website;
        return this;
    }
}