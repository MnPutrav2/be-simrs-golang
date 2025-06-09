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

func CreateRegistrationPatient(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string, m string) {
	// ---- Needed for every request ---
	if !pkg.CheckRequestHeader(w, r, sql, path, m) {
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
	body, err := helper.GetRequestBodyRegisterData(w, r, path)
	if err != nil {
		helper.ResponseError(w, "empty request body", "empty request body : 400", 400, path)
		return
	}

	registerRepo := repository.NewRegisterRepository(sql, w, r)
	err = registerRepo.CreateRegistrationData(body, path)
	if err != nil {
		return
	}

	s, _ := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "created"})
	helper.ResponseSuccess(w, "create patient registration : 201", path, s, 201)
}

func DeleteRegistrationPatient(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string, m string) {
	// ---- Needed for every request ---
	if !pkg.CheckRequestHeader(w, r, sql, path, m) {
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

	query := r.URL.Query().Get("care-number")

	registerRepo := repository.NewRegisterRepository(sql, w, r)
	err := registerRepo.DeleteRegistrationData(query)
	if err != nil {
		helper.ResponseError(w, "failed delete data", "failed delete data : 400", 400, path)
		return
	}

	s, _ := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "deleted"})
	helper.ResponseSuccess(w, "delete registration : 200", path, s, 200)
}

func GetRegistrationPatient(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string, m string) {
	// ---- Needed for every request ---
	if !pkg.CheckRequestHeader(w, r, sql, path, m) {
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

	query := r.URL.Query()

	date1 := query.Get("date1")
	date2 := query.Get("date2")
	limit := query.Get("limit")
	search := "%" + query.Get("search") + "%"

	registerRepo := repository.NewRegisterRepository(sql, w, r)
	data, err := registerRepo.GetRegistrationData(date1, date2, limit, search)
	if err != nil {
		helper.ResponseError(w, "failed get data", err.Error()+" : 400", 401, path)
	}

	s, err := json.Marshal(data)
	if err != nil {
		panic(err.Error())
	}

	helper.ResponseSuccess(w, "get patient data : 200", path, s, 200)
}

func GetCurrentRegisterNum(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string, m string) {
	// ---- Needed for every request ---
	if !pkg.CheckRequestHeader(w, r, sql, path, m) {
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

	query := r.URL.Query()

	date := query.Get("date")
	policlinic := query.Get("policlinic")

	registerRepo := repository.NewRegisterRepository(sql, w, r)
	data, err := registerRepo.GetCurrentRegisterNumber(date, policlinic)
	if err != nil {
		helper.ResponseError(w, "failed get data", err.Error()+" : 400", 401, path)
	}

	s, err := json.Marshal(models.ResponseDataSuccessInt{Status: "success", Response: data})
	if err != nil {
		panic(err.Error())
	}

	helper.ResponseSuccess(w, "get current reg number : 200", path, s, 200)
}

func GetCurrentCareNum(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string, m string) {
	// ---- Needed for every request ---
	if !pkg.CheckRequestHeader(w, r, sql, path, m) {
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

	query := r.URL.Query()

	date := query.Get("date")

	registerRepo := repository.NewRegisterRepository(sql, w, r)
	data, err := registerRepo.GetCurrentCareNumber(date)
	if err != nil {
		helper.ResponseError(w, "failed get data", err.Error()+" : 400", 401, path)
	}

	s, err := json.Marshal(models.ResponseDataSuccessInt{Status: "success", Response: data})
	if err != nil {
		panic(err.Error())
	}

	helper.ResponseSuccess(w, "get current care number : 200", path, s, 200)
}
