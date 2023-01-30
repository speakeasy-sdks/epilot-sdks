package .models.operations;


public class UpdateUserV2Response {
    public String contentType;
    public UpdateUserV2Response withContentType(String contentType) {
        this.contentType = contentType;
        return this;
    }
    public Long statusCode;
    public UpdateUserV2Response withStatusCode(Long statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public .models.shared.UserV2 userV2;
    public UpdateUserV2Response withUserV2(.models.shared.UserV2 userV2) {
        this.userV2 = userV2;
        return this;
    }
}