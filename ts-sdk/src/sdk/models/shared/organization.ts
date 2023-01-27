import { SpeakeasyMetadata, SpeakeasyBase } from "../../../internal/utils";



export class OrganizationAddress extends SpeakeasyBase {
  @SpeakeasyMetadata({ data: "json, name=city" })
  city?: string;

  @SpeakeasyMetadata({ data: "json, name=country" })
  country?: string;

  @SpeakeasyMetadata({ data: "json, name=postal_code" })
  postalCode?: string;

  @SpeakeasyMetadata({ data: "json, name=street" })
  street?: string;

  @SpeakeasyMetadata({ data: "json, name=street_number" })
  streetNumber?: string;
}

export enum OrganizationTypeEnum {
    Vendor = "Vendor",
    Partner = "Partner"
}


export class Organization extends SpeakeasyBase {
  @SpeakeasyMetadata({ data: "json, name=address" })
  address?: OrganizationAddress;

  @SpeakeasyMetadata({ data: "json, name=email" })
  email?: string;

  @SpeakeasyMetadata({ data: "json, name=id" })
  id?: string;

  @SpeakeasyMetadata({ data: "json, name=is_unlicensed_org" })
  isUnlicensedOrg?: boolean;

  @SpeakeasyMetadata({ data: "json, name=logo_thumbnail_url" })
  logoThumbnailUrl?: string;

  @SpeakeasyMetadata({ data: "json, name=logo_url" })
  logoUrl?: string;

  @SpeakeasyMetadata({ data: "json, name=name" })
  name?: string;

  @SpeakeasyMetadata({ data: "json, name=phone" })
  phone?: string;

  @SpeakeasyMetadata({ data: "json, name=pricing_tier" })
  pricingTier?: string;

  @SpeakeasyMetadata({ data: "json, name=signature" })
  signature?: string;

  @SpeakeasyMetadata({ data: "json, name=symbol" })
  symbol?: string;

  @SpeakeasyMetadata({ data: "json, name=type" })
  type?: OrganizationTypeEnum;

  @SpeakeasyMetadata({ data: "json, name=website" })
  website?: string;
}
