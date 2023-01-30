import dataclasses
from typing import Optional
from ..shared import userv2 as shared_userv2


@dataclasses.dataclass
class GetMeV2Response:
    content_type: str = dataclasses.field()
    status_code: int = dataclasses.field()
    user_v2: Optional[shared_userv2.UserV2] = dataclasses.field(default=None)
    
