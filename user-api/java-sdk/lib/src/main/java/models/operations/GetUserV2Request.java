package .models.operations;


public class GetUserV2Request {
    public GetUserV2PathParams pathParams;
    public GetUserV2Request withPathParams(GetUserV2PathParams pathParams) {
        this.pathParams = pathParams;
        return this;
    }
}