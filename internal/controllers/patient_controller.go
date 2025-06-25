package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

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

	val := pkg.CheckUserLogin(w, r, sql, path)
	if val.Status == "authorization" {
		helper.ResponseWarn(w, "", "unauthorization", "unauthorization", 401, path)
		return
	} else if val.Status == "error_format" {
		helper.ResponseWarn(w, "", "unauthorization error format", "unauthorization error format", 400, path)
		return
	}

	// get client request body
	patient, err := helper.GetRequestBodyPatientData(w, r, path)
	if err != nil {
		helper.ResponseWarn(w, val.Id, "invalid request body", err.Error(), 400, path)
		return
	}

	patientRepo := repository.NewPatientRepository(w, r, sql)
	if err := patientRepo.CreatePatientData(patient, val.Id, path); err != nil {
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

	val := pkg.CheckUserLogin(w, r, sql, path)
	if val.Status == "authorization" {
		helper.ResponseWarn(w, "", "unauthorization", "unauthorization", 401, path)
		return
	} else if val.Status == "error_format" {
		helper.ResponseWarn(w, "", "unauthorization error format", "unauthorization error format", 400, path)
		return
	}

	// get client request body
	param := r.URL.Query()

	limit := param.Get("limit")
	search := "%" + param.Get("search") + "%"

	patientRepo := repository.NewPatientRepository(w, r, sql)
	patients, err := patientRepo.GetPatientData(limit, search, val.Token, path)
	if err != nil {
		return
	}

	s, err := json.Marshal(patients)
	if err != nil {
		panic(err.Error())
	}

	helper.ResponseSuccess(w, val.Id, "get patient data ", path, s, 200)
}

func UpdatePatientData(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string, m string) {
	// ---- Needed for every request ---
	if !pkg.CheckRequestHeader(w, r, sql, path, m) {
		return
	}

	val := pkg.CheckUserLogin(w, r, sql, path)
	if val.Status == "authorization" {
		helper.ResponseWarn(w, "", "unauthorization", "unauthorization", 401, path)
		return
	} else if val.Status == "error_format" {
		helper.ResponseWarn(w, "", "unauthorization error format", "unauthorization error format", 400, path)
		return
	}

	patient, err := helper.GetRequestBodyPatientDataUpdate(w, r, path)
	if err != nil {
		helper.ResponseWarn(w, val.Id, "invalid request body", err.Error(), 400, path)
		return
	}

	patientRepo := repository.NewPatientRepository(w, r, sql)
	if err := patientRepo.UpdatePatientData(patient); err != nil {
		helper.ResponseWarn(w, val.Id, "patient data not found", err.Error(), 404, path)
		return
	}

	s, err := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "updated"})
	if err != nil {
		helper.ResponseError(w, val.Id, "error server", err.Error(), 500, path)
		return
	}

	helper.ResponseSuccess(w, val.Id, "update patient data", path, s, 200)
}

func DeletePatient(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string, m string) {
	// ---- Needed for every request ---
	if !pkg.CheckRequestHeader(w, r, sql, path, m) {
		return
	}

	val := pkg.CheckUserLogin(w, r, sql, path)
	if val.Status == "authorization" {
		helper.ResponseWarn(w, "", "unauthorization", "unauthorization", 401, path)
		return
	} else if val.Status == "error_format" {
		helper.ResponseWarn(w, "", "unauthorization error format", "unauthorization error format", 400, path)
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

	helper.ResponseSuccess(w, val.Id, "delete patient data", path, s, 200)
}

func GetCurrentMR(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string, m string) {
	// ---- Needed for every request ---
	if !pkg.CheckRequestHeader(w, r, sql, path, m) {
		return
	}

	val := pkg.CheckUserLogin(w, r, sql, path)
	if val.Status == "authorization" {
		helper.ResponseWarn(w, "", "unauthorization", "unauthorization", 401, path)
		return
	} else if val.Status == "error_format" {
		helper.ResponseWarn(w, "", "unauthorization error format", "unauthorization error format", 400, path)
		return
	}

	patientRepo := repository.NewPatientRepository(w, r, sql)
	mr := patientRepo.GetCurrentMedicalRecord()

	s, err := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: mr})
	if err != nil {
		panic(err.Error())
	}

	helper.ResponseSuccess(w, val.Id, "get current MR", path, s, 200)
}
