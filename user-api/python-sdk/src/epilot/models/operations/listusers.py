import dataclasses
from typing import Optional
from dataclasses_json import dataclass_json
from epilot import utils
from ..shared import user as shared_user


@dataclasses.dataclass
class ListUsersQueryParams:
    limit: Optional[float] = dataclasses.field(default=None, metadata={'query_param': { 'field_name': 'limit', 'style': 'form', 'explode': True }})
    offset: Optional[float] = dataclasses.field(default=None, metadata={'query_param': { 'field_name': 'offset', 'style': 'form', 'explode': True }})
    org_ids: Optional[list[str]] = dataclasses.field(default=None, metadata={'query_param': { 'field_name': 'org_ids', 'style': 'form', 'explode': False }})
    query: Optional[str] = dataclasses.field(default=None, metadata={'query_param': { 'field_name': 'query', 'style': 'form', 'explode': True }})
    

@dataclass_json
@dataclasses.dataclass
class ListUsers200ApplicationJSON:
    users: Optional[list[shared_user.User]] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.field_name('users') }})
    

@dataclasses.dataclass
class ListUsersRequest:
    query_params: ListUsersQueryParams = dataclasses.field()
    

@dataclasses.dataclass
class ListUsersResponse:
    content_type: str = dataclasses.field()
    status_code: int = dataclasses.field()
    list_users_200_application_json_object: Optional[ListUsers200ApplicationJSON] = dataclasses.field(default=None)
    
