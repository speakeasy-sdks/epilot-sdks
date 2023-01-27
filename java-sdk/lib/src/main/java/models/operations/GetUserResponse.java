package .models.operations;


public class GetUserResponse {
    public String contentType;
    public GetUserResponse withContentType(String contentType) {
        this.contentType = contentType;
        return this;
    }
    public Long statusCode;
    public GetUserResponse withStatusCode(Long statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public .models.shared.User user;
    public GetUserResponse withUser(.models.shared.User user) {
        this.user = user;
        return this;
    }
}