package .models.shared;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonInclude.Include;
public class User {
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("display_name")
    public String displayName;
    public User withDisplayName(String displayName) {
        this.displayName = displayName;
        return this;
    }
    @JsonProperty("email")
    public String email;
    public User withEmail(String email) {
        this.email = email;
        return this;
    }
    @JsonProperty("id")
    public String id;
    public User withId(String id) {
        this.id = id;
        return this;
    }
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("image_uri")
    public java.util.Map<String, Object> imageUri;
    public User withImageUri(java.util.Map<String, Object> imageUri) {
        this.imageUri = imageUri;
        return this;
    }
    @JsonProperty("name")
    public String name;
    public User withName(String name) {
        this.name = name;
        return this;
    }
    @JsonProperty("organization_id")
    public String organizationId;
    public User withOrganizationId(String organizationId) {
        this.organizationId = organizationId;
        return this;
    }
    @JsonProperty("preferred_language")
    public String preferredLanguage;
    public User withPreferredLanguage(String preferredLanguage) {
        this.preferredLanguage = preferredLanguage;
        return this;
    }
    @JsonProperty("properties")
    public UserProperties[] properties;
    public User withProperties(UserProperties[] properties) {
        this.properties = properties;
        return this;
    }
    @JsonProperty("roles")
    public String[] roles;
    public User withRoles(String[] roles) {
        this.roles = roles;
        return this;
    }
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("signature")
    public String signature;
    public User withSignature(String signature) {
        this.signature = signature;
        return this;
    }
}