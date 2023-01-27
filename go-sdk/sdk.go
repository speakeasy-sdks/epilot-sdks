package epilot

import (
	"net/http"

	"openapi/pkg/models/shared"
	"openapi/pkg/utils"
)

var ServerList = []string{
	"https://user.sls.epilot.io",
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Epilot struct {
	UserV1 *UserV1
	UserV2 *UserV2

	_defaultClient  HTTPClient
	_securityClient HTTPClient
	_security       *shared.Security
	_serverURL      string
	_language       string
	_sdkVersion     string
	_genVersion     string
}

type SDKOption func(*Epilot)

func WithServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *Epilot) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk._serverURL = serverURL
	}
}

func WithClient(client HTTPClient) SDKOption {
	return func(sdk *Epilot) {
		sdk._defaultClient = client
	}
}

func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *Epilot) {
		sdk._security = &security
	}
}

func New(opts ...SDKOption) *Epilot {
	sdk := &Epilot{
		_language:   "go",
		_sdkVersion: "",
		_genVersion: "0.21.0",
	}
	for _, opt := range opts {
		opt(sdk)
	}

	if sdk._defaultClient == nil {
		sdk._defaultClient = http.DefaultClient
	}
	if sdk._securityClient == nil {

		if sdk._security != nil {
			sdk._securityClient = utils.ConfigureSecurityClient(sdk._defaultClient, sdk._security)
		} else {
			sdk._securityClient = sdk._defaultClient
		}

	}

	if sdk._serverURL == "" {
		sdk._serverURL = ServerList[0]
	}

	sdk.UserV1 = NewUserV1(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.UserV2 = NewUserV2(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	return sdk
}
