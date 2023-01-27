import { SpeakeasyMetadata, SpeakeasyBase } from "../../../internal/utils";
import * as shared from "../shared";



export class UpdateUserV2PathParams extends SpeakeasyBase {
  @SpeakeasyMetadata({ data: "pathParam, style=simple;explode=false;name=id" })
  id: string;
}


export class UpdateUserV2Request extends SpeakeasyBase {
  @SpeakeasyMetadata()
  pathParams: UpdateUserV2PathParams;

  @SpeakeasyMetadata({ data: "request, media_type=application/json" })
  request?: shared.UserV2;
}


export class UpdateUserV2Response extends SpeakeasyBase {
  @SpeakeasyMetadata()
  contentType: string;

  @SpeakeasyMetadata()
  statusCode: number;

  @SpeakeasyMetadata()
  userV2?: shared.UserV2;
}
