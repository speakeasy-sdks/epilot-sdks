package .models.operations;


public class DeleteUserV2Response {
    public String contentType;
    public DeleteUserV2Response withContentType(String contentType) {
        this.contentType = contentType;
        return this;
    }
    public Long statusCode;
    public DeleteUserV2Response withStatusCode(Long statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public .models.shared.User user;
    public DeleteUserV2Response withUser(.models.shared.User user) {
        this.user = user;
        return this;
    }
}