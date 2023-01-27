import dataclasses
from typing import Optional
from dataclasses_json import dataclass_json
from epilot import utils
from ..shared import loginparameters as shared_loginparameters


@dataclasses.dataclass
class GetUserLoginParametersV2PathParams:
    username: str = dataclasses.field(metadata={'path_param': { 'field_name': 'username', 'style': 'simple', 'explode': False }})
    

@dataclass_json
@dataclasses.dataclass
class GetUserLoginParametersV2200ApplicationJSON:
    login_parameters: Optional[list[shared_loginparameters.LoginParameters]] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.field_name('login_parameters') }})
    

@dataclasses.dataclass
class GetUserLoginParametersV2Request:
    path_params: GetUserLoginParametersV2PathParams = dataclasses.field()
    

@dataclasses.dataclass
class GetUserLoginParametersV2Response:
    content_type: str = dataclasses.field()
    status_code: int = dataclasses.field()
    get_user_login_parameters_v2_200_application_json_object: Optional[GetUserLoginParametersV2200ApplicationJSON] = dataclasses.field(default=None)
    
