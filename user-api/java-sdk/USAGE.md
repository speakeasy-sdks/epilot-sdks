<!-- Start SDK Example Usage -->
```java
package hello.world;

import .Epilot;
import .models.shared.Security;

public class Application {
    public static void main(String[] args) {
        try {
            Epilot.Builder builder = Epilot.builder();

            builder.setSecurity(
                new Security() {{
                    epilotAuth = new SchemeEpilotAuth() {{
                        authorization = "Bearer YOUR_BEARER_TOKEN_HERE";
                    }};
                }}
            );

            Epilot sdk = builder.build();

            GetMeResponse res = sdk.userV1.getMe();

            if (res.user.isPresent()) {
                // handle response
            }
        } catch (Exception e) {
            // handle exception
        }
```
<!-- End SDK Example Usage -->