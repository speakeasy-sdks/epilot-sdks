package .models.shared;

import com.fasterxml.jackson.annotation.JsonProperty;
public class UserDetail {
    @JsonProperty("email")
    public String email;
    public UserDetail withEmail(String email) {
        this.email = email;
        return this;
    }
    @JsonProperty("full_name")
    public String fullName;
    public UserDetail withFullName(String fullName) {
        this.fullName = fullName;
        return this;
    }
    @JsonProperty("password")
    public String password;
    public UserDetail withPassword(String password) {
        this.password = password;
        return this;
    }
}