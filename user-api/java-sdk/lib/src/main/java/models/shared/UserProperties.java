package .models.shared;

import com.fasterxml.jackson.annotation.JsonProperty;
public class UserProperties {
    @JsonProperty("name")
    public String name;
    public UserProperties withName(String name) {
        this.name = name;
        return this;
    }
    @JsonProperty("value")
    public String value;
    public UserProperties withValue(String value) {
        this.value = value;
        return this;
    }
}