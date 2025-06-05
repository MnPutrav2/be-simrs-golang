package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/MnPutrav2/be-simrs-golang/internal/helper"
	"github.com/MnPutrav2/be-simrs-golang/internal/models"
	"github.com/MnPutrav2/be-simrs-golang/internal/pkg"
)

func CreatePatient(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
	// ---- Needed for every request ---
	if !pkg.CheckRequestHeader(w, r, sql, path, "POST") {
		return
	}

	// Check Header
	auth := r.Header.Get("Authorization")
	if !pkg.CheckAuthorization(w, path, sql, auth) {
		helper.ResponseError(w, "unauthorization", "unauthorization : 400", 401, path)
		return
	}

	split := strings.SplitN(auth, " ", 2)

	if len(split) != 2 || split[0] != "Bearer" {
		helper.ResponseError(w, "unauthorization error format", "unauthorization error format : 400", 400, path)
		return
	}
	// Check Header
	// --- ---

	// get client request body
	patient, err := helper.GetRequestBodyPatientData(w, r, path)
	if err != nil {
		helper.ResponseError(w, "empty request body", "empty request body : 400", 400, path)
		return
	}

	var checkExists bool
	err = sql.QueryRow("SELECT EXISTS(SELECT 1 FROM patients WHERE medical_record = ?)", patient.MedicalRecord).Scan(&checkExists)
	if err != nil {
		panic(err.Error())
	}

	if checkExists {
		helper.ResponseError(w, "duplicate entry", "duplicate entry : 400", 400, path)
		return
	}

	insert, err := sql.Exec("INSERT INTO patients(patients.medical_record, patients.name, patients.gender, patients.wedding, patients.religion, patients.education, patients.birth_place, patients.birth_date, patients.work, patients.address, patients.village, patients.district, patients.regencie, patients.province, patients.nik, patients.bpjs, patients.phone_number, patients.parent_name, patients.parent_gender) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", patient.MedicalRecord, patient.Name, patient.Gender, patient.Wedding, patient.Religion, patient.Education, patient.BirthPlace, patient.BirthDate, patient.Work, patient.Address, patient.Village, patient.District, patient.Regencie, patient.Province, patient.NIK, patient.BPJS, patient.PhoneNumber, patient.ParentName, patient.ParentGender)
	if err != nil {
		helper.ResponseError(w, "error server", err.Error()+" : 500", 500, path)
		return
	}

	_, err = insert.RowsAffected()
	if err != nil {
		helper.ResponseError(w, "failed insert data", "failed insert data : 400", 400, path)
		return
	}

	s, err := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "created"})
	if err != nil {
		helper.ResponseError(w, "error server", err.Error()+" : 500", 500, path)
		return
	}

	helper.ResponseSuccess(w, "create patient : 201", path, s, 201)
}

func GetPatient(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
	// ---- Needed for every request ---
	if !pkg.CheckRequestHeader(w, r, sql, path, "GET") {
		return
	}

	// Check Header
	auth := r.Header.Get("Authorization")
	if !pkg.CheckAuthorization(w, path, sql, auth) {
		helper.ResponseError(w, "unauthorization", "unauthorization : 400", 401, path)
		return
	}

	split := strings.SplitN(auth, " ", 2)

	if len(split) != 2 || split[0] != "Bearer" {
		helper.ResponseError(w, "unauthorization error format", "unauthorization error format : 400", 400, path)
		return
	}
	// Check Header
	// --- ---

	// get client request body
	param := r.URL.Query()

	limit := param.Get("limit")
	search := "%" + param.Get("search") + "%"

	datas, err := sql.Query("SELECT * FROM patients WHERE medical_record LIKE ? OR name LIKE ? LIMIT ?", search, search, limit)
	if err != nil {
		panic(err.Error())
	}

	var patients []models.PatientData

	for datas.Next() {
		var patient models.PatientData

		err := datas.Scan(&patient.MedicalRecord, &patient.Name, &patient.Gender, &patient.Wedding, &patient.Religion, &patient.Education, &patient.BirthPlace, &patient.BirthDate, &patient.Work, &patient.Address, &patient.Village, &patient.District, &patient.Regencie, &patient.Province, &patient.NIK, &patient.BPJS, &patient.PhoneNumber, &patient.ParentName, &patient.ParentGender)
		if err != nil {
			helper.ResponseError(w, "error server", err.Error()+" : 500", 500, path)
			return
		}

		patients = append(patients, patient)
	}

	s, err := json.Marshal(patients)
	if err != nil {
		panic(err.Error())
	}

	helper.ResponseSuccess(w, "get patient data : 200", path, s, 200)
}

func DeletePatient(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
	// ---- Needed for every request ---
	if !pkg.CheckRequestHeader(w, r, sql, path, "DELETE") {
		return
	}

	// Check Header
	auth := r.Header.Get("Authorization")
	if !pkg.CheckAuthorization(w, path, sql, auth) {
		helper.ResponseError(w, "unauthorization", "unauthorization : 400", 401, path)
		return
	}

	split := strings.SplitN(auth, " ", 2)

	if len(split) != 2 || split[0] != "Bearer" {
		helper.ResponseError(w, "unauthorization error format", "unauthorization error format : 400", 400, path)
		return
	}
	// Check Header
	// --- ---

	// get client request body
	param := r.URL.Query().Get("mr")

	_, err := sql.Exec("DELETE FROM patients WHERE medical_record = ?", param)
	if err != nil {
		panic(err.Error())
	}

	s, err := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "deleted"})
	if err != nil {
		panic(err.Error())
	}

	helper.ResponseSuccess(w, "delete patient data : 200", path, s, 200)
}
