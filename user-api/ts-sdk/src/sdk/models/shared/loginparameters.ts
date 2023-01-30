import { SpeakeasyMetadata, SpeakeasyBase } from "../../../internal/utils";


export enum LoginParametersOauthResponseTypeEnum {
    Code = "code",
    Token = "token"
}


export class LoginParameters extends SpeakeasyBase {
  @SpeakeasyMetadata({ data: "json, name=cognito_identity_pool_id" })
  cognitoIdentityPoolId?: string;

  @SpeakeasyMetadata({ data: "json, name=cognito_oauth_domain" })
  cognitoOauthDomain?: string;

  @SpeakeasyMetadata({ data: "json, name=cognito_oauth_scopes" })
  cognitoOauthScopes?: string[];

  @SpeakeasyMetadata({ data: "json, name=cognito_region" })
  cognitoRegion?: string;

  @SpeakeasyMetadata({ data: "json, name=cognito_user_pool_client_id" })
  cognitoUserPoolClientId?: string;

  @SpeakeasyMetadata({ data: "json, name=cognito_user_pool_id" })
  cognitoUserPoolId?: string;

  @SpeakeasyMetadata({ data: "json, name=oauth_response_type" })
  oauthResponseType?: LoginParametersOauthResponseTypeEnum;

  @SpeakeasyMetadata({ data: "json, name=organization_id" })
  organizationId?: string;

  @SpeakeasyMetadata({ data: "json, name=organization_name" })
  organizationName?: string;

  @SpeakeasyMetadata({ data: "json, name=organization_type" })
  organizationType?: string;
}
