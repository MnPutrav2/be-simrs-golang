package handlers

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/MnPutrav2/be-simrs-golang/internal/clients/satu_sehat/models"
	"github.com/MnPutrav2/be-simrs-golang/internal/clients/satu_sehat/services"
	"github.com/MnPutrav2/be-simrs-golang/internal/helper"
	"github.com/MnPutrav2/be-simrs-golang/internal/pkg"
)

func CreateSatuSehatEncounter(w http.ResponseWriter, r *http.Request, db *sql.DB, path string, m string) {
	// ---- Needed for every request ---
	if !pkg.CheckRequestHeader(w, r, db, path, m) {
		return
	}

	// Check Header
	auth := r.Header.Get("Authorization")
	if !pkg.CheckAuthorization(w, path, db, auth) {
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
	if err := db.QueryRow("SELECT users.id FROM users INNER JOIN session_token ON users.id = session_token.users_id WHERE session_token.token = $1", split[1]).Scan(&id); err != nil {
		return
	}

	token, err := pkg.CreateSatuSehatToken(db)
	if err != nil {
		helper.ResponseError(w, id, "error create token satu sehat", err.Error(), 400, path)
		return
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		helper.ResponseWarn(w, id, "empty request body", err.Error(), 400, path)
		return
	}

	var patient models.EncounterResponse
	_ = json.Unmarshal(data, &patient)

	encounterService := services.NewSatuSehatEncounter(db)
	res, err := encounterService.CreateEncounterData(patient, token)
	if err != nil {
		helper.ResponseError(w, id, "error fetch data", err.Error(), 400, path)
		return
	}

	if res.Code == 201 {
		dt, _ := json.Marshal(models.SatuSehatToClientResponse{Status: "success", Response: res.Data})

		helper.ResponseSuccess(w, id, "success create data", path, dt, 201)
	} else {
		dt, _ := json.Marshal(models.SatuSehatToClientResponse{Status: "failed", Response: res.Data})

		helper.ResponseSuccess(w, id, "failed fetch data", path, dt, 400)
	}
}
