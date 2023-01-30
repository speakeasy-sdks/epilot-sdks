import requests
from typing import Optional
from epilot.models import shared, operations
from . import utils

class UserV1:
    _client: requests.Session
    _security_client: requests.Session
    _server_url: str
    _language: str
    _sdk_version: str
    _gen_version: str

    def __init__(self, client: requests.Session, security_client: requests.Session, server_url: str, language: str, sdk_version: str, gen_version: str) -> None:
        self._client = client
        self._security_client = security_client
        self._server_url = server_url
        self._language = language
        self._sdk_version = sdk_version
        self._gen_version = gen_version

    
    def get_me(self) -> operations.GetMeResponse:
        r"""getMe
        Get currently logged in user
        """
        
        base_url = self._server_url
        
        url = base_url.removesuffix("/") + "/v1/users/me"
        
        
        client = self._security_client
        
        r = client.request("GET", url)
        content_type = r.headers.get("Content-Type")

        res = operations.GetMeResponse(status_code=r.status_code, content_type=content_type)
        
        if r.status_code == 200:
            if utils.match_content_type(content_type, "application/json"):
                out = utils.unmarshal_json(r.text, Optional[shared.User])
                res.user = out

        return res

    
    def get_user(self, request: operations.GetUserRequest) -> operations.GetUserResponse:
        r"""getUser
        Get user by id
        """
        
        base_url = self._server_url
        
        url = utils.generate_url(base_url, "/v1/users/{id}", request.path_params)
        
        
        client = self._security_client
        
        r = client.request("GET", url)
        content_type = r.headers.get("Content-Type")

        res = operations.GetUserResponse(status_code=r.status_code, content_type=content_type)
        
        if r.status_code == 200:
            if utils.match_content_type(content_type, "application/json"):
                out = utils.unmarshal_json(r.text, Optional[shared.User])
                res.user = out

        return res

    
    def get_user_login_parameters(self, request: operations.GetUserLoginParametersRequest) -> operations.GetUserLoginParametersResponse:
        r"""getUserLoginParameters
        Get user organization login parameters by username
        """
        
        base_url = self._server_url
        
        url = utils.generate_url(base_url, "/v1/users/username/{username}:getLoginParameters", request.path_params)
        
        
        client = self._security_client
        
        r = client.request("GET", url)
        content_type = r.headers.get("Content-Type")

        res = operations.GetUserLoginParametersResponse(status_code=r.status_code, content_type=content_type)
        
        if r.status_code == 200:
            if utils.match_content_type(content_type, "application/json"):
                out = utils.unmarshal_json(r.text, Optional[operations.GetUserLoginParameters200ApplicationJSON])
                res.get_user_login_parameters_200_application_json_object = out

        return res

    
    def list_users(self, request: operations.ListUsersRequest) -> operations.ListUsersResponse:
        r"""listUsers
        Lists users in organizations you have access to
        """
        
        base_url = self._server_url
        
        url = base_url.removesuffix("/") + "/v1/users"
        
        query_params = utils.get_query_params(request.query_params)
        
        client = self._security_client
        
        r = client.request("GET", url, params=query_params)
        content_type = r.headers.get("Content-Type")

        res = operations.ListUsersResponse(status_code=r.status_code, content_type=content_type)
        
        if r.status_code == 200:
            if utils.match_content_type(content_type, "application/json"):
                out = utils.unmarshal_json(r.text, Optional[operations.ListUsers200ApplicationJSON])
                res.list_users_200_application_json_object = out

        return res

    