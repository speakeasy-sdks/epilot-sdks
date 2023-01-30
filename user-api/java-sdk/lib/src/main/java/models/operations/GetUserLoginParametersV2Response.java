package .models.operations;


public class GetUserLoginParametersV2Response {
    public String contentType;
    public GetUserLoginParametersV2Response withContentType(String contentType) {
        this.contentType = contentType;
        return this;
    }
    public Long statusCode;
    public GetUserLoginParametersV2Response withStatusCode(Long statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public GetUserLoginParametersV2200ApplicationJson getUserLoginParametersV2200ApplicationJSONObject;
    public GetUserLoginParametersV2Response withGetUserLoginParametersV2200ApplicationJsonObject(GetUserLoginParametersV2200ApplicationJson getUserLoginParametersV2200ApplicationJSONObject) {
        this.getUserLoginParametersV2200ApplicationJSONObject = getUserLoginParametersV2200ApplicationJSONObject;
        return this;
    }
}