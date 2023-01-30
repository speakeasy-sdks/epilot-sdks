package .models.operations;

import .utils.SpeakeasyMetadata;
public class GetUserV2PathParams {
    @SpeakeasyMetadata("pathParam:style=simple,explode=false,name=id")
    public String id;
    public GetUserV2PathParams withId(String id) {
        this.id = id;
        return this;
    }
}