import { SpeakeasyMetadata, SpeakeasyBase } from "../../../internal/utils";



export class UserDetail extends SpeakeasyBase {
  @SpeakeasyMetadata({ data: "json, name=email" })
  email: string;

  @SpeakeasyMetadata({ data: "json, name=full_name" })
  fullName: string;

  @SpeakeasyMetadata({ data: "json, name=password" })
  password: string;
}
