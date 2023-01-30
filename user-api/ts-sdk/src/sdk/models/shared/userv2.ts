import { SpeakeasyMetadata, SpeakeasyBase } from "../../../internal/utils";



export class UserV2Properties extends SpeakeasyBase {
  @SpeakeasyMetadata({ data: "json, name=name" })
  name: string;

  @SpeakeasyMetadata({ data: "json, name=value" })
  value: string;
}

export enum UserV2StatusEnum {
    Active = "Active",
    Pending = "Pending",
    Deactivated = "Deactivated",
    Deleted = "Deleted"
}


export class UserV2 extends SpeakeasyBase {
  @SpeakeasyMetadata({ data: "json, name=created_at" })
  createdAt?: string;

  @SpeakeasyMetadata({ data: "json, name=department" })
  department?: string;

  @SpeakeasyMetadata({ data: "json, name=display_name" })
  displayName?: string;

  @SpeakeasyMetadata({ data: "json, name=draft_email" })
  draftEmail?: string;

  @SpeakeasyMetadata({ data: "json, name=email" })
  email?: string;

  @SpeakeasyMetadata({ data: "json, name=id" })
  id?: string;

  @SpeakeasyMetadata({ data: "json, name=image_uri" })
  imageUri?: Record<string, any>;

  @SpeakeasyMetadata({ data: "json, name=is_signature_enabled" })
  isSignatureEnabled?: boolean;

  @SpeakeasyMetadata({ data: "json, name=mfa_enabled" })
  mfaEnabled?: boolean;

  @SpeakeasyMetadata({ data: "json, name=organization_id" })
  organizationId?: string;

  @SpeakeasyMetadata({ data: "json, name=phone" })
  phone?: string;

  @SpeakeasyMetadata({ data: "json, name=phone_verified" })
  phoneVerified?: boolean;

  @SpeakeasyMetadata({ data: "json, name=preferred_language" })
  preferredLanguage?: string;

  @SpeakeasyMetadata({ data: "json, name=properties", elemType: UserV2Properties })
  properties?: UserV2Properties[];

  @SpeakeasyMetadata({ data: "json, name=signature" })
  signature?: string;

  @SpeakeasyMetadata({ data: "json, name=status" })
  status?: UserV2StatusEnum;

  @SpeakeasyMetadata({ data: "json, name=token" })
  token?: string;
}
