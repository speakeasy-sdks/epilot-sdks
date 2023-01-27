package .models.operations;


public class GetUserV2Response {
    public String contentType;
    public GetUserV2Response withContentType(String contentType) {
        this.contentType = contentType;
        return this;
    }
    public Long statusCode;
    public GetUserV2Response withStatusCode(Long statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public .models.shared.UserV2 userV2;
    public GetUserV2Response withUserV2(.models.shared.UserV2 userV2) {
        this.userV2 = userV2;
        return this;
    }
}