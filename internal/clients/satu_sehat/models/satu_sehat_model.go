package models

type SatuSehatErrorResponse struct {
	ResourceType string  `json:"resourceType"`
	Issue        []Issue `json:"issue"`
}

type Issue struct {
	Severity string  `json:"severity"`
	Code     string  `json:"code"`
	Details  Details `json:"details"`
}

type Details struct {
	Text string `json:"text"`
}

type SatuSehatResponseReturn struct {
	Data string
	Code int
}

type SatuSehatToClientResponse struct {
	Status   string `json:"status"`
	Response string `json:"response"`
}
