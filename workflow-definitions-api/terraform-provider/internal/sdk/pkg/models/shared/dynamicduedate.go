package shared

type DynamicDueDateActionTypeConditionEnum string

const (
	DynamicDueDateActionTypeConditionEnumWorkflowStarted DynamicDueDateActionTypeConditionEnum = "WORKFLOW_STARTED"
	DynamicDueDateActionTypeConditionEnumStepClosed      DynamicDueDateActionTypeConditionEnum = "STEP_CLOSED"
)

type DynamicDueDateTimePeriodEnum string

const (
	DynamicDueDateTimePeriodEnumDays   DynamicDueDateTimePeriodEnum = "days"
	DynamicDueDateTimePeriodEnumWeeks  DynamicDueDateTimePeriodEnum = "weeks"
	DynamicDueDateTimePeriodEnumMonths DynamicDueDateTimePeriodEnum = "months"
)

// DynamicDueDate
// set a Duedate for a step then a specific
type DynamicDueDate struct {
	ActionTypeCondition DynamicDueDateActionTypeConditionEnum `json:"actionTypeCondition"`
	NumberOfUnits       float64                               `json:"numberOfUnits"`
	StepID              *string                               `json:"stepId,omitempty"`
	TimePeriod          DynamicDueDateTimePeriodEnum          `json:"timePeriod"`
}
