package shared

type StepRequirementConditionEnum string

const (
	StepRequirementConditionEnumClosed StepRequirementConditionEnum = "CLOSED"
)

// StepRequirement
// describe the requirement for step enablement
type StepRequirement struct {
	Condition    StepRequirementConditionEnum `json:"condition"`
	DefinitionID string                       `json:"definitionId"`
	Type         ItemTypeEnum                 `json:"type"`
}
