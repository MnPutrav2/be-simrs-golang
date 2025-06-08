package repository

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/MnPutrav2/be-simrs-golang/internal/helper"
	"github.com/MnPutrav2/be-simrs-golang/internal/models"
)

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
	err := q.sql.QueryRow("SELECT EXISTS(SELECT 1 FROM registration WHERE treatment_number = ?)", reg.TreatmentNumber).Scan(&check)
	if err != nil {
		panic(err.Error())
	}

	if check {
		helper.ResponseError(q.w, "duplicate entry", "duplicate entry : 400", 400, path)
		return fmt.Errorf("duplicate entry")
	}

	_, err = q.sql.Exec("INSERT INTO registration(treatment_number, register_number, register_date, medical_record, payment_method, policlinic, doctor) VALUES(?, ?, ?, ?, ?, ?, ?);", reg.TreatmentNumber, reg.RegisterNumber, reg.RegisterDate, reg.MedicalRecord, reg.PaymentMethod, reg.Policlinic, reg.Doctor)
	if err != nil {
		helper.ResponseError(q.w, "error request data : check your data", "error data : 400", 400, path)
		return fmt.Errorf("error request data")
	}

	return nil
}

func (q *registerRepository) GetRegistrationData(date1 string, date2 string, limit string, search string) ([]models.ResponseRegisterPatient, error) {
	result, err := q.sql.Query("SELECT registration.treatment_number, registration.register_number, registration.register_date, registration.medical_record, patients.name, patients.gender, registration.payment_method, registration.policlinic, policlinics.name, registration.doctor, doctors.name FROM registration INNER JOIN patients ON registration.medical_record = patients.medical_record INNER JOIN policlinics ON registration.policlinic = policlinics.id INNER JOIN doctors ON registration.doctor = doctors.id WHERE registration.register_date BETWEEN ? AND ? AND (registration.treatment_number LIKE ? OR patients.name LIKE ?) ORDER BY registration.treatment_number DESC LIMIT ?", date1, date2, search, search, limit)
	if err != nil {
		return []models.ResponseRegisterPatient{}, nil
	}

	var datas []models.ResponseRegisterPatient

	for result.Next() {
		var dt models.ResponseRegisterPatient

		err := result.Scan(&dt.TreatmentNumber, &dt.RegisterNumber, &dt.RegisterDate, &dt.MedicalRecord, &dt.Name, &dt.Gender, &dt.PaymentMethod, &dt.Policlinic_id, &dt.Policlinic_name, &dt.Doctor_id, &dt.Doctor_name)
		if err != nil {
			panic(err.Error())
		}

		datas = append(datas, dt)
	}

	return datas, nil
}

func (q *registerRepository) DeleteRegistrationData(tn string) error {
	_, err := q.sql.Exec("DELETE FROM registration WHERE treatment_number = ?", tn)
	return err
}
