package models

type RequestAmbulatoryCare struct {
	CareNumber      string `json:"care_number"`
	MedicalRecord   string `json:"medical_record"`
	Name            string `json:"name"`
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
	Officer         string `json:"officer"`
}

type RequestUpdateAmbulatorCare struct {
	CareNumber string                `json:"care_number"`
	Date       string                `json:"date"`
	Data       RequestAmbulatoryCare `json:"data"`
}

type ResponseAmbulatoryCare struct {
	CareNumber      string `json:"care_number"`
	MedicalRecord   string `json:"medical_record"`
	Name            string `json:"name"`
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
	Officer         string `json:"officer"`
	OfficerName     string `json:"officer_name"`
}
