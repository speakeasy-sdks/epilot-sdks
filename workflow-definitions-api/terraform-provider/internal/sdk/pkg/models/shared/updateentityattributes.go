package shared

type UpdateEntityAttributesSourceEnum string

const (
	UpdateEntityAttributesSourceEnumWorkflowStatus UpdateEntityAttributesSourceEnum = "workflow_status"
	UpdateEntityAttributesSourceEnumCurrentSection UpdateEntityAttributesSourceEnum = "current_section"
	UpdateEntityAttributesSourceEnumCurrentStep    UpdateEntityAttributesSourceEnum = "current_step"
)

type UpdateEntityAttributesTarget struct {
	EntityAttribute string `json:"entityAttribute"`
	EntitySchema    string `json:"entitySchema"`
}

type UpdateEntityAttributes struct {
	Source UpdateEntityAttributesSourceEnum `json:"source"`
	Target UpdateEntityAttributesTarget     `json:"target"`
}
