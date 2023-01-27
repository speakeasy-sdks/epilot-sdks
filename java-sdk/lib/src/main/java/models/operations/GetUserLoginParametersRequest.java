package .models.operations;


public class GetUserLoginParametersRequest {
    public GetUserLoginParametersPathParams pathParams;
    public GetUserLoginParametersRequest withPathParams(GetUserLoginParametersPathParams pathParams) {
        this.pathParams = pathParams;
        return this;
    }
}