import { SpeakeasyMetadata, SpeakeasyBase } from "../../../internal/utils";



export class UserProperties extends SpeakeasyBase {
  @SpeakeasyMetadata({ data: "json, name=name" })
  name: string;

  @SpeakeasyMetadata({ data: "json, name=value" })
  value: string;
}


export class User extends SpeakeasyBase {
  @SpeakeasyMetadata({ data: "json, name=display_name" })
  displayName?: string;

  @SpeakeasyMetadata({ data: "json, name=email" })
  email: string;

  @SpeakeasyMetadata({ data: "json, name=id" })
  id: string;

  @SpeakeasyMetadata({ data: "json, name=image_uri" })
  imageUri?: Record<string, any>;

  @SpeakeasyMetadata({ data: "json, name=name" })
  name: string;

  @SpeakeasyMetadata({ data: "json, name=organization_id" })
  organizationId: string;

  @SpeakeasyMetadata({ data: "json, name=preferred_language" })
  preferredLanguage: string;

  @SpeakeasyMetadata({ data: "json, name=properties", elemType: UserProperties })
  properties: UserProperties[];

  @SpeakeasyMetadata({ data: "json, name=roles" })
  roles: string[];

  @SpeakeasyMetadata({ data: "json, name=signature" })
  signature?: string;
}
