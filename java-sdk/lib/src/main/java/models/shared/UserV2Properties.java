package .models.shared;

import com.fasterxml.jackson.annotation.JsonProperty;
public class UserV2Properties {
    @JsonProperty("name")
    public String name;
    public UserV2Properties withName(String name) {
        this.name = name;
        return this;
    }
    @JsonProperty("value")
    public String value;
    public UserV2Properties withValue(String value) {
        this.value = value;
        return this;
    }
}