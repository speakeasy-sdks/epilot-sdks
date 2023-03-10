# epilot-workflows-definition/sdk

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get epilot-workflows-definition/sdk
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import (
    "epilot-workflows-definition/sdk"
    "epilot-workflows-definition/sdk/pkg/models/shared"
    "epilot-workflows-definition/sdk/pkg/models/operations"
)

func main() {
    opts := []sdk.SDKOption{
        sdk.WithSecurity(
            shared.Security{
                BearerAuth: shared.SchemeBearerAuth{
                    Authorization: "Bearer YOUR_BEARER_TOKEN_HERE",
                },
            }
        ),
    }

    s := sdk.New(opts...)
    
    req := operations.ChangeReasonStatusRequest{
        PathParams: operations.ChangeReasonStatusPathParams{
            ReasonID: "sit",
        },
        Request: &shared.ChangeReasonStatusReq{
            Status: "ACTIVE",
        },
    }
    
    res, err := s.ClosingReason.ChangeReasonStatus(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## SDK Available Operations

### Closing Reason

* `ChangeReasonStatus` - changeReasonStatus
* `CreateClosingReason` - createClosingReason
* `GetAllClosingReasons` - getAllClosingReasons

### Workflows

* `CreateDefinition` - createDefinition
* `DeleteDefinition` - deleteDefinition
* `GetDefinition` - getDefinition
* `GetDefinitions` - getDefinitions
* `GetMaxAllowedLimit` - getMaxAllowedLimit
* `GetWorkflowClosingReasons` - getWorkflowClosingReasons
* `SetWorkflowClosingReasons` - setWorkflowClosingReasons
* `UpdateDefinition` - updateDefinition

<!-- End SDK Available Operations -->

### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
