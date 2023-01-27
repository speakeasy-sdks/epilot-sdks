package .models.operations;

import .utils.SpeakeasyMetadata;
public class DeleteUserV2PathParams {
    @SpeakeasyMetadata("pathParam:style=simple,explode=false,name=id")
    public String id;
    public DeleteUserV2PathParams withId(String id) {
        this.id = id;
        return this;
    }
}