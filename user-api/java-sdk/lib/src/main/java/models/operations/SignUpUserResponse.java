package .models.operations;


public class SignUpUserResponse {
    public String contentType;
    public SignUpUserResponse withContentType(String contentType) {
        this.contentType = contentType;
        return this;
    }
    public Long statusCode;
    public SignUpUserResponse withStatusCode(Long statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public SignUpUser200ApplicationJson signUpUser200ApplicationJSONObject;
    public SignUpUserResponse withSignUpUser200ApplicationJsonObject(SignUpUser200ApplicationJson signUpUser200ApplicationJSONObject) {
        this.signUpUser200ApplicationJSONObject = signUpUser200ApplicationJSONObject;
        return this;
    }
}