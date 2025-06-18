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

func CreateDrugDatas(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string, m string) {
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
