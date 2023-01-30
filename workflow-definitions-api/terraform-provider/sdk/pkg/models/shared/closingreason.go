package shared

// ClosingReason
// One Closing reason for a workflow
type ClosingReason struct {
	CreationTime   *string                  `json:"creationTime,omitempty"`
	ID             *string                  `json:"id,omitempty"`
	LastUpdateTime *string                  `json:"lastUpdateTime,omitempty"`
	Status         ClosingReasonsStatusEnum `json:"status"`
	Title          string                   `json:"title"`
}
