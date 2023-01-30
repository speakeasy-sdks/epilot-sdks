import { SpeakeasyMetadata, SpeakeasyBase } from "../../../internal/utils";
import * as shared from "../shared";


export enum ResendUserInvitationRequestBodyLanguageEnum {
    En = "en",
    De = "de"
}


export class ResendUserInvitationRequestBody extends SpeakeasyBase {
  @SpeakeasyMetadata({ data: "json, name=email" })
  email?: string;

  @SpeakeasyMetadata({ data: "json, name=language" })
  language?: ResendUserInvitationRequestBodyLanguageEnum;
}


export class ResendUserInvitationRequest extends SpeakeasyBase {
  @SpeakeasyMetadata({ data: "request, media_type=application/json" })
  request?: ResendUserInvitationRequestBody;
}


export class ResendUserInvitationResponse extends SpeakeasyBase {
  @SpeakeasyMetadata()
  contentType: string;

  @SpeakeasyMetadata()
  statusCode: number;

  @SpeakeasyMetadata()
  userV2?: shared.UserV2;
}
