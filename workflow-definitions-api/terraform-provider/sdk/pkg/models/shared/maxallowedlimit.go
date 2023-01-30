package shared

type MaxAllowedLimit struct {
	CurrentNoOfWorkflows *float64 `json:"currentNoOfWorkflows,omitempty"`
	MaxAllowed           *float64 `json:"maxAllowed,omitempty"`
}
