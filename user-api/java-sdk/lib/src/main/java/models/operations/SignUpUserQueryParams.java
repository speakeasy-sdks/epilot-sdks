package .models.operations;

import .utils.SpeakeasyMetadata;
public class SignUpUserQueryParams {
    @SpeakeasyMetadata("queryParam:style=form,explode=true,name=token")
    public String token;
    public SignUpUserQueryParams withToken(String token) {
        this.token = token;
        return this;
    }
}