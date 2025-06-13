package models

type CarePlantRequest struct {
	Encounter        string `json:"encounter"`
	CareNumber       string `json:"care_number"`
	PatientID        string `json:"patient_id"`
	PatientName      string `json:"patient_name"`
	PractitionerID   string `json:"practitioner_id"`
	PractitionerName string `json:"practitioner_name"`
	Description      string `json:"description"`
	Date             string `json:"date"`
	DateDisplay      string `json:"date_display"`
}

type CarePlan struct {
	ResourceType string     `json:"resourceType"`
	Identifier   Identifier `json:"identifier"`
	Title        string     `json:"title"`
	Status       string     `json:"status"`
	Category     []Category `json:"category"`
	Intent       string     `json:"intent"`
	Description  string     `json:"description"`
	Subject      Reference  `json:"subject"`
	Encounter    Encounter2 `json:"encounter"`
	Created      string     `json:"created"`
	Author       Reference  `json:"author"`
}

type Reference struct {
	Reference string `json:"reference"`
	Display   string `json:"display"`
}
