import dataclasses
from datetime import date, datetime
from marshmallow import fields
import dateutil.parser
from typing import Optional
from ..shared import userv2 as shared_userv2


@dataclasses.dataclass
class UpdateUserV2PathParams:
    id: str = dataclasses.field(metadata={'path_param': { 'field_name': 'id', 'style': 'simple', 'explode': False }})
    

@dataclasses.dataclass
class UpdateUserV2Request:
    path_params: UpdateUserV2PathParams = dataclasses.field()
    request: Optional[shared_userv2.UserV2] = dataclasses.field(default=None, metadata={'request': { 'media_type': 'application/json' }})
    

@dataclasses.dataclass
class UpdateUserV2Response:
    content_type: str = dataclasses.field()
    status_code: int = dataclasses.field()
    user_v2: Optional[shared_userv2.UserV2] = dataclasses.field(default=None)
    
