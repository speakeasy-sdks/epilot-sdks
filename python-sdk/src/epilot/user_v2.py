import requests
from typing import Optional
from epilot.models import shared, operations
from . import utils

class UserV2:
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

    
    def activate_user(self, request: operations.ActivateUserRequest) -> operations.ActivateUserResponse:
        r"""activateUser
        Activate user using an invite token
        """
        
        base_url = self._server_url
        
        url = base_url.removesuffix("/") + "/v2/users/public/activate"
        
        headers = {}
        req_content_type, data, json, files = utils.serialize_request_body(request)
        if req_content_type != "multipart/form-data" and req_content_type != "multipart/mixed":
            headers["content-type"] = req_content_type
        query_params = utils.get_query_params(request.query_params)
        
        client = self._security_client
        
        r = client.request("POST", url, params=query_params, data=data, json=json, files=files, headers=headers)
        content_type = r.headers.get("Content-Type")

        res = operations.ActivateUserResponse(status_code=r.status_code, content_type=content_type)
        
        if r.status_code == 200:
            pass
        elif r.status_code == 404:
            pass

        return res

    
    def delete_user_v2(self, request: operations.DeleteUserV2Request) -> operations.DeleteUserV2Response:
        r"""deleteUserV2
        Delete user by user id
        """
        
        base_url = self._server_url
        
        url = utils.generate_url(base_url, "/v2/users/{id}", request.path_params)
        
        
        client = self._security_client
        
        r = client.request("DELETE", url)
        content_type = r.headers.get("Content-Type")

        res = operations.DeleteUserV2Response(status_code=r.status_code, content_type=content_type)
        
        if r.status_code == 200:
            if utils.match_content_type(content_type, "application/json"):
                out = utils.unmarshal_json(r.text, Optional[shared.User])
                res.user = out

        return res

    
    def get_me_v2(self) -> operations.GetMeV2Response:
        r"""getMeV2
        Get currently logged in user
        """
        
        base_url = self._server_url
        
        url = base_url.removesuffix("/") + "/v2/users/me"
        
        
        client = self._security_client
        
        r = client.request("GET", url)
        content_type = r.headers.get("Content-Type")

        res = operations.GetMeV2Response(status_code=r.status_code, content_type=content_type)
        
        if r.status_code == 200:
            if utils.match_content_type(content_type, "application/json"):
                out = utils.unmarshal_json(r.text, Optional[shared.UserV2])
                res.user_v2 = out

        return res

    
    def get_user_login_parameters_v2(self, request: operations.GetUserLoginParametersV2Request) -> operations.GetUserLoginParametersV2Response:
        r"""getUserLoginParametersV2
        Get user organization login parameters by username
        """
        
        base_url = self._server_url
        
        url = utils.generate_url(base_url, "/v2/users/public/username/{username}:getLoginParameters", request.path_params)
        
        
        client = self._security_client
        
        r = client.request("GET", url)
        content_type = r.headers.get("Content-Type")

        res = operations.GetUserLoginParametersV2Response(status_code=r.status_code, content_type=content_type)
        
        if r.status_code == 200:
            if utils.match_content_type(content_type, "application/json"):
                out = utils.unmarshal_json(r.text, Optional[operations.GetUserLoginParametersV2200ApplicationJSON])
                res.get_user_login_parameters_v2_200_application_json_object = out

        return res

    
    def get_user_v2(self, request: operations.GetUserV2Request) -> operations.GetUserV2Response:
        r"""getUserV2
        Get user details by user id
        """
        
        base_url = self._server_url
        
        url = utils.generate_url(base_url, "/v2/users/{id}", request.path_params)
        
        
        client = self._security_client
        
        r = client.request("GET", url)
        content_type = r.headers.get("Content-Type")

        res = operations.GetUserV2Response(status_code=r.status_code, content_type=content_type)
        
        if r.status_code == 200:
            if utils.match_content_type(content_type, "application/json"):
                out = utils.unmarshal_json(r.text, Optional[shared.UserV2])
                res.user_v2 = out

        return res

    
    def invite_user(self, request: operations.InviteUserRequest) -> operations.InviteUserResponse:
        r"""inviteUser
        Creates a new user in the caller's organization and sends an invite email to activate
        """
        
        base_url = self._server_url
        
        url = base_url.removesuffix("/") + "/v2/users/invite"
        
        headers = {}
        req_content_type, data, json, files = utils.serialize_request_body(request)
        if req_content_type != "multipart/form-data" and req_content_type != "multipart/mixed":
            headers["content-type"] = req_content_type
        
        client = self._security_client
        
        r = client.request("POST", url, data=data, json=json, files=files, headers=headers)
        content_type = r.headers.get("Content-Type")

        res = operations.InviteUserResponse(status_code=r.status_code, content_type=content_type)
        
        if r.status_code == 201:
            if utils.match_content_type(content_type, "application/json"):
                out = utils.unmarshal_json(r.text, Optional[shared.UserV2])
                res.user_v2 = out
        elif r.status_code == 400:
            pass

        return res

    
    def list_users_v2(self, request: operations.ListUsersV2Request) -> operations.ListUsersV2Response:
        r"""listUsersV2
        Get the list of organization users
        """
        
        base_url = self._server_url
        
        url = base_url.removesuffix("/") + "/v2/users"
        
        query_params = utils.get_query_params(request.query_params)
        
        client = self._security_client
        
        r = client.request("GET", url, params=query_params)
        content_type = r.headers.get("Content-Type")

        res = operations.ListUsersV2Response(status_code=r.status_code, content_type=content_type)
        
        if r.status_code == 200:
            if utils.match_content_type(content_type, "application/json"):
                out = utils.unmarshal_json(r.text, Optional[operations.ListUsersV2200ApplicationJSON])
                res.list_users_v2_200_application_json_object = out

        return res

    
    def resend_user_invitation(self, request: operations.ResendUserInvitationRequest) -> operations.ResendUserInvitationResponse:
        r"""resendUserInvitation
        Resend user invitation email
        """
        
        base_url = self._server_url
        
        url = base_url.removesuffix("/") + "/v2/users/invite:resendEmail"
        
        headers = {}
        req_content_type, data, json, files = utils.serialize_request_body(request)
        if req_content_type != "multipart/form-data" and req_content_type != "multipart/mixed":
            headers["content-type"] = req_content_type
        
        client = self._security_client
        
        r = client.request("POST", url, data=data, json=json, files=files, headers=headers)
        content_type = r.headers.get("Content-Type")

        res = operations.ResendUserInvitationResponse(status_code=r.status_code, content_type=content_type)
        
        if r.status_code == 200:
            if utils.match_content_type(content_type, "application/json"):
                out = utils.unmarshal_json(r.text, Optional[shared.UserV2])
                res.user_v2 = out
        elif r.status_code == 400:
            pass

        return res

    
    def sign_up_user(self, request: operations.SignUpUserRequest) -> operations.SignUpUserResponse:
        r"""signUpUser
        """
        
        base_url = self._server_url
        
        url = base_url.removesuffix("/") + "/v2/users/public/signup"
        
        headers = {}
        req_content_type, data, json, files = utils.serialize_request_body(request)
        if req_content_type != "multipart/form-data" and req_content_type != "multipart/mixed":
            headers["content-type"] = req_content_type
        query_params = utils.get_query_params(request.query_params)
        
        client = self._security_client
        
        r = client.request("POST", url, params=query_params, data=data, json=json, files=files, headers=headers)
        content_type = r.headers.get("Content-Type")

        res = operations.SignUpUserResponse(status_code=r.status_code, content_type=content_type)
        
        if r.status_code == 200:
            if utils.match_content_type(content_type, "application/json"):
                out = utils.unmarshal_json(r.text, Optional[operations.SignUpUser200ApplicationJSON])
                res.sign_up_user_200_application_json_object = out

        return res

    
    def update_user_v2(self, request: operations.UpdateUserV2Request) -> operations.UpdateUserV2Response:
        r"""updateUserV2
        Update user details
        """
        
        base_url = self._server_url
        
        url = utils.generate_url(base_url, "/v2/users/{id}", request.path_params)
        
        headers = {}
        req_content_type, data, json, files = utils.serialize_request_body(request)
        if req_content_type != "multipart/form-data" and req_content_type != "multipart/mixed":
            headers["content-type"] = req_content_type
        
        client = self._security_client
        
        r = client.request("PATCH", url, data=data, json=json, files=files, headers=headers)
        content_type = r.headers.get("Content-Type")

        res = operations.UpdateUserV2Response(status_code=r.status_code, content_type=content_type)
        
        if r.status_code == 200:
            if utils.match_content_type(content_type, "application/json"):
                out = utils.unmarshal_json(r.text, Optional[shared.UserV2])
                res.user_v2 = out

        return res

    
    def verify_email_with_token(self, request: operations.VerifyEmailWithTokenRequest) -> operations.VerifyEmailWithTokenResponse:
        r"""verifyEmailWithToken
        Update new email using an verification token
        """
        
        base_url = self._server_url
        
        url = base_url.removesuffix("/") + "/v2/users/public/verifyEmail"
        
        headers = {}
        req_content_type, data, json, files = utils.serialize_request_body(request)
        if req_content_type != "multipart/form-data" and req_content_type != "multipart/mixed":
            headers["content-type"] = req_content_type
        query_params = utils.get_query_params(request.query_params)
        
        client = self._security_client
        
        r = client.request("POST", url, params=query_params, data=data, json=json, files=files, headers=headers)
        content_type = r.headers.get("Content-Type")

        res = operations.VerifyEmailWithTokenResponse(status_code=r.status_code, content_type=content_type)
        
        if r.status_code == 200:
            pass
        elif r.status_code == 404:
            pass

        return res

    