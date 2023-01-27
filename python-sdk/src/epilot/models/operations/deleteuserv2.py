import dataclasses
from typing import Optional
from ..shared import user as shared_user


@dataclasses.dataclass
class DeleteUserV2PathParams:
    id: str = dataclasses.field(metadata={'path_param': { 'field_name': 'id', 'style': 'simple', 'explode': False }})
    

@dataclasses.dataclass
class DeleteUserV2Request:
    path_params: DeleteUserV2PathParams = dataclasses.field()
    

@dataclasses.dataclass
class DeleteUserV2Response:
    content_type: str = dataclasses.field()
    status_code: int = dataclasses.field()
    user: Optional[shared_user.User] = dataclasses.field(default=None)
    
