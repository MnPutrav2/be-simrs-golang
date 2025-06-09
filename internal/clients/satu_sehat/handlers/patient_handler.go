package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/MnPutrav2/be-simrs-golang/internal/clients/satu_sehat/services"
	"github.com/MnPutrav2/be-simrs-golang/internal/helper"
	"github.com/MnPutrav2/be-simrs-golang/internal/models"
	"github.com/MnPutrav2/be-simrs-golang/internal/pkg"
)

func GetSatuSehatPatient(w http.ResponseWriter, r *http.Request, db *sql.DB, path string, m string) {
	if r.Method != m {
		helper.ResponseError(w, "method not allowed", "method not allowed : 400", 400, path)
		return
	}

	token, err := pkg.CreateToken(db)
	if err != nil {
		helper.ResponseError(w, "error create token satu sehat", err.Error()+" : 400", 400, path)
		return
	}

	param := r.URL.Query()

	patientService := services.NewSatuSehatPatient(db, r)
	data, err := patientService.GetDataPatientByNIK(param.Get("nik"), token)
	if err != nil {
		helper.ResponseError(w, "failed fetch data", err.Error()+" : 400", 400, path)
		return
	}

	s, _ := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: data})
	helper.ResponseSuccess(w, "success get patient id (satu-sehat)", "success get patient id (satu-sehat) : 200", s, 200)
}
