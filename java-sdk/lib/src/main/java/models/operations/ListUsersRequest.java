package .models.operations;


public class ListUsersRequest {
    public ListUsersQueryParams queryParams;
    public ListUsersRequest withQueryParams(ListUsersQueryParams queryParams) {
        this.queryParams = queryParams;
        return this;
    }
}