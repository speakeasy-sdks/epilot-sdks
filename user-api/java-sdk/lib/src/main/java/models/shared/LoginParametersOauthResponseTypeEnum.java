package .models.shared;


public enum LoginParametersOauthResponseTypeEnum {
    CODE("code"),
    TOKEN("token");

    public final String value;

    private LoginParametersOauthResponseTypeEnum(String value) {
        this.value = value;
    }
}
