package .models.operations;

import .utils.SpeakeasyMetadata;
public class ListUsersQueryParams {
    @SpeakeasyMetadata("queryParam:style=form,explode=true,name=limit")
    public Double limit;
    public ListUsersQueryParams withLimit(Double limit) {
        this.limit = limit;
        return this;
    }
    @SpeakeasyMetadata("queryParam:style=form,explode=true,name=offset")
    public Double offset;
    public ListUsersQueryParams withOffset(Double offset) {
        this.offset = offset;
        return this;
    }
    @SpeakeasyMetadata("queryParam:style=form,explode=false,name=org_ids")
    public String[] orgIds;
    public ListUsersQueryParams withOrgIds(String[] orgIds) {
        this.orgIds = orgIds;
        return this;
    }
    @SpeakeasyMetadata("queryParam:style=form,explode=true,name=query")
    public String query;
    public ListUsersQueryParams withQuery(String query) {
        this.query = query;
        return this;
    }
}