package .models.operations;

import .utils.SpeakeasyMetadata;
public class ListUsersV2QueryParams {
    @SpeakeasyMetadata("queryParam:style=form,explode=true,name=limit")
    public Double limit;
    public ListUsersV2QueryParams withLimit(Double limit) {
        this.limit = limit;
        return this;
    }
    @SpeakeasyMetadata("queryParam:style=form,explode=true,name=offset")
    public Double offset;
    public ListUsersV2QueryParams withOffset(Double offset) {
        this.offset = offset;
        return this;
    }
    @SpeakeasyMetadata("queryParam:style=form,explode=true,name=query")
    public String query;
    public ListUsersV2QueryParams withQuery(String query) {
        this.query = query;
        return this;
    }
}