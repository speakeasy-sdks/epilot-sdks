import dataclasses
from typing import Optional
from dataclasses_json import dataclass_json
from epilot import utils
from ..shared import loginparameters as shared_loginparameters


@dataclasses.dataclass
class GetUserLoginParametersPathParams:
    username: str = dataclasses.field(metadata={'path_param': { 'field_name': 'username', 'style': 'simple', 'explode': False }})
    

@dataclass_json
@dataclasses.dataclass
class GetUserLoginParameters200ApplicationJSON:
    login_parameters: Optional[list[shared_loginparameters.LoginParameters]] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.field_name('login_parameters') }})
    

@dataclasses.dataclass
class GetUserLoginParametersRequest:
    path_params: GetUserLoginParametersPathParams = dataclasses.field()
    

@dataclasses.dataclass
class GetUserLoginParametersResponse:
    content_type: str = dataclasses.field()
    status_code: int = dataclasses.field()
    get_user_login_parameters_200_application_json_object: Optional[GetUserLoginParameters200ApplicationJSON] = dataclasses.field(default=None)
    
