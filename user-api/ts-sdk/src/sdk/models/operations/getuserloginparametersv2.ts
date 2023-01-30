import { SpeakeasyMetadata, SpeakeasyBase } from "../../../internal/utils";
import * as shared from "../shared";



export class GetUserLoginParametersV2PathParams extends SpeakeasyBase {
  @SpeakeasyMetadata({ data: "pathParam, style=simple;explode=false;name=username" })
  username: string;
}


export class GetUserLoginParametersV2200ApplicationJson extends SpeakeasyBase {
  @SpeakeasyMetadata({ data: "json, name=login_parameters", elemType: shared.LoginParameters })
  loginParameters?: shared.LoginParameters[];
}


export class GetUserLoginParametersV2Request extends SpeakeasyBase {
  @SpeakeasyMetadata()
  pathParams: GetUserLoginParametersV2PathParams;
}


export class GetUserLoginParametersV2Response extends SpeakeasyBase {
  @SpeakeasyMetadata()
  contentType: string;

  @SpeakeasyMetadata()
  statusCode: number;

  @SpeakeasyMetadata()
  getUserLoginParametersV2200ApplicationJSONObject?: GetUserLoginParametersV2200ApplicationJson;
}
