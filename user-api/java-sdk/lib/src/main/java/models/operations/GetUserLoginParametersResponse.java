package .models.operations;


public class GetUserLoginParametersResponse {
    public String contentType;
    public GetUserLoginParametersResponse withContentType(String contentType) {
        this.contentType = contentType;
        return this;
    }
    public Long statusCode;
    public GetUserLoginParametersResponse withStatusCode(Long statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public GetUserLoginParameters200ApplicationJson getUserLoginParameters200ApplicationJSONObject;
    public GetUserLoginParametersResponse withGetUserLoginParameters200ApplicationJsonObject(GetUserLoginParameters200ApplicationJson getUserLoginParameters200ApplicationJSONObject) {
        this.getUserLoginParameters200ApplicationJSONObject = getUserLoginParameters200ApplicationJSONObject;
        return this;
    }
}