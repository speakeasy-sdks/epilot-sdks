package .models.operations;


public class GetMeResponse {
    public String contentType;
    public GetMeResponse withContentType(String contentType) {
        this.contentType = contentType;
        return this;
    }
    public Long statusCode;
    public GetMeResponse withStatusCode(Long statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public .models.shared.User user;
    public GetMeResponse withUser(.models.shared.User user) {
        this.user = user;
        return this;
    }
}