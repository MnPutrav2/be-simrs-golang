package models

type Identifier struct {
	System string `json:"system"`
	Value  string `json:"value"`
}

type Class struct {
	System  string `json:"system"`
	Code    string `json:"code"`
	Display string `json:"display"`
}

type Subject struct {
	Reference string `json:"reference"`
	Display   string `json:"display"`
}

type Type struct {
	Coding []Class `json:"coding"`
}

type Participan struct {
	Individual Subject `json:"individual"`
	Type       []Type  `json:"type"`
}

type Period struct {
	Start string `json:"start"`
}

type Extension2 struct {
	URL                  string `json:"url"`
	ValueCodeableConcept Type   `json:"valueCodeableConcept"`
}

type Extension struct {
	URL       string       `json:"url"`
	Extension []Extension2 `json:"extension"`
}

type Location struct {
	Location  Subject     `json:"location"`
	Period    Period      `json:"period"`
	Extension []Extension `json:"extension"`
}

type StatusHistory struct {
	Status string `json:"status"`
	Period Period `json:"period"`
}

type ServiceProvider struct {
	Reference string `json:"reference"`
}

type EncounterBody struct {
	ResourceType    string          `json:"resourceType"`
	Identifier      []Identifier    `json:"identifier"`
	Status          string          `json:"status"`
	Class           Class           `json:"class"`
	Subject         Subject         `json:"subject"`
	Participant     []Participan    `json:"participant"`
	Period          Period          `json:"period"`
	Location        []Location      `json:"location"`
	StatusHistory   []StatusHistory `json:"statusHistory"`
	ServiceProvider ServiceProvider `json:"serviceProvider"`
}

type EncounterUpdate struct {
	ResourceType    string          `json:"resourceType"`
	ID              string          `json:"id"`
	Identifier      []Identifier    `json:"identifier"`
	Status          string          `json:"status"`
	Class           Class           `json:"class"`
	Subject         Subject         `json:"subject"`
	Participant     []Participan    `json:"participant"`
	Period          Period          `json:"period"`
	Location        []Location      `json:"location"`
	StatusHistory   []StatusHistory `json:"statusHistory"`
	ServiceProvider ServiceProvider `json:"serviceProvider"`
}

type EncounterResponse struct {
	CareNumber       string `json:"care_number"`
	PatientID        string `json:"patient_id"`
	PatientName      string `json:"patient_name"`
	PractitionerID   string `json:"practitioner_id"`
	PractitionerName string `json:"practitioner_name"`
	LocationID       string `json:"location_id"`
	Start            string `json:"start"`
}

type EncounterBodyResponse struct {
	Class           Class           `json:"class"`
	ID              string          `json:"id"`
	Identifier      []Identifier    `json:"identifier"`
	Location        []Location      `json:"location"`
	Meta            Meta            `json:"meta"`
	Participant     []Participant   `json:"participant"`
	Period          Period          `json:"period"`
	ResourceType    string          `json:"resourceType"`
	ServiceProvider ServiceProvider `json:"serviceProvider"`
	Status          string          `json:"status"`
	StatusHistory   []StatusHistory `json:"statusHistory"`
	Subject         Subject         `json:"subject"`
}

type LocationExtension struct {
	Extension []ExtensionExtension `json:"extension"`
	URL       string               `json:"url"`
}

type ExtensionExtension struct {
	URL                  string `json:"url"`
	ValueCodeableConcept Type   `json:"valueCodeableConcept"`
}

type Meta struct {
	LastUpdated string `json:"lastUpdated"`
	VersionID   string `json:"versionId"`
}

type Participant struct {
	Individual Subject `json:"individual"`
	Type       []Type  `json:"type"`
}
