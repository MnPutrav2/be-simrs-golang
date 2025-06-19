package repository

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/MnPutrav2/be-simrs-golang/internal/models"
)

type PatientRepository interface {
	CreatePatientData(patient models.PatientData, token string, path string) error
	GetPatientData(limit string, search string, token string, path string) ([]models.PatientData, error)
	DeletePatientData(mr string) error
	UpdatePatientData(patient models.PatientDataUpdate) error
	GetCurrentMedicalRecord() string
}

type patientRepository struct {
	w   http.ResponseWriter
	r   *http.Request
	sql *sql.DB
}

func NewPatientRepository(w http.ResponseWriter, r *http.Request, sql *sql.DB) PatientRepository {
	return &patientRepository{w, r, sql}
}

func (q *patientRepository) CreatePatientData(patient models.PatientData, token string, path string) error {
	var checkExists bool
	err := q.sql.QueryRow("SELECT EXISTS(SELECT 1 FROM patients WHERE medical_record = $1)", patient.MedicalRecord).Scan(&checkExists)
	if err != nil {
		panic(err.Error())
	}

	if checkExists {
		return fmt.Errorf("duplicate entry")
	}

	insert, err := q.sql.Exec("INSERT INTO patients(medical_record, name, gender, wedding, religion, education, birth_place, birth_date, work, address, village, district, regencie, province, nik, bpjs, phone_number, parent_name, relationship, parent_gender) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20);", patient.MedicalRecord, patient.Name, patient.Gender, patient.Wedding, patient.Religion, patient.Education, patient.BirthPlace, patient.BirthDate, patient.Work, patient.Address, patient.Village, patient.District, patient.Regencie, patient.Province, patient.NIK, patient.BPJS, patient.PhoneNumber, patient.ParentName, patient.Relationship, patient.ParentGender)
	if err != nil {
		panic(err.Error())
	}

	_, err = insert.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (q *patientRepository) GetPatientData(limit string, search string, token string, path string) ([]models.PatientData, error) {
	datas, err := q.sql.Query("SELECT * FROM patients WHERE medical_record LIKE $1 OR name LIKE $2 ORDER BY medical_record DESC LIMIT $3 ", search, search, limit)
	if err != nil {
		panic(err.Error())
	}

	var patients []models.PatientData

	for datas.Next() {
		var patient models.PatientData

		err := datas.Scan(&patient.MedicalRecord, &patient.Name, &patient.Gender, &patient.Wedding, &patient.Religion, &patient.Education, &patient.BirthPlace, &patient.BirthDate, &patient.Work, &patient.Address, &patient.Village, &patient.District, &patient.Regencie, &patient.Province, &patient.NIK, &patient.BPJS, &patient.PhoneNumber, &patient.ParentName, &patient.Relationship, &patient.ParentGender)
		if err != nil {
			return []models.PatientData{}, err
		}

		patients = append(patients, patient)
	}

	return patients, nil

}

func (q *patientRepository) UpdatePatientData(datas models.PatientDataUpdate) error {
	mr := datas.MedicalRecordKey
	patient := datas.Update

	_, err := q.sql.Exec("UPDATE patients SET medical_record = $1, name = $2, gender = $3, wedding = $4, religion = $5, education = $6, birth_place = $7, birth_date = $8, work = $9, address = $10, village = $11, district = $12, regencie = $13, province = $14, nik = $15, bpjs = $16, phone_number = $17, parent_name = $18, relationship = $19, parent_gender = $20 WHERE medical_record = $21;", patient.MedicalRecord, patient.Name, patient.Gender, patient.Wedding, patient.Religion, patient.Education, patient.BirthPlace, patient.BirthDate, patient.Work, patient.Address, patient.Village, patient.District, patient.Regencie, patient.Province, patient.NIK, patient.BPJS, patient.PhoneNumber, patient.ParentName, patient.Relationship, patient.ParentGender, mr)
	return err
}

func (q *patientRepository) DeletePatientData(mr string) error {
	_, err := q.sql.Exec("DELETE FROM patients WHERE medical_record = $1", mr)
	return err
}

func (q *patientRepository) GetCurrentMedicalRecord() string {
	var mr string

	err := q.sql.QueryRow("SELECT COUNT(medical_record) + 1 FROM patients").Scan(&mr)
	if err != nil {
		panic(err.Error())
	}

	return mr
}
