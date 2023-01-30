import dataclasses
from typing import Optional
from enum import Enum
from dataclasses_json import dataclass_json
from epilot import utils

class UserInvitationPayloadLanguageEnum(str, Enum):
    EN = "en"
    DE = "de"


@dataclass_json
@dataclasses.dataclass
class UserInvitationPayload:
    email: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.field_name('email') }})
    language: Optional[UserInvitationPayloadLanguageEnum] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.field_name('language') }})
    roles: Optional[list[str]] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.field_name('roles') }})
    
