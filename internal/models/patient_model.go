package models

type PatientData struct {
	MedicalRecord string `json:"medical_record"`
	Name          string `json:"name"`
	Gender        string `json:"gender"`
	Wedding       string `json:"wedding"`
	Religion      string `json:"religion"`
	Education     string `json:"education"`
	BirthPlace    string `json:"birth_place"`
	BirthDate     string `json:"birth_date"`
	Work          string `json:"work"`
	Address       string `json:"address"`
	Village       int    `json:"village" validate:"max=4"`
	District      int    `json:"district" validate:"max=2"`
	Regencie      int    `json:"regencie" validate:"max=2"`
	Province      int    `json:"province" validate:"max=2"`
	NIK           string `json:"nik"`
	BPJS          string `json:"bpjs"`
	PhoneNumber   string `json:"phone_number"`
	ParentName    string `json:"parent_name"`
	ParentGender  string `json:"parent_gender"`
}

type PatientDataUpdate struct {
	MedicalRecordKey string      `json:"medical_record"`
	Update           PatientData `json:"update"`
}
