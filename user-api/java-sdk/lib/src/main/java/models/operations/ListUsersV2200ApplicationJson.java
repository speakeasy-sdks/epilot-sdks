package .models.operations;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonInclude.Include;
public class ListUsersV2200ApplicationJson {
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("results")
    public .models.shared.UserV2[] results;
    public ListUsersV2200ApplicationJson withResults(.models.shared.UserV2[] results) {
        this.results = results;
        return this;
    }
}