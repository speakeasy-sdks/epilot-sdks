import { SpeakeasyMetadata, SpeakeasyBase } from "../../../internal/utils";
import * as shared from "../shared";



export class ListUsersV2QueryParams extends SpeakeasyBase {
  @SpeakeasyMetadata({ data: "queryParam, style=form;explode=true;name=limit" })
  limit?: number;

  @SpeakeasyMetadata({ data: "queryParam, style=form;explode=true;name=offset" })
  offset?: number;

  @SpeakeasyMetadata({ data: "queryParam, style=form;explode=true;name=query" })
  query?: string;
}


export class ListUsersV2200ApplicationJson extends SpeakeasyBase {
  @SpeakeasyMetadata({ data: "json, name=results", elemType: shared.UserV2 })
  results?: shared.UserV2[];
}


export class ListUsersV2Request extends SpeakeasyBase {
  @SpeakeasyMetadata()
  queryParams: ListUsersV2QueryParams;
}


export class ListUsersV2Response extends SpeakeasyBase {
  @SpeakeasyMetadata()
  contentType: string;

  @SpeakeasyMetadata()
  statusCode: number;

  @SpeakeasyMetadata()
  listUsersV2200ApplicationJSONObject?: ListUsersV2200ApplicationJson;
}
