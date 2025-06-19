package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/MnPutrav2/be-simrs-golang/internal/helper"
	"github.com/MnPutrav2/be-simrs-golang/internal/models"
	"github.com/MnPutrav2/be-simrs-golang/internal/pkg"
	"github.com/MnPutrav2/be-simrs-golang/internal/repository"
)

func CreatePatient(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string, m string) {
	// ---- Needed for every request ---
	if !pkg.CheckRequestHeader(w, r, sql, path, m) {
		return
	}

	// Check Header
	auth := r.Header.Get("Authorization")
	if !pkg.CheckAuthorization(w, path, sql, auth) {
		helper.ResponseWarn(w, "", "unauthorization", "unauthorization : 400", 401, path)
		return
	}

	split := strings.SplitN(auth, " ", 2)

	if len(split) != 2 || split[0] != "Bearer" {
		helper.ResponseWarn(w, "", "unauthorization error format", "unauthorization error format : 400", 400, path)
		return
	}
	// Check Header
	// --- ---

	// get client request body
	patient, err := helper.GetRequestBodyPatientData(w, r, path)
	if err != nil {
		helper.ResponseWarn(w, "", "invalid request body", err.Error(), 400, path)
		return
	}

	token := split[1]

	patientRepo := repository.NewPatientRepository(w, r, sql)
	if err := patientRepo.CreatePatientData(patient, token, path); err != nil {
		return
	}

	s, err := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "created"})
	if err != nil {
		helper.ResponseError(w, "", "error server", err.Error()+" : 500", 500, path)
		return
	}

	helper.ResponseSuccess(w, "", "create patient : 201", path, s, 201)
}

func GetPatient(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string, m string) {
	// ---- Needed for every request ---
	if !pkg.CheckRequestHeader(w, r, sql, path, m) {
		return
	}

	// Check Header
	auth := r.Header.Get("Authorization")
	if !pkg.CheckAuthorization(w, path, sql, auth) {
		helper.ResponseWarn(w, "", "unauthorization", "unauthorization", 401, path)
		return
	}

	split := strings.SplitN(auth, " ", 2)

	if len(split) != 2 || split[0] != "Bearer" {
		helper.ResponseWarn(w, "", "unauthorization error format", "unauthorization error format", 400, path)
		return
	}
	// Check Header
	// --- ---
	var id string
	if err := sql.QueryRow("SELECT users.id FROM users INNER JOIN session_token ON users.id = session_token.users_id WHERE session_token.token = $1", split[1]).Scan(&id); err != nil {
		return
	}

	// get client request body
	param := r.URL.Query()

	limit := param.Get("limit")
	search := "%" + param.Get("search") + "%"

	token := split[1]

	patientRepo := repository.NewPatientRepository(w, r, sql)
	patients, err := patientRepo.GetPatientData(limit, search, token, path)
	if err != nil {
		return
	}

	s, err := json.Marshal(patients)
	if err != nil {
		panic(err.Error())
	}

	helper.ResponseSuccess(w, id, "get patient data ", path, s, 200)
}

func UpdatePatientData(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string, m string) {
	// ---- Needed for every request ---
	if !pkg.CheckRequestHeader(w, r, sql, path, m) {
		return
	}

	// Check Header
	auth := r.Header.Get("Authorization")
	if !pkg.CheckAuthorization(w, path, sql, auth) {
		helper.ResponseWarn(w, "", "unauthorization", "unauthorization", 401, path)
		return
	}

	split := strings.SplitN(auth, " ", 2)

	if len(split) != 2 || split[0] != "Bearer" {
		helper.ResponseWarn(w, "", "unauthorization error format", "unauthorization error format", 400, path)
		return
	}
	// Check Header
	// --- ---
	var id string
	if err := sql.QueryRow("SELECT users.id FROM users INNER JOIN session_token ON users.id = session_token.users_id WHERE session_token.token = $1", split[1]).Scan(&id); err != nil {
		return
	}

	patient, err := helper.GetRequestBodyPatientDataUpdate(w, r, path)
	if err != nil {
		helper.ResponseWarn(w, id, "invalid request body", err.Error(), 400, path)
		return
	}

	patientRepo := repository.NewPatientRepository(w, r, sql)
	if err := patientRepo.UpdatePatientData(patient); err != nil {
		helper.ResponseWarn(w, id, "patient data not found", err.Error(), 404, path)
		return
	}

	s, err := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "updated"})
	if err != nil {
		helper.ResponseError(w, id, "error server", err.Error(), 500, path)
		return
	}

	helper.ResponseSuccess(w, id, "update patient data", path, s, 200)
}

func DeletePatient(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string, m string) {
	// ---- Needed for every request ---
	if !pkg.CheckRequestHeader(w, r, sql, path, m) {
		return
	}

	// Check Header
	auth := r.Header.Get("Authorization")
	if !pkg.CheckAuthorization(w, path, sql, auth) {
		helper.ResponseWarn(w, "", "unauthorization", "unauthorization", 401, path)
		return
	}

	split := strings.SplitN(auth, " ", 2)

	if len(split) != 2 || split[0] != "Bearer" {
		helper.ResponseWarn(w, "", "unauthorization error format", "unauthorization error format", 400, path)
		return
	}
	// Check Header
	// --- ---
	var id string
	if err := sql.QueryRow("SELECT users.id FROM users INNER JOIN session_token ON users.id = session_token.users_id WHERE session_token.token = $1", split[1]).Scan(&id); err != nil {
		return
	}

	// get client request body
	param := r.URL.Query().Get("mr")

	patientRepo := repository.NewPatientRepository(w, r, sql)
	if err := patientRepo.DeletePatientData(param); err != nil {
		panic(err.Error())
	}

	s, err := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "deleted"})
	if err != nil {
		panic(err.Error())
	}

	helper.ResponseSuccess(w, id, "delete patient data", path, s, 200)
}

func GetCurrentMR(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string, m string) {
	// ---- Needed for every request ---
	if !pkg.CheckRequestHeader(w, r, sql, path, m) {
		return
	}

	// Check Header
	auth := r.Header.Get("Authorization")
	if !pkg.CheckAuthorization(w, path, sql, auth) {
		helper.ResponseWarn(w, "", "unauthorization", "unauthorization", 401, path)
		return
	}

	split := strings.SplitN(auth, " ", 2)

	if len(split) != 2 || split[0] != "Bearer" {
		helper.ResponseWarn(w, "", "unauthorization error format", "unauthorization error format", 400, path)
		return
	}
	// Check Header
	// --- ---
	var id string
	if err := sql.QueryRow("SELECT users.id FROM users INNER JOIN session_token ON users.id = session_token.users_id WHERE session_token.token = $1", split[1]).Scan(&id); err != nil {
		return
	}

	patientRepo := repository.NewPatientRepository(w, r, sql)
	mr := patientRepo.GetCurrentMedicalRecord()

	s, err := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: mr})
	if err != nil {
		panic(err.Error())
	}

	helper.ResponseSuccess(w, id, "get current MR", path, s, 200)
}
