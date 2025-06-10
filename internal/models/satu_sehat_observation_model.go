package models

type ObservatioClientRequest struct {
	Encounter      string  `json:"encounter"`
	PatientID      string  `json:"patient_id"`
	PatientName    string  `json:"patient_name"`
	PractitionerID string  `json:"practitioner_id"`
	Date           string  `json:"date"`
	DateDisplay    string  `json:"date_display"`
	Pulse          int     `json:"pulse"`
	Respiratory    int     `json:"respiratory"`
	Temperature    float64 `json:"temperature"`
	Sistol         int     `json:"sistol"`
	Diastol        int     `json:"diastol"`
}

type ObservationCommon struct {
	ResourceType      string            `json:"resourceType"`
	Status            string            `json:"status"`
	Category          []Code            `json:"category"`
	Code              Code              `json:"code"`
	Subject           ServiceProvider   `json:"subject"`
	Performer         []ServiceProvider `json:"performer"`
	Encounter         Encounter2        `json:"encounter"`
	EffectiveDateTime string            `json:"effectiveDateTime"`
	Issued            string            `json:"issued"`
	ValueQuantity     ValueQuantity     `json:"valueQuantity"`
}

type ObservationTension struct {
	ResourceType      string            `json:"resourceType"`
	Status            string            `json:"status"`
	Category          []Code            `json:"category"`
	Code              Code              `json:"code"`
	Subject           Subject           `json:"subject"`
	Performer         []ServiceProvider `json:"performer"`
	Encounter         Encounter2        `json:"encounter"`
	EffectiveDateTime string            `json:"effectiveDateTime"`
	Issued            string            `json:"issued"`
	BodySite          BodySite          `json:"bodySite"`
	ValueQuantity     ValueQuantity     `json:"valueQuantity"`
	Interpretation    []Interpretation  `json:"interpretation"`
}

type ObservationTemperature struct {
	ResourceType      string            `json:"resourceType"`
	Status            string            `json:"status"`
	Category          []Code            `json:"category"`
	Code              Code              `json:"code"`
	Subject           ServiceProvider   `json:"subject"`
	Performer         []ServiceProvider `json:"performer"`
	Encounter         Encounter2        `json:"encounter"`
	EffectiveDateTime string            `json:"effectiveDateTime"`
	Issued            string            `json:"issued"`
	ValueQuantity     ValueQuantityTemp `json:"valueQuantity"`
	Interpretation    []Interpretation  `json:"interpretation"`
}

type ValueQuantity struct {
	Value  int64  `json:"value"`
	Unit   string `json:"unit"`
	System string `json:"system"`
	Code   string `json:"code"`
}

type ValueQuantityTemp struct {
	Value  float64 `json:"value"`
	Unit   string  `json:"unit"`
	System string  `json:"system"`
	Code   string  `json:"code"`
}

type Encounter2 struct {
	Reference string `json:"reference"`
	Display   string `json:"display"`
}

type BodySite struct {
	Coding []Coding `json:"coding"`
}

type Interpretation struct {
	Coding []Coding `json:"coding"`
	Text   string   `json:"text"`
}

type Performer struct {
	Reference string `json:"reference"`
}
