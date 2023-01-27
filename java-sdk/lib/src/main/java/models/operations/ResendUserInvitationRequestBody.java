package .models.operations;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonInclude.Include;
public class ResendUserInvitationRequestBody {
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("email")
    public String email;
    public ResendUserInvitationRequestBody withEmail(String email) {
        this.email = email;
        return this;
    }
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("language")
    public ResendUserInvitationRequestBodyLanguageEnum language;
    public ResendUserInvitationRequestBody withLanguage(ResendUserInvitationRequestBodyLanguageEnum language) {
        this.language = language;
        return this;
    }
}