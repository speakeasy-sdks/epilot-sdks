import { SpeakeasyMetadata, SpeakeasyBase } from "../../../internal/utils";



export class UserVerificationPayload extends SpeakeasyBase {
  @SpeakeasyMetadata({ data: "json, name=password" })
  password?: string;
}
