package .models.shared;


public enum OrganizationTypeEnum {
    VENDOR("Vendor"),
    PARTNER("Partner");

    public final String value;

    private OrganizationTypeEnum(String value) {
        this.value = value;
    }
}
