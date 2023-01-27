package .models.shared;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonInclude.Include;
public class SignupUserPayload {
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("language")
    public SignupUserPayloadLanguageEnum language;
    public SignupUserPayload withLanguage(SignupUserPayloadLanguageEnum language) {
        this.language = language;
        return this;
    }
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("organization_detail")
    public java.util.Map<String, Object> organizationDetail;
    public SignupUserPayload withOrganizationDetail(java.util.Map<String, Object> organizationDetail) {
        this.organizationDetail = organizationDetail;
        return this;
    }
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("user_detail")
    public UserDetail userDetail;
    public SignupUserPayload withUserDetail(UserDetail userDetail) {
        this.userDetail = userDetail;
        return this;
    }
}