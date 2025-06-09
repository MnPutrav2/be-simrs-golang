package repository

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/MnPutrav2/be-simrs-golang/internal/helper"
	"github.com/MnPutrav2/be-simrs-golang/internal/models"
)

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
	err := q.sql.QueryRow("SELECT EXISTS(SELECT 1 FROM patients WHERE medical_record = ?)", patient.MedicalRecord).Scan(&checkExists)
	if err != nil {
		panic(err.Error())
	}

	if checkExists {
		helper.ResponseError(q.w, "duplicate entry", "duplicate entry : 400", 400, path)
		return fmt.Errorf("duplicate entry")
	}

	insert, err := q.sql.Exec("INSERT INTO patients(patients.medical_record, patients.name, patients.gender, patients.wedding, patients.religion, patients.education, patients.birth_place, patients.birth_date, patients.work, patients.address, patients.village, patients.district, patients.regencie, patients.province, patients.nik, patients.bpjs, patients.phone_number, patients.parent_name, patients.relationship, patients.parent_gender) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", patient.MedicalRecord, patient.Name, patient.Gender, patient.Wedding, patient.Religion, patient.Education, patient.BirthPlace, patient.BirthDate, patient.Work, patient.Address, patient.Village, patient.District, patient.Regencie, patient.Province, patient.NIK, patient.BPJS, patient.PhoneNumber, patient.ParentName, patient.Relationship, patient.ParentGender)
	if err != nil {
		helper.ResponseError(q.w, "error server", err.Error()+" : 500", 500, path)
		return err
	}

	_, err = insert.RowsAffected()
	if err != nil {
		helper.ResponseError(q.w, "failed insert data", "failed insert data : 400", 400, path)
		return err
	}

	return nil
}

func (q *patientRepository) GetPatientData(limit string, search string, token string, path string) ([]models.PatientData, error) {
	datas, err := q.sql.Query("SELECT * FROM patients WHERE medical_record LIKE ? OR name LIKE ? ORDER BY medical_record DESC LIMIT ? ", search, search, limit)
	if err != nil {
		panic(err.Error())
	}

	var patients []models.PatientData

	for datas.Next() {
		var patient models.PatientData

		err := datas.Scan(&patient.MedicalRecord, &patient.Name, &patient.Gender, &patient.Wedding, &patient.Religion, &patient.Education, &patient.BirthPlace, &patient.BirthDate, &patient.Work, &patient.Address, &patient.Village, &patient.District, &patient.Regencie, &patient.Province, &patient.NIK, &patient.BPJS, &patient.PhoneNumber, &patient.ParentName, &patient.Relationship, &patient.ParentGender)
		if err != nil {
			helper.ResponseError(q.w, "error server", err.Error()+" : 500", 500, path)
			return []models.PatientData{}, err
		}

		patients = append(patients, patient)
	}

	return patients, nil

}

func (q *patientRepository) UpdatePatientData(datas models.PatientDataUpdate) error {
	mr := datas.MedicalRecordKey
	patient := datas.Update

	_, err := q.sql.Exec("UPDATE patients SET medical_record = ?, name = ?, gender = ?, wedding = ?, religion = ?, education = ?, birth_place = ?, birth_date = ?, work = ?, address = ?, village = ?, district = ?, regencie = ?, province = ?, nik = ?, bpjs = ?, phone_number = ?, parent_name = ?, relationship = ?, parent_gender = ? WHERE medical_record = ?", patient.MedicalRecord, patient.Name, patient.Gender, patient.Wedding, patient.Religion, patient.Education, patient.BirthPlace, patient.BirthDate, patient.Work, patient.Address, patient.Village, patient.District, patient.Regencie, patient.Province, patient.NIK, patient.BPJS, patient.PhoneNumber, patient.ParentName, patient.Relationship, patient.ParentGender, mr)
	return err
}

func (q *patientRepository) DeletePatientData(mr string) error {
	_, err := q.sql.Exec("DELETE FROM patients WHERE medical_record = ?", mr)
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
