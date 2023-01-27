package .models.operations;

import .utils.SpeakeasyMetadata;
public class SignUpUserRequest {
    public SignUpUserQueryParams queryParams;
    public SignUpUserRequest withQueryParams(SignUpUserQueryParams queryParams) {
        this.queryParams = queryParams;
        return this;
    }
    @SpeakeasyMetadata("request:mediaType=application/json")
    public .models.shared.SignupUserPayload request;
    public SignUpUserRequest withRequest(.models.shared.SignupUserPayload request) {
        this.request = request;
        return this;
    }
}