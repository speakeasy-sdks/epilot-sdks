import dataclasses
from typing import Optional
from enum import Enum
from dataclasses_json import dataclass_json
from epilot import utils

class LoginParametersOauthResponseTypeEnum(str, Enum):
    CODE = "code"
    TOKEN = "token"


@dataclass_json
@dataclasses.dataclass
class LoginParameters:
    cognito_identity_pool_id: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.field_name('cognito_identity_pool_id') }})
    cognito_oauth_domain: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.field_name('cognito_oauth_domain') }})
    cognito_oauth_scopes: Optional[list[str]] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.field_name('cognito_oauth_scopes') }})
    cognito_region: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.field_name('cognito_region') }})
    cognito_user_pool_client_id: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.field_name('cognito_user_pool_client_id') }})
    cognito_user_pool_id: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.field_name('cognito_user_pool_id') }})
    oauth_response_type: Optional[LoginParametersOauthResponseTypeEnum] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.field_name('oauth_response_type') }})
    organization_id: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.field_name('organization_id') }})
    organization_name: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.field_name('organization_name') }})
    organization_type: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.field_name('organization_type') }})
    
