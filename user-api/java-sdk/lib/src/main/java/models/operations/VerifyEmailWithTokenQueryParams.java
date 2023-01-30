package .models.operations;

import .utils.SpeakeasyMetadata;
public class VerifyEmailWithTokenQueryParams {
    @SpeakeasyMetadata("queryParam:style=form,explode=true,name=token")
    public String token;
    public VerifyEmailWithTokenQueryParams withToken(String token) {
        this.token = token;
        return this;
    }
}