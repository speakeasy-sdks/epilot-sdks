package .models.operations;

import .utils.SpeakeasyMetadata;
public class InviteUserRequest {
    @SpeakeasyMetadata("request:mediaType=application/json")
    public .models.shared.UserInvitationPayload request;
    public InviteUserRequest withRequest(.models.shared.UserInvitationPayload request) {
        this.request = request;
        return this;
    }
}