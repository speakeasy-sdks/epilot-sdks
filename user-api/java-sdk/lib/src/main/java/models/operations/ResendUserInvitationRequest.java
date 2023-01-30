package .models.operations;

import .utils.SpeakeasyMetadata;
public class ResendUserInvitationRequest {
    @SpeakeasyMetadata("request:mediaType=application/json")
    public ResendUserInvitationRequestBody request;
    public ResendUserInvitationRequest withRequest(ResendUserInvitationRequestBody request) {
        this.request = request;
        return this;
    }
}