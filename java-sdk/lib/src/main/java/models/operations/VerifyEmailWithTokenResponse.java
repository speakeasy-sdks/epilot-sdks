package .models.operations;


public class VerifyEmailWithTokenResponse {
    public String contentType;
    public VerifyEmailWithTokenResponse withContentType(String contentType) {
        this.contentType = contentType;
        return this;
    }
    public Long statusCode;
    public VerifyEmailWithTokenResponse withStatusCode(Long statusCode) {
        this.statusCode = statusCode;
        return this;
    }
}