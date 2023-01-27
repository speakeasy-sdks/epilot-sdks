package .models.shared;


public enum UserInvitationPayloadLanguageEnum {
    EN("en"),
    DE("de");

    public final String value;

    private UserInvitationPayloadLanguageEnum(String value) {
        this.value = value;
    }
}
