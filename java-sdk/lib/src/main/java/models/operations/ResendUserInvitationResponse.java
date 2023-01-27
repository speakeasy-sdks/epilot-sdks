package .models.operations;


public class ResendUserInvitationResponse {
    public String contentType;
    public ResendUserInvitationResponse withContentType(String contentType) {
        this.contentType = contentType;
        return this;
    }
    public Long statusCode;
    public ResendUserInvitationResponse withStatusCode(Long statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public .models.shared.UserV2 userV2;
    public ResendUserInvitationResponse withUserV2(.models.shared.UserV2 userV2) {
        this.userV2 = userV2;
        return this;
    }
}