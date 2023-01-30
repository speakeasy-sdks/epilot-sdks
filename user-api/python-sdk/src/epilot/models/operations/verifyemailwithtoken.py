import dataclasses
from typing import Optional
from ..shared import userverificationpayload as shared_userverificationpayload


@dataclasses.dataclass
class VerifyEmailWithTokenQueryParams:
    token: str = dataclasses.field(metadata={'query_param': { 'field_name': 'token', 'style': 'form', 'explode': True }})
    

@dataclasses.dataclass
class VerifyEmailWithTokenRequest:
    query_params: VerifyEmailWithTokenQueryParams = dataclasses.field()
    request: Optional[shared_userverificationpayload.UserVerificationPayload] = dataclasses.field(default=None, metadata={'request': { 'media_type': 'application/json' }})
    

@dataclasses.dataclass
class VerifyEmailWithTokenResponse:
    content_type: str = dataclasses.field()
    status_code: int = dataclasses.field()
    
