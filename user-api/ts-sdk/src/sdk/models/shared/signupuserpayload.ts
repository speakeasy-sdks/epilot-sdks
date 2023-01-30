import { SpeakeasyMetadata, SpeakeasyBase } from "../../../internal/utils";
import { UserDetail } from "./userdetail";


export enum SignupUserPayloadLanguageEnum {
    En = "en",
    De = "de"
}


export class SignupUserPayload extends SpeakeasyBase {
  @SpeakeasyMetadata({ data: "json, name=language" })
  language?: SignupUserPayloadLanguageEnum;

  @SpeakeasyMetadata({ data: "json, name=organization_detail" })
  organizationDetail?: Record<string, any>;

  @SpeakeasyMetadata({ data: "json, name=user_detail" })
  userDetail?: UserDetail;
}
