package .models.operations;

import .utils.SpeakeasyMetadata;
public class ActivateUserQueryParams {
    @SpeakeasyMetadata("queryParam:style=form,explode=true,name=token")
    public String token;
    public ActivateUserQueryParams withToken(String token) {
        this.token = token;
        return this;
    }
}