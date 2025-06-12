package models

type ConditionBodyResponse struct {
	Category       []ClinicalStatus `json:"category"`
	ClinicalStatus ClinicalStatus   `json:"clinicalStatus"`
	Code           ClinicalStatus   `json:"code"`
	Encounter      Encounter        `json:"encounter"`
	ID             string           `json:"id"`
	Identifier     []interface{}    `json:"identifier"`
	Meta           Meta             `json:"meta"`
	OnsetDateTime  string           `json:"onsetDateTime"`
	RecordedDate   string           `json:"recordedDate"`
	ResourceType   string           `json:"resourceType"`
	Subject        Subject          `json:"subject"`
}

type ConditionRequest struct {
	ResourceType   string         `json:"resourceType"`
	ClinicalStatus ClinicalStatus `json:"clinicalStatus"`
	Category       []Category     `json:"category"`
	Code           Code           `json:"code"`
	Subject        Subject        `json:"subject"`
	Encounter      Encounter      `json:"encounter"`
	OnSetDateTime  string         `json:"onsetDateTime"`
	RecordedDate   string         `json:"recordedDate"`
}

type Diagnosis struct {
	ICD    int    `json:"icd"`
	Code   string `json:"code"`
	Detail string `json:"detail"`
}

type ConditionClientRequest struct {
	Encounter   string   `json:"encounter"`
	PatientID   string   `json:"patient_id"`
	PatientName string   `json:"patient_name"`
	Date        string   `json:"date"`
	Diagnosis   []Coding `json:"diagnosis"`
}

type ClinicalStatus struct {
	Coding []Coding `json:"coding"`
}

type Category struct {
	Coding []Coding `json:"coding"`
}

type Encounter struct {
	Reference string `json:"reference"`
}

type Code struct {
	Coding []Coding `json:"coding"`
}

type Note struct {
	Text string `json:"text"`
}

type Coding struct {
	System  string `json:"system"`
	Code    string `json:"code"`
	Display string `json:"display"`
}
