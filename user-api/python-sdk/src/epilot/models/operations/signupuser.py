import dataclasses
from typing import Optional
from dataclasses_json import dataclass_json
from epilot import utils
from ..shared import organization as shared_organization
from ..shared import user as shared_user
from ..shared import signupuserpayload as shared_signupuserpayload


@dataclasses.dataclass
class SignUpUserQueryParams:
    token: Optional[str] = dataclasses.field(default=None, metadata={'query_param': { 'field_name': 'token', 'style': 'form', 'explode': True }})
    

@dataclass_json
@dataclasses.dataclass
class SignUpUser200ApplicationJSON:
    organization: Optional[shared_organization.Organization] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.field_name('organization') }})
    user: Optional[shared_user.User] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.field_name('user') }})
    

@dataclasses.dataclass
class SignUpUserRequest:
    query_params: SignUpUserQueryParams = dataclasses.field()
    request: Optional[shared_signupuserpayload.SignupUserPayload] = dataclasses.field(default=None, metadata={'request': { 'media_type': 'application/json' }})
    

@dataclasses.dataclass
class SignUpUserResponse:
    content_type: str = dataclasses.field()
    status_code: int = dataclasses.field()
    sign_up_user_200_application_json_object: Optional[SignUpUser200ApplicationJSON] = dataclasses.field(default=None)
    
