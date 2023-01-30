package .models.operations;


public class ListUsersV2Response {
    public String contentType;
    public ListUsersV2Response withContentType(String contentType) {
        this.contentType = contentType;
        return this;
    }
    public Long statusCode;
    public ListUsersV2Response withStatusCode(Long statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    public ListUsersV2200ApplicationJson listUsersV2200ApplicationJSONObject;
    public ListUsersV2Response withListUsersV2200ApplicationJsonObject(ListUsersV2200ApplicationJson listUsersV2200ApplicationJSONObject) {
        this.listUsersV2200ApplicationJSONObject = listUsersV2200ApplicationJSONObject;
        return this;
    }
}