package .models.shared;


public enum SignupUserPayloadLanguageEnum {
    EN("en"),
    DE("de");

    public final String value;

    private SignupUserPayloadLanguageEnum(String value) {
        this.value = value;
    }
}
