package .models.operations;


public class InviteUserResponse {
    public String contentType;
    public InviteUserResponse withContentType(String contentType) {
        this.contentType = contentType;
        return this;
    }
    public Long statusCode;
    public InviteUserResponse withStatusCode(Long statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public .models.shared.UserV2 userV2;
    public InviteUserResponse withUserV2(.models.shared.UserV2 userV2) {
        this.userV2 = userV2;
        return this;
    }
}