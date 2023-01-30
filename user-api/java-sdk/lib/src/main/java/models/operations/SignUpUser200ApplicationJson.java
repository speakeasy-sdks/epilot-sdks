package .models.operations;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonInclude.Include;
public class SignUpUser200ApplicationJson {
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("organization")
    public .models.shared.Organization organization;
    public SignUpUser200ApplicationJson withOrganization(.models.shared.Organization organization) {
        this.organization = organization;
        return this;
    }
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("user")
    public .models.shared.User user;
    public SignUpUser200ApplicationJson withUser(.models.shared.User user) {
        this.user = user;
        return this;
    }
}