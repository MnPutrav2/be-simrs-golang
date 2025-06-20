package repository

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/MnPutrav2/be-simrs-golang/internal/helper"
	"github.com/MnPutrav2/be-simrs-golang/internal/models"
)

type RegisterPatient interface {
	CreateRegistrationData(reg models.RequestRegisterPatient, path string) error
	DeleteRegistrationData(treatmentNumber string) error
	GetRegistrationData(date1 string, date2 string, limit int, search string) ([]models.ResponseRegisterPatient, error)
	GetCurrentRegisterNumber(date string, policlinic string) (int, error)
	GetCurrentCareNumber(date string) (int, error)
}

type registerRepository struct {
	sql *sql.DB
	w   http.ResponseWriter
	r   *http.Request
}

func NewRegisterRepository(sql *sql.DB, w http.ResponseWriter, r *http.Request) RegisterPatient {
	return &registerRepository{sql, w, r}
}

func (q *registerRepository) CreateRegistrationData(reg models.RequestRegisterPatient, path string) error {
	var check bool
	err := q.sql.QueryRow("SELECT EXISTS(SELECT 1 FROM registration WHERE care_number = $1)", reg.CareNumber).Scan(&check)
	if err != nil {
		return err
	}

	if check {
		helper.ResponseError(q.w, "", "duplicate entry", "duplicate entry : 400", 400, path)
		return fmt.Errorf("duplicate entry")
	}

	_, err = q.sql.Exec("INSERT INTO registration(care_number, register_number, register_date, medical_record, payment_method, policlinic, doctor) VALUES($1, $2, $3, $4, $5, $6, $7);", reg.CareNumber, reg.RegisterNumber, reg.RegisterDate, reg.MedicalRecord, reg.PaymentMethod, reg.Policlinic, reg.Doctor)
	if err != nil {
		helper.ResponseError(q.w, "", "error request data : check your data", err.Error(), 400, path)
		return fmt.Errorf("error request data")
	}

	return nil
}

func (q *registerRepository) GetRegistrationData(date1 string, date2 string, limit int, search string) ([]models.ResponseRegisterPatient, error) {
	result, err := q.sql.Query("SELECT registration.care_number, registration.register_number, registration.register_date, registration.medical_record, patients.name, patients.gender, registration.payment_method, registration.policlinic, policlinic.name, registration.doctor, doctors.name FROM registration INNER JOIN patients ON registration.medical_record = patients.medical_record INNER JOIN policlinic ON registration.policlinic = policlinic.id INNER JOIN doctors ON registration.doctor = doctors.id WHERE registration.register_date BETWEEN $1 AND $2 AND (registration.care_number LIKE $3 OR patients.name LIKE $4) ORDER BY registration.care_number DESC LIMIT $5", date1, date2, search, search, limit)
	if err != nil {
		return nil, err
	}

	var datas []models.ResponseRegisterPatient

	for result.Next() {
		var dt models.ResponseRegisterPatient

		err := result.Scan(&dt.CareNumber, &dt.RegisterNumber, &dt.RegisterDate, &dt.MedicalRecord, &dt.Name, &dt.Gender, &dt.PaymentMethod, &dt.Policlinic_id, &dt.Policlinic_name, &dt.Doctor_id, &dt.Doctor_name)
		if err != nil {
			return nil, err
		}

		datas = append(datas, dt)
	}

	return datas, nil
}

func (q *registerRepository) DeleteRegistrationData(tn string) error {
	_, err := q.sql.Exec("DELETE FROM registration WHERE care_number = $1", tn)
	return err
}

func (q *registerRepository) GetCurrentRegisterNumber(date string, policlinic string) (int, error) {
	var data int

	err := q.sql.QueryRow("SELECT COUNT(*) FROM registration WHERE registration.register_date = $1 AND registration.policlinic = $2", date, policlinic).Scan(&data)
	if err != nil {
		return 0, err
	}

	return data, nil
}

func (q *registerRepository) GetCurrentCareNumber(date string) (int, error) {
	var data int

	err := q.sql.QueryRow("SELECT COUNT(*) FROM registration WHERE registration.register_date = $1", date).Scan(&data)
	if err != nil {
		return 0, err
	}

	return data, nil
}
