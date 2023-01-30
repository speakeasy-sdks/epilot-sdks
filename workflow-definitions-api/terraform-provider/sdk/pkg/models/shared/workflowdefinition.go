package shared

type WorkflowDefinition struct {
	AssignedTo             []string                 `json:"assignedTo,omitempty"`
	ClosingReasons         []ClosingReasonID        `json:"closingReasons,omitempty"`
	CreationTime           *string                  `json:"creationTime,omitempty"`
	Description            *string                  `json:"description,omitempty"`
	DueDate                *string                  `json:"dueDate,omitempty"`
	DynamicDueDate         *DynamicDueDate          `json:"dynamicDueDate,omitempty"`
	EnableEcpWorkflow      *bool                    `json:"enableECPWorkflow,omitempty"`
	Flow                   []interface{}            `json:"flow"`
	ID                     *string                  `json:"id,omitempty"`
	LastUpdateTime         *string                  `json:"lastUpdateTime,omitempty"`
	Name                   string                   `json:"name"`
	UpdateEntityAttributes []UpdateEntityAttributes `json:"updateEntityAttributes,omitempty"`
	UserIds                []float64                `json:"userIds,omitempty"`
}
