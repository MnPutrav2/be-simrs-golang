package models

type EmployeeData struct {
	Employee_ID string `json:"employee_id"`
	Name        string `json:"name"`
	Gender      string `json:"gender"`
	BirthPlace  string `json:"birth_place"`
	BirthDate   string `json:"birth_date"`
	Wedding     string `json:"wedding"`
	Address     string `json:"address"`
	Village     int    `json:"village"`
	District    int    `json:"district"`
	Regencie    int    `json:"regencie"`
	Province    int    `json:"province"`
	NIK         string `json:"nik"`
	BPJS        string `json:"bpjs"`
	NPWP        string `json:"npwp"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}
