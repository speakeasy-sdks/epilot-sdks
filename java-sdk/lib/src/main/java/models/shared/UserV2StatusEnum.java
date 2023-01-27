package .models.shared;


public enum UserV2StatusEnum {
    ACTIVE("Active"),
    PENDING("Pending"),
    DEACTIVATED("Deactivated"),
    DELETED("Deleted");

    public final String value;

    private UserV2StatusEnum(String value) {
        this.value = value;
    }
}
