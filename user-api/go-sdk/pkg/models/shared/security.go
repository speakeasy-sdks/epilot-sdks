package shared

type SchemeEpilotAuth struct {
	Authorization string `security:"name=Authorization"`
}

type Security struct {
	EpilotAuth SchemeEpilotAuth `security:"scheme,type=http,subtype=bearer"`
}
