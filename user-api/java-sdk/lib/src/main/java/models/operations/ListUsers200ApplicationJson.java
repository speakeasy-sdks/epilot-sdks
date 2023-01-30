package .models.operations;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonInclude.Include;
public class ListUsers200ApplicationJson {
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("users")
    public .models.shared.User[] users;
    public ListUsers200ApplicationJson withUsers(.models.shared.User[] users) {
        this.users = users;
        return this;
    }
}