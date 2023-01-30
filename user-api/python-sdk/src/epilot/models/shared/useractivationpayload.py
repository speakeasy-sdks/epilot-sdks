import dataclasses
from typing import Optional
from dataclasses_json import dataclass_json
from epilot import utils


@dataclass_json
@dataclasses.dataclass
class UserActivationPayload:
    display_name: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.field_name('display_name') }})
    password: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.field_name('password') }})
    
