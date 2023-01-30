package shared

type StepAutomationConfig struct {
	FlowID string `json:"flowId"`
}

// Step
// Action that needs to be done in a Workflow
type Step struct {
	AssignedTo       []string              `json:"assignedTo,omitempty"`
	AutomationConfig *StepAutomationConfig `json:"automationConfig,omitempty"`
	DueDate          *string               `json:"dueDate,omitempty"`
	DynamicDueDate   *DynamicDueDate       `json:"dynamicDueDate,omitempty"`
	Ecp              *EcpDetails           `json:"ecp,omitempty"`
	ExecutionType    *StepTypeEnum         `json:"executionType,omitempty"`
	ID               *string               `json:"id,omitempty"`
	Name             string                `json:"name"`
	Order            float64               `json:"order"`
	Requirements     []StepRequirement     `json:"requirements,omitempty"`
	Type             ItemTypeEnum          `json:"type"`
	UserIds          []float64             `json:"userIds,omitempty"`
}
