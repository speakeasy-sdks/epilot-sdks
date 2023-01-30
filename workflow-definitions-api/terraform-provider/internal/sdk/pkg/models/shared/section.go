package shared

// Section
// A group of Steps that define the progress of the Workflow
type Section struct {
	ID    *string      `json:"id,omitempty"`
	Name  string       `json:"name"`
	Order float64      `json:"order"`
	Steps []Step       `json:"steps"`
	Type  ItemTypeEnum `json:"type"`
}
