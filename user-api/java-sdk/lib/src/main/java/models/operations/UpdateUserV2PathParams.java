package .models.operations;

import .utils.SpeakeasyMetadata;
public class UpdateUserV2PathParams {
    @SpeakeasyMetadata("pathParam:style=simple,explode=false,name=id")
    public String id;
    public UpdateUserV2PathParams withId(String id) {
        this.id = id;
        return this;
    }
}