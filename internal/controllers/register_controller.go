package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/MnPutrav2/be-simrs-golang/internal/helper"
	"github.com/MnPutrav2/be-simrs-golang/internal/models"
	"github.com/MnPutrav2/be-simrs-golang/internal/pkg"
	"github.com/MnPutrav2/be-simrs-golang/internal/repository"
)

func CreateRegistrationPatient(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string, m string) {
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
	body, err := helper.GetRequestBodyRegisterData(w, r, path)
	if err != nil {
		helper.ResponseError(w, val.Id, "invalid request body", err.Error(), 400, path)
		return
	}

	registerRepo := repository.NewRegisterRepository(sql, w, r)
	err = registerRepo.CreateRegistrationData(body, path)
	if err != nil {
		return
	}

	s, _ := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "created"})
	helper.ResponseSuccess(w, val.Id, "create patient registration", path, s, 201)
}

func DeleteRegistrationPatient(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string, m string) {
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

	query := r.URL.Query().Get("care-number")

	registerRepo := repository.NewRegisterRepository(sql, w, r)
	err := registerRepo.DeleteRegistrationData(query)
	if err != nil {
		helper.ResponseError(w, val.Id, "failed delete data", "failed delete data", 400, path)
		return
	}

	s, _ := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "deleted"})
	helper.ResponseSuccess(w, val.Id, "delete registration", path, s, 200)
}

func GetRegistrationPatient(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string, m string) {
	// ---- Needed for every request ---
	if !pkg.CheckRequestHeader(w, r, sql, path, m) {
		return
	}

	// Check Header
	val := pkg.CheckUserLogin(w, r, sql, path)
	if val.Status == "authorization" {
		helper.ResponseWarn(w, "", "unauthorization", "unauthorization", 401, path)
		return
	} else if val.Status == "error_format" {
		helper.ResponseWarn(w, "", "unauthorization error format", "unauthorization error format", 400, path)
		return
	}

	query := r.URL.Query()

	date1 := query.Get("date1")
	date2 := query.Get("date2")
	limit := query.Get("limit")
	search := "%" + query.Get("search") + "%"
	lim, _ := strconv.Atoi(limit)

	registerRepo := repository.NewRegisterRepository(sql, w, r)
	data, err := registerRepo.GetRegistrationData(date1, date2, lim, search)
	if err != nil {
		helper.ResponseError(w, val.Id, "failed get data", err.Error(), 401, path)
		return
	}

	s, err := json.Marshal(data)
	if err != nil {
		panic(err.Error())
	}

	helper.ResponseSuccess(w, val.Id, "get patient data", path, s, 200)
}

func GetCurrentRegisterNum(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string, m string) {
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

	query := r.URL.Query()

	date := query.Get("date")
	policlinic := query.Get("policlinic")

	registerRepo := repository.NewRegisterRepository(sql, w, r)
	data, err := registerRepo.GetCurrentRegisterNumber(date, policlinic)
	if err != nil {
		helper.ResponseError(w, val.Id, "failed get data", err.Error(), 401, path)
	}

	s, err := json.Marshal(models.ResponseDataSuccessInt{Status: "success", Response: data})
	if err != nil {
		panic(err.Error())
	}

	helper.ResponseSuccess(w, val.Id, "get current reg number", path, s, 200)
}

func GetCurrentCareNum(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string, m string) {
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

	query := r.URL.Query()

	date := query.Get("date")

	registerRepo := repository.NewRegisterRepository(sql, w, r)
	data, err := registerRepo.GetCurrentCareNumber(date)
	if err != nil {
		helper.ResponseError(w, val.Id, "failed get data", err.Error(), 401, path)
	}

	s, err := json.Marshal(models.ResponseDataSuccessInt{Status: "success", Response: data})
	if err != nil {
		panic(err.Error())
	}

	helper.ResponseSuccess(w, val.Id, "get current care number", path, s, 200)
}
