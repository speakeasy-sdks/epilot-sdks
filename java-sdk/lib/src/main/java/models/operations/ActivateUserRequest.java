package .models.operations;

import .utils.SpeakeasyMetadata;
public class ActivateUserRequest {
    public ActivateUserQueryParams queryParams;
    public ActivateUserRequest withQueryParams(ActivateUserQueryParams queryParams) {
        this.queryParams = queryParams;
        return this;
    }
    @SpeakeasyMetadata("request:mediaType=application/json")
    public .models.shared.UserActivationPayload request;
    public ActivateUserRequest withRequest(.models.shared.UserActivationPayload request) {
        this.request = request;
        return this;
    }
}