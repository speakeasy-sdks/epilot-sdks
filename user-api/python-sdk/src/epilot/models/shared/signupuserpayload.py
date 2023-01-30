import dataclasses
from typing import Any,Optional
from enum import Enum
from dataclasses_json import dataclass_json
from epilot import utils
from ..shared import userdetail as shared_userdetail

class SignupUserPayloadLanguageEnum(str, Enum):
    EN = "en"
    DE = "de"


@dataclass_json
@dataclasses.dataclass
class SignupUserPayload:
    language: Optional[SignupUserPayloadLanguageEnum] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.field_name('language') }})
    organization_detail: Optional[dict[str, Any]] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.field_name('organization_detail') }})
    user_detail: Optional[shared_userdetail.UserDetail] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.field_name('user_detail') }})
    
