import dataclasses
from typing import Optional
from dataclasses_json import dataclass_json
from epilot import utils
from ..shared import userv2 as shared_userv2


@dataclasses.dataclass
class ListUsersV2QueryParams:
    limit: Optional[float] = dataclasses.field(default=None, metadata={'query_param': { 'field_name': 'limit', 'style': 'form', 'explode': True }})
    offset: Optional[float] = dataclasses.field(default=None, metadata={'query_param': { 'field_name': 'offset', 'style': 'form', 'explode': True }})
    query: Optional[str] = dataclasses.field(default=None, metadata={'query_param': { 'field_name': 'query', 'style': 'form', 'explode': True }})
    

@dataclass_json
@dataclasses.dataclass
class ListUsersV2200ApplicationJSON:
    results: Optional[list[shared_userv2.UserV2]] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.field_name('results') }})
    

@dataclasses.dataclass
class ListUsersV2Request:
    query_params: ListUsersV2QueryParams = dataclasses.field()
    

@dataclasses.dataclass
class ListUsersV2Response:
    content_type: str = dataclasses.field()
    status_code: int = dataclasses.field()
    list_users_v2_200_application_json_object: Optional[ListUsersV2200ApplicationJSON] = dataclasses.field(default=None)
    
