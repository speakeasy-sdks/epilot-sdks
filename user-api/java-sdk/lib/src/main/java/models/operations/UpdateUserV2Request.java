package .models.operations;

import .utils.SpeakeasyMetadata;
public class UpdateUserV2Request {
    public UpdateUserV2PathParams pathParams;
    public UpdateUserV2Request withPathParams(UpdateUserV2PathParams pathParams) {
        this.pathParams = pathParams;
        return this;
    }
    @SpeakeasyMetadata("request:mediaType=application/json")
    public .models.shared.UserV2 request;
    public UpdateUserV2Request withRequest(.models.shared.UserV2 request) {
        this.request = request;
        return this;
    }
}