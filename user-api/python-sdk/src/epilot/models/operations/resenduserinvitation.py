import dataclasses
from typing import Optional
from enum import Enum
from dataclasses_json import dataclass_json
from epilot import utils
from ..shared import userv2 as shared_userv2

class ResendUserInvitationRequestBodyLanguageEnum(str, Enum):
    EN = "en"
    DE = "de"


@dataclass_json
@dataclasses.dataclass
class ResendUserInvitationRequestBody:
    email: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.field_name('email') }})
    language: Optional[ResendUserInvitationRequestBodyLanguageEnum] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.field_name('language') }})
    

@dataclasses.dataclass
class ResendUserInvitationRequest:
    request: Optional[ResendUserInvitationRequestBody] = dataclasses.field(default=None, metadata={'request': { 'media_type': 'application/json' }})
    

@dataclasses.dataclass
class ResendUserInvitationResponse:
    content_type: str = dataclasses.field()
    status_code: int = dataclasses.field()
    user_v2: Optional[shared_userv2.UserV2] = dataclasses.field(default=None)
    
