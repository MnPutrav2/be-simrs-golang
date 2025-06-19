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

func CreateAmbulatoryCarePatient(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string, m string) {
	// ---- Needed for every request ---
	if !pkg.CheckRequestHeader(w, r, sql, path, m) {
		return
	}

	// Check Header
	auth := r.Header.Get("Authorization")
	if !pkg.CheckAuthorization(w, path, sql, auth) {
		helper.ResponseError(w, "", "unauthorization", "unauthorization", 401, path)
		return
	}

	split := strings.SplitN(auth, " ", 2)

	if len(split) != 2 || split[0] != "Bearer" {
		helper.ResponseError(w, "", "unauthorization error format", "unauthorization error format", 400, path)
		return
	}
	// Check Header
	// --- ---

	var id string
	if err := sql.QueryRow("SELECT users.id FROM users INNER JOIN session_token ON users.id = session_token.users_id WHERE session_token.token = $1", split[1]).Scan(&id); err != nil {
		return
	}

	care, err := helper.GetAmbulatoryRequest(w, r, path)
	if err != nil {
		helper.ResponseError(w, id, "invalid json format", err.Error(), 400, path)
		return
	}

	ambulatoryRepo := repository.NewAmbulatoryCareRepository(sql, w, r)
	err = ambulatoryRepo.CreateAmbulatoryCareData(care)
	if err != nil {
		helper.ResponseError(w, id, "failed create data", err.Error(), 400, path)
		return
	}

	s, _ := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "created"})
	helper.ResponseSuccess(w, id, "create ambulatory care", path, s, 201)
}

func DeleteAmbulatoryCarePatient(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string, m string) {
	// ---- Needed for every request ---
	if !pkg.CheckRequestHeader(w, r, sql, path, m) {
		return
	}

	// Check Header
	auth := r.Header.Get("Authorization")
	if !pkg.CheckAuthorization(w, path, sql, auth) {
		helper.ResponseError(w, "", "unauthorization", "unauthorization", 401, path)
		return
	}

	split := strings.SplitN(auth, " ", 2)

	if len(split) != 2 || split[0] != "Bearer" {
		helper.ResponseError(w, "", "unauthorization error format", "unauthorization error format", 400, path)
		return
	}
	// Check Header
	// --- ---

	var id string
	if err := sql.QueryRow("SELECT users.id FROM users INNER JOIN session_token ON users.id = session_token.users_id WHERE session_token.token = $1", split[1]).Scan(&id); err != nil {
		panic(err.Error)
	}

	query := r.URL.Query()
	careNum := query.Get("care-number")
	date := query.Get("date")

	ambulatoryRepo := repository.NewAmbulatoryCareRepository(sql, w, r)
	err := ambulatoryRepo.DeleteAmbulatoryCareData(careNum, date)
	if err != nil {
		helper.ResponseError(w, id, "failed delete data", err.Error(), 400, path)
		return
	}

	s, _ := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "deleted"})
	helper.ResponseSuccess(w, id, "delete ambulatory care", path, s, 200)
}

func GetAmbulatoryCarePatient(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string, m string) {
	// ---- Needed for every request ---
	if !pkg.CheckRequestHeader(w, r, sql, path, m) {
		return
	}

	// Check Header
	auth := r.Header.Get("Authorization")
	if !pkg.CheckAuthorization(w, path, sql, auth) {
		helper.ResponseError(w, "", "unauthorization", "unauthorization", 401, path)
		return
	}

	split := strings.SplitN(auth, " ", 2)

	if len(split) != 2 || split[0] != "Bearer" {
		helper.ResponseError(w, "", "unauthorization error format", "unauthorization error format", 400, path)
		return
	}
	// Check Header
	// --- ---

	var id string
	if err := sql.QueryRow("SELECT users.id FROM users INNER JOIN session_token ON users.id = session_token.users_id WHERE session_token.token = $1", split[1]).Scan(&id); err != nil {
		panic(err.Error)
	}

	query := r.URL.Query()
	careNum := query.Get("care-number")
	date1 := query.Get("date1")
	date2 := query.Get("date2")

	ambulatoryRepo := repository.NewAmbulatoryCareRepository(sql, w, r)
	res, err := ambulatoryRepo.GetAmbulatoryCareData(careNum, date1, date2)
	if err != nil {
		helper.ResponseError(w, id, "failed get data", err.Error(), 400, path)
		return
	}

	s, _ := json.Marshal(res)
	helper.ResponseSuccess(w, id, "get ambulatory care", path, s, 200)
}

func UpdateAmbulatoryCarePatient(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string, m string) {
	// ---- Needed for every request ---
	if !pkg.CheckRequestHeader(w, r, sql, path, m) {
		return
	}

	// Check Header
	auth := r.Header.Get("Authorization")
	if !pkg.CheckAuthorization(w, path, sql, auth) {
		helper.ResponseError(w, "", "unauthorization", "unauthorization", 401, path)
		return
	}

	split := strings.SplitN(auth, " ", 2)

	if len(split) != 2 || split[0] != "Bearer" {
		helper.ResponseError(w, "", "unauthorization error format", "unauthorization error format", 400, path)
		return
	}
	// Check Header
	// --- ---

	var id string
	if err := sql.QueryRow("SELECT users.id FROM users INNER JOIN session_token ON users.id = session_token.users_id WHERE session_token.token = $1", split[1]).Scan(&id); err != nil {
		panic(err.Error)
	}

	care, err := helper.GetAmbulatoryRequestUpdate(w, r, path)
	if err != nil {
		helper.ResponseError(w, id, "invalid json format", err.Error(), 400, path)
		return
	}

	ambulatoryRepo := repository.NewAmbulatoryCareRepository(sql, w, r)
	err = ambulatoryRepo.UpdateAmbulatoryCareData(care)
	if err != nil {
		helper.ResponseError(w, id, "failed create data", err.Error(), 400, path)
		return
	}

	s, _ := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "update"})
	helper.ResponseSuccess(w, id, "update ambulatory care", path, s, 200)
}
