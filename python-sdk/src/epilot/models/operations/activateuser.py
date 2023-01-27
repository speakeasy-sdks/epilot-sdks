import dataclasses
from typing import Optional
from ..shared import useractivationpayload as shared_useractivationpayload


@dataclasses.dataclass
class ActivateUserQueryParams:
    token: str = dataclasses.field(metadata={'query_param': { 'field_name': 'token', 'style': 'form', 'explode': True }})
    

@dataclasses.dataclass
class ActivateUserRequest:
    query_params: ActivateUserQueryParams = dataclasses.field()
    request: Optional[shared_useractivationpayload.UserActivationPayload] = dataclasses.field(default=None, metadata={'request': { 'media_type': 'application/json' }})
    

@dataclasses.dataclass
class ActivateUserResponse:
    content_type: str = dataclasses.field()
    status_code: int = dataclasses.field()
    
