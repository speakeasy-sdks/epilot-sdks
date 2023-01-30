package .models.operations;


public class ListUsersV2Request {
    public ListUsersV2QueryParams queryParams;
    public ListUsersV2Request withQueryParams(ListUsersV2QueryParams queryParams) {
        this.queryParams = queryParams;
        return this;
    }
}