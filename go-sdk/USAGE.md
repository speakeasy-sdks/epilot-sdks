<!-- Start SDK Example Usage -->
```go
package main

import (
    "openapi"
    "openapi/pkg/models/shared"
    "openapi/pkg/models/operations"
)

func main() {
    opts := []epilot.SDKOption{
        epilot.WithSecurity(
            shared.Security{
                EpilotAuth: shared.SchemeEpilotAuth{
                    Authorization: "Bearer YOUR_BEARER_TOKEN_HERE",
                },
            }
        ),
    }

    s := epilot.New(opts...)
    
    res, err := s.UserV1.GetMe(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.User != nil {
        // handle response
    }
```
<!-- End SDK Example Usage -->