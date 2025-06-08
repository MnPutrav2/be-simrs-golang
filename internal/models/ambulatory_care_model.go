package models

type AmbulatoryCare struct {
	CareNumber      string `json:"care_number"`
	Date            string `json:"date"`
	BodyTemperature int    `json:"body_temperature"`
	Tension         string `json:"tension"`
	Pulse           int    `json:"pulse"`
	Respiration     int    `json:"respiration"`
	Height          int    `json:"height"`
	Weight          int    `json:"weight"`
	Spo2            int    `json:"spo2"`
	GCS             int    `json:"gcs"`
	Awareness       string `json:"awareness"`
	Complaint       string `json:"complaint"`
	Examination     string `json:"examination"`
	Allergy         string `json:"allergy"`
	FollowUp        string `json:"followup"`
	Assessment      string `json:"assessment"`
	Instructions    string `json:"instructions"`
	Evaluation      string `json:"evaluation"`
	Officer         int    `json:"officer"`
}
