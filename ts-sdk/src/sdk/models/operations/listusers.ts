import { SpeakeasyMetadata, SpeakeasyBase } from "../../../internal/utils";
import * as shared from "../shared";



export class ListUsersQueryParams extends SpeakeasyBase {
  @SpeakeasyMetadata({ data: "queryParam, style=form;explode=true;name=limit" })
  limit?: number;

  @SpeakeasyMetadata({ data: "queryParam, style=form;explode=true;name=offset" })
  offset?: number;

  @SpeakeasyMetadata({ data: "queryParam, style=form;explode=false;name=org_ids" })
  orgIds?: string[];

  @SpeakeasyMetadata({ data: "queryParam, style=form;explode=true;name=query" })
  query?: string;
}


export class ListUsers200ApplicationJson extends SpeakeasyBase {
  @SpeakeasyMetadata({ data: "json, name=users", elemType: shared.User })
  users?: shared.User[];
}


export class ListUsersRequest extends SpeakeasyBase {
  @SpeakeasyMetadata()
  queryParams: ListUsersQueryParams;
}


export class ListUsersResponse extends SpeakeasyBase {
  @SpeakeasyMetadata()
  contentType: string;

  @SpeakeasyMetadata()
  statusCode: number;

  @SpeakeasyMetadata()
  listUsers200ApplicationJSONObject?: ListUsers200ApplicationJson;
}
