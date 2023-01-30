package .models.operations;

import .utils.SpeakeasyMetadata;
public class GetUserLoginParametersV2PathParams {
    @SpeakeasyMetadata("pathParam:style=simple,explode=false,name=username")
    public String username;
    public GetUserLoginParametersV2PathParams withUsername(String username) {
        this.username = username;
        return this;
    }
}