package .models.operations;

import .utils.SpeakeasyMetadata;
public class GetUserLoginParametersPathParams {
    @SpeakeasyMetadata("pathParam:style=simple,explode=false,name=username")
    public String username;
    public GetUserLoginParametersPathParams withUsername(String username) {
        this.username = username;
        return this;
    }
}