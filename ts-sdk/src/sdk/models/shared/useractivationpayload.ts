import { SpeakeasyMetadata, SpeakeasyBase } from "../../../internal/utils";



export class UserActivationPayload extends SpeakeasyBase {
  @SpeakeasyMetadata({ data: "json, name=display_name" })
  displayName?: string;

  @SpeakeasyMetadata({ data: "json, name=password" })
  password?: string;
}
