package .models.operations;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonInclude.Include;
public class GetUserLoginParametersV2200ApplicationJson {
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("login_parameters")
    public .models.shared.LoginParameters[] loginParameters;
    public GetUserLoginParametersV2200ApplicationJson withLoginParameters(.models.shared.LoginParameters[] loginParameters) {
        this.loginParameters = loginParameters;
        return this;
    }
}