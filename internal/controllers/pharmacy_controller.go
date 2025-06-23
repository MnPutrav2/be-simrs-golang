package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/MnPutrav2/be-simrs-golang/internal/helper"
	"github.com/MnPutrav2/be-simrs-golang/internal/models"
	"github.com/MnPutrav2/be-simrs-golang/internal/pkg"
	"github.com/MnPutrav2/be-simrs-golang/internal/repository"
)

func CreateDrugDatas(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string, m string) {
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

	patient, err := helper.GetRequestBodyDrugData(w, r, path)
	if err != nil {
		helper.ResponseWarn(w, id, "invalid request body", err.Error(), 400, path)
		return
	}

	pharmacyRepo := repository.NewPharmacyRepository(sql)
	if err := pharmacyRepo.CreateDrugData(patient); err != nil {
		helper.ResponseWarn(w, id, "failed create drug data", err.Error(), 404, path)
		return
	}

	s, err := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "created"})
	if err != nil {
		helper.ResponseError(w, id, "error server", err.Error(), 500, path)
		return
	}

	helper.ResponseSuccess(w, id, "create drug data", path, s, 201)
}

func GetDrugDatas(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string, m string) {
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

	var param = r.URL.Query()
	var search = "%" + param.Get("search") + "%"
	limit, _ := strconv.Atoi(param.Get("limit"))

	pharmacyRepo := repository.NewPharmacyRepository(sql)
	data, err := pharmacyRepo.GetDrugData(search, limit)
	if err != nil {
		helper.ResponseWarn(w, id, "failed get drug data", err.Error(), 404, path)
		return
	}

	s, err := json.Marshal(data)
	if err != nil {
		panic(err.Error())
	}

	helper.ResponseSuccess(w, id, "get drug data", path, s, 200)
}

func UpdateDrugDatas(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string, m string) {
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

	drug, err := helper.GetRequestBodyDrugDataUpdate(w, r, path)
	if err != nil {
		helper.ResponseWarn(w, id, "invalid request body", err.Error(), 400, path)
		return
	}

	pharmacyRepo := repository.NewPharmacyRepository(sql)
	if err := pharmacyRepo.UpdateDrugData(drug); err != nil {
		helper.ResponseWarn(w, id, "failed create drug data", err.Error(), 404, path)
		return
	}

	s, err := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "updated"})
	if err != nil {
		helper.ResponseError(w, id, "error server", err.Error(), 500, path)
		return
	}

	helper.ResponseSuccess(w, id, "update drug data", path, s, 200)
}

func DeleteDrugDatas(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string, m string) {
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

	var param = r.URL.Query()
	var drug = param.Get("drug_id")

	pharmacyRepo := repository.NewPharmacyRepository(sql)
	if err := pharmacyRepo.DeleteDrugData(drug); err != nil {
		helper.ResponseWarn(w, id, "failed create drug data", err.Error(), 404, path)
		return
	}

	s, err := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "deleted"})
	if err != nil {
		helper.ResponseError(w, id, "error server", err.Error(), 500, path)
		return
	}

	helper.ResponseSuccess(w, id, "delete drug data", path, s, 200)
}

func GetDistributor(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string, m string) {
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

	pharmacyRepo := repository.NewPharmacyRepository(sql)
	data, err := pharmacyRepo.GetDistributor()
	if err != nil {
		helper.ResponseWarn(w, id, "failed distributor data", err.Error(), 404, path)
		return
	}

	s, err := json.Marshal(data)
	if err != nil {
		panic(err.Error())
	}

	helper.ResponseSuccess(w, id, "get distributor data", path, s, 200)
}

func CreateRecipe(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string, m string) {
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

	patient, err := helper.GetRequestBodyDrugRecipe(w, r, path)
	if err != nil {
		helper.ResponseWarn(w, id, "invalid request body", err.Error(), 400, path)
		return
	}

	pharmacyRepo := repository.NewPharmacyRepository(sql)
	if err := pharmacyRepo.CreateRecipe(patient); err != nil {
		helper.ResponseWarn(w, id, "failed create recipe", err.Error(), 404, path)
		return
	}

	s, err := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "created"})
	if err != nil {
		helper.ResponseError(w, id, "error server", err.Error(), 500, path)
		return
	}

	helper.ResponseSuccess(w, id, "create recipe success", path, s, 201)
}
