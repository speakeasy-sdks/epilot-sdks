<!-- Start SDK Example Usage -->
```typescript
import { Epilot, withSecurity} from "openapi";
import { GetMeResponse } from "openapi/src/sdk/models/operations";
import { AxiosError } from "axios";

const sdk = new Epilot(withSecurity(
  security: {
    epilotAuth: {
      authorization: "Bearer YOUR_BEARER_TOKEN_HERE",
    },
  }
));

sdk.userV1.getMe().then((res: GetMeResponse | AxiosError) => {
   // handle response
});
```
<!-- End SDK Example Usage -->