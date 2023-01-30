package .models.shared;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonInclude.Include;
public class LoginParameters {
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("cognito_identity_pool_id")
    public String cognitoIdentityPoolId;
    public LoginParameters withCognitoIdentityPoolId(String cognitoIdentityPoolId) {
        this.cognitoIdentityPoolId = cognitoIdentityPoolId;
        return this;
    }
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("cognito_oauth_domain")
    public String cognitoOauthDomain;
    public LoginParameters withCognitoOauthDomain(String cognitoOauthDomain) {
        this.cognitoOauthDomain = cognitoOauthDomain;
        return this;
    }
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("cognito_oauth_scopes")
    public String[] cognitoOauthScopes;
    public LoginParameters withCognitoOauthScopes(String[] cognitoOauthScopes) {
        this.cognitoOauthScopes = cognitoOauthScopes;
        return this;
    }
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("cognito_region")
    public String cognitoRegion;
    public LoginParameters withCognitoRegion(String cognitoRegion) {
        this.cognitoRegion = cognitoRegion;
        return this;
    }
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("cognito_user_pool_client_id")
    public String cognitoUserPoolClientId;
    public LoginParameters withCognitoUserPoolClientId(String cognitoUserPoolClientId) {
        this.cognitoUserPoolClientId = cognitoUserPoolClientId;
        return this;
    }
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("cognito_user_pool_id")
    public String cognitoUserPoolId;
    public LoginParameters withCognitoUserPoolId(String cognitoUserPoolId) {
        this.cognitoUserPoolId = cognitoUserPoolId;
        return this;
    }
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("oauth_response_type")
    public LoginParametersOauthResponseTypeEnum oauthResponseType;
    public LoginParameters withOauthResponseType(LoginParametersOauthResponseTypeEnum oauthResponseType) {
        this.oauthResponseType = oauthResponseType;
        return this;
    }
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("organization_id")
    public String organizationId;
    public LoginParameters withOrganizationId(String organizationId) {
        this.organizationId = organizationId;
        return this;
    }
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("organization_name")
    public String organizationName;
    public LoginParameters withOrganizationName(String organizationName) {
        this.organizationName = organizationName;
        return this;
    }
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("organization_type")
    public String organizationType;
    public LoginParameters withOrganizationType(String organizationType) {
        this.organizationType = organizationType;
        return this;
    }
}