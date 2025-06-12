package models

type ClinicalImpressionClientRequest struct {
	Encounter      string   `json:"encounter"`
	Condition      string   `json:"condition"`
	PatientID      string   `json:"patient_id"`
	PatientName    string   `json:"patient_name"`
	PractitionerID string   `json:"practitioner_id"`
	Date           string   `json:"date"`
	DateDisplay    string   `json:"date_display"`
	Description    string   `json:"description"`
	Summary        string   `json:"summary"`
	Diagnosis      []Coding `json:"diagnosis"`
}

type ClinicalImpressionRequest struct {
	ResourceType             string            `json:"resourceType"`
	Status                   string            `json:"status"`
	Description              string            `json:"description"`
	Subject                  Encounter2        `json:"subject"`
	Encounter                Encounter2        `json:"encounter"`
	EffectiveDateTime        string            `json:"effectiveDateTime"`
	Date                     string            `json:"date"`
	Assessor                 Assessor          `json:"assessor"`
	Summary                  string            `json:"summary"`
	Finding                  []Finding         `json:"finding"`
	PrognosisCodeableConcept []CodeableConcept `json:"prognosisCodeableConcept"`
}

type Assessor struct {
	Reference string `json:"reference"`
}

type Finding struct {
	ItemCodeableConcept CodeableConcept `json:"itemCodeableConcept"`
	ItemReference       Assessor        `json:"itemReference"`
}

type CodeableConcept struct {
	Coding []Coding `json:"coding"`
}
