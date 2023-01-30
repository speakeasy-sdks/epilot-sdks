package .models.operations;


public class DeleteUserV2Request {
    public DeleteUserV2PathParams pathParams;
    public DeleteUserV2Request withPathParams(DeleteUserV2PathParams pathParams) {
        this.pathParams = pathParams;
        return this;
    }
}