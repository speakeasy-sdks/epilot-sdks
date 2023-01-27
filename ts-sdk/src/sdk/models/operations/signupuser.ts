import { SpeakeasyMetadata, SpeakeasyBase } from "../../../internal/utils";
import * as shared from "../shared";



export class SignUpUserQueryParams extends SpeakeasyBase {
  @SpeakeasyMetadata({ data: "queryParam, style=form;explode=true;name=token" })
  token?: string;
}


export class SignUpUser200ApplicationJson extends SpeakeasyBase {
  @SpeakeasyMetadata({ data: "json, name=organization" })
  organization?: shared.Organization;

  @SpeakeasyMetadata({ data: "json, name=user" })
  user?: shared.User;
}


export class SignUpUserRequest extends SpeakeasyBase {
  @SpeakeasyMetadata()
  queryParams: SignUpUserQueryParams;

  @SpeakeasyMetadata({ data: "request, media_type=application/json" })
  request?: shared.SignupUserPayload;
}


export class SignUpUserResponse extends SpeakeasyBase {
  @SpeakeasyMetadata()
  contentType: string;

  @SpeakeasyMetadata()
  statusCode: number;

  @SpeakeasyMetadata()
  signUpUser200ApplicationJSONObject?: SignUpUser200ApplicationJson;
}
