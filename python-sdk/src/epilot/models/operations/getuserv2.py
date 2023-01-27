import dataclasses
from typing import Optional
from ..shared import userv2 as shared_userv2


@dataclasses.dataclass
class GetUserV2PathParams:
    id: str = dataclasses.field(metadata={'path_param': { 'field_name': 'id', 'style': 'simple', 'explode': False }})
    

@dataclasses.dataclass
class GetUserV2Request:
    path_params: GetUserV2PathParams = dataclasses.field()
    

@dataclasses.dataclass
class GetUserV2Response:
    content_type: str = dataclasses.field()
    status_code: int = dataclasses.field()
    user_v2: Optional[shared_userv2.UserV2] = dataclasses.field(default=None)
    
