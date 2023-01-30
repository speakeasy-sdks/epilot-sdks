package .models.operations;

import .utils.SpeakeasyMetadata;
public class VerifyEmailWithTokenRequest {
    public VerifyEmailWithTokenQueryParams queryParams;
    public VerifyEmailWithTokenRequest withQueryParams(VerifyEmailWithTokenQueryParams queryParams) {
        this.queryParams = queryParams;
        return this;
    }
    @SpeakeasyMetadata("request:mediaType=application/json")
    public .models.shared.UserVerificationPayload request;
    public VerifyEmailWithTokenRequest withRequest(.models.shared.UserVerificationPayload request) {
        this.request = request;
        return this;
    }
}