import { SpeakeasyMetadata, SpeakeasyBase } from "../../../internal/utils";


export enum UserInvitationPayloadLanguageEnum {
    En = "en",
    De = "de"
}


export class UserInvitationPayload extends SpeakeasyBase {
  @SpeakeasyMetadata({ data: "json, name=email" })
  email?: string;

  @SpeakeasyMetadata({ data: "json, name=language" })
  language?: UserInvitationPayloadLanguageEnum;

  @SpeakeasyMetadata({ data: "json, name=roles" })
  roles?: string[];
}
