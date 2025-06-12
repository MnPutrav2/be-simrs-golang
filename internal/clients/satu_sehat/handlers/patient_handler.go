package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/MnPutrav2/be-simrs-golang/internal/clients/satu_sehat/services"
	"github.com/MnPutrav2/be-simrs-golang/internal/helper"
	"github.com/MnPutrav2/be-simrs-golang/internal/models"
	"github.com/MnPutrav2/be-simrs-golang/internal/pkg"
)

func GetSatuSehatPatient(w http.ResponseWriter, r *http.Request, db *sql.DB, path string, m string) {
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

	param := r.URL.Query()

	patientService := services.NewSatuSehatPatient(db, r)
	data, err := patientService.GetDataPatientByNIK(param.Get("nik"), token)
	if err != nil {
		helper.ResponseError(w, 0, "failed fetch data", err.Error(), 400, path)
		return
	}

	s, _ := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: data})
	helper.ResponseSuccess(w, 0, "success get patient id (satu-sehat)", "success get patient id (satu-sehat)", s, 200)
}
