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