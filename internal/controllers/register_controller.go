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
		helper.ResponseError(w, 0, "unauthorization", "unauthorization", 401, path)
		return
	}

	split := strings.SplitN(auth, " ", 2)

	if len(split) != 2 || split[0] != "Bearer" {
		helper.ResponseError(w, 0, "unauthorization error format", "unauthorization error format", 400, path)
		return
	}
	// Check Header
	// --- ---
	var id int
	if err := sql.QueryRow("SELECT users.id FROM users INNER JOIN session_token ON users.id = session_token.users_id WHERE session_token.token = ?", split[1]).Scan(&id); err != nil {
		return
	}

	// get client request body
	body, err := helper.GetRequestBodyRegisterData(w, r, path)
	if err != nil {
		helper.ResponseError(w, id, "invalid request body", err.Error(), 400, path)
		return
	}

	registerRepo := repository.NewRegisterRepository(sql, w, r)
	err = registerRepo.CreateRegistrationData(body, path)
	if err != nil {
		return
	}

	s, _ := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "created"})
	helper.ResponseSuccess(w, id, "create patient registration", path, s, 201)
}

func DeleteRegistrationPatient(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string, m string) {
	// ---- Needed for every request ---
	if !pkg.CheckRequestHeader(w, r, sql, path, m) {
		return
	}

	// Check Header
	auth := r.Header.Get("Authorization")
	if !pkg.CheckAuthorization(w, path, sql, auth) {
		helper.ResponseError(w, 0, "unauthorization", "unauthorization", 401, path)
		return
	}

	split := strings.SplitN(auth, " ", 2)

	if len(split) != 2 || split[0] != "Bearer" {
		helper.ResponseError(w, 0, "unauthorization error format", "unauthorization error format", 400, path)
		return
	}
	// Check Header
	// --- ---
	var id int
	if err := sql.QueryRow("SELECT users.id FROM users INNER JOIN session_token ON users.id = session_token.users_id WHERE session_token.token = ?", split[1]).Scan(&id); err != nil {
		return
	}

	query := r.URL.Query().Get("care-number")

	registerRepo := repository.NewRegisterRepository(sql, w, r)
	err := registerRepo.DeleteRegistrationData(query)
	if err != nil {
		helper.ResponseError(w, id, "failed delete data", "failed delete data", 400, path)
		return
	}

	s, _ := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "deleted"})
	helper.ResponseSuccess(w, id, "delete registration", path, s, 200)
}

func GetRegistrationPatient(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string, m string) {
	// ---- Needed for every request ---
	if !pkg.CheckRequestHeader(w, r, sql, path, m) {
		return
	}

	// Check Header
	auth := r.Header.Get("Authorization")
	if !pkg.CheckAuthorization(w, path, sql, auth) {
		helper.ResponseWarn(w, 0, "unauthorization", "unauthorization ", 401, path)
		return
	}

	split := strings.SplitN(auth, " ", 2)

	if len(split) != 2 || split[0] != "Bearer" {
		helper.ResponseWarn(w, 0, "unauthorization error format", "unauthorization error format", 400, path)
		return
	}
	// Check Header
	// --- ---
	var id int
	if err := sql.QueryRow("SELECT users.id FROM users INNER JOIN session_token ON users.id = session_token.users_id WHERE session_token.token = ?", split[1]).Scan(&id); err != nil {
		return
	}

	query := r.URL.Query()

	date1 := query.Get("date1")
	date2 := query.Get("date2")
	limit := query.Get("limit")
	search := "%" + query.Get("search") + "%"

	registerRepo := repository.NewRegisterRepository(sql, w, r)
	data, err := registerRepo.GetRegistrationData(date1, date2, limit, search)
	if err != nil {
		helper.ResponseError(w, id, "failed get data", err.Error(), 401, path)
	}

	s, err := json.Marshal(data)
	if err != nil {
		panic(err.Error())
	}

	helper.ResponseSuccess(w, id, "get patient data", path, s, 200)
}

func GetCurrentRegisterNum(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string, m string) {
	// ---- Needed for every request ---
	if !pkg.CheckRequestHeader(w, r, sql, path, m) {
		return
	}

	// Check Header
	auth := r.Header.Get("Authorization")
	if !pkg.CheckAuthorization(w, path, sql, auth) {
		helper.ResponseWarn(w, 0, "unauthorization", "unauthorization", 401, path)
		return
	}

	split := strings.SplitN(auth, " ", 2)

	if len(split) != 2 || split[0] != "Bearer" {
		helper.ResponseWarn(w, 0, "unauthorization error format", "unauthorization error format", 400, path)
		return
	}
	// Check Header
	// --- ---
	var id int
	if err := sql.QueryRow("SELECT users.id FROM users INNER JOIN session_token ON users.id = session_token.users_id WHERE session_token.token = ?", split[1]).Scan(&id); err != nil {
		return
	}

	query := r.URL.Query()

	date := query.Get("date")
	policlinic := query.Get("policlinic")

	registerRepo := repository.NewRegisterRepository(sql, w, r)
	data, err := registerRepo.GetCurrentRegisterNumber(date, policlinic)
	if err != nil {
		helper.ResponseError(w, id, "failed get data", err.Error(), 401, path)
	}

	s, err := json.Marshal(models.ResponseDataSuccessInt{Status: "success", Response: data})
	if err != nil {
		panic(err.Error())
	}

	helper.ResponseSuccess(w, id, "get current reg number", path, s, 200)
}

func GetCurrentCareNum(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string, m string) {
	// ---- Needed for every request ---
	if !pkg.CheckRequestHeader(w, r, sql, path, m) {
		return
	}

	// Check Header
	auth := r.Header.Get("Authorization")
	if !pkg.CheckAuthorization(w, path, sql, auth) {
		helper.ResponseWarn(w, 0, "unauthorization", "unauthorization", 401, path)
		return
	}

	split := strings.SplitN(auth, " ", 2)

	if len(split) != 2 || split[0] != "Bearer" {
		helper.ResponseWarn(w, 0, "unauthorization error format", "unauthorization error format", 400, path)
		return
	}
	// Check Header
	// --- ---
	var id int
	if err := sql.QueryRow("SELECT users.id FROM users INNER JOIN session_token ON users.id = session_token.users_id WHERE session_token.token = ?", split[1]).Scan(&id); err != nil {
		return
	}

	query := r.URL.Query()

	date := query.Get("date")

	registerRepo := repository.NewRegisterRepository(sql, w, r)
	data, err := registerRepo.GetCurrentCareNumber(date)
	if err != nil {
		helper.ResponseError(w, id, "failed get data", err.Error(), 401, path)
	}

	s, err := json.Marshal(models.ResponseDataSuccessInt{Status: "success", Response: data})
	if err != nil {
		panic(err.Error())
	}

	helper.ResponseSuccess(w, id, "get current care number", path, s, 200)
}
