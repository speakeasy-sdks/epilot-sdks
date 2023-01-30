import { SpeakeasyMetadata, SpeakeasyBase } from "../../../internal/utils";
import * as shared from "../shared";



export class GetUserLoginParametersPathParams extends SpeakeasyBase {
  @SpeakeasyMetadata({ data: "pathParam, style=simple;explode=false;name=username" })
  username: string;
}


export class GetUserLoginParameters200ApplicationJson extends SpeakeasyBase {
  @SpeakeasyMetadata({ data: "json, name=login_parameters", elemType: shared.LoginParameters })
  loginParameters?: shared.LoginParameters[];
}


export class GetUserLoginParametersRequest extends SpeakeasyBase {
  @SpeakeasyMetadata()
  pathParams: GetUserLoginParametersPathParams;
}


export class GetUserLoginParametersResponse extends SpeakeasyBase {
  @SpeakeasyMetadata()
  contentType: string;

  @SpeakeasyMetadata()
  statusCode: number;

  @SpeakeasyMetadata()
  getUserLoginParameters200ApplicationJSONObject?: GetUserLoginParameters200ApplicationJson;
}
