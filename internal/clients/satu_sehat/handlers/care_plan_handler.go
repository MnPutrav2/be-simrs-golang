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

func CreateSatuSehatCarePlan(w http.ResponseWriter, r *http.Request, db *sql.DB, path string, m string) {
	// ---- Needed for every request ---
	if !pkg.CheckRequestHeader(w, r, db, path, m) {
		return
	}

	// Check Header
	auth := r.Header.Get("Authorization")
	if !pkg.CheckAuthorization(w, path, db, auth) {
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

	token, err := pkg.CreateSatuSehatToken(db)
	if err != nil {
		helper.ResponseError(w, 0, "error create token satu sehat", err.Error(), 400, path)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		helper.ResponseWarn(w, 0, "empty request body", err.Error(), 400, path)
		return
	}

	var patient models.CarePlantRequest
	_ = json.Unmarshal(body, &patient)

	carePlanService := services.NewSatuSehatCarePlan(db)
	res, err := carePlanService.CreateCarePlan(patient, token)
	if err != nil {
		helper.ResponseError(w, 0, "error fetch data", err.Error(), 400, path)
		return
	}

	if res.Code == 201 {
		dt, _ := json.Marshal(models.SatuSehatToClientResponse{Status: "success", Response: res.Data})

		helper.ResponseSuccess(w, 0, "success create data", path, dt, 201)
	} else {
		dt, _ := json.Marshal(models.SatuSehatToClientResponse{Status: "failed", Response: res.Data})

		helper.ResponseSuccess(w, 0, "failed fetch data", path, dt, 400)
	}
}
