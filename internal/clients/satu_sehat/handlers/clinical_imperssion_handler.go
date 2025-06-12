package handlers

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"

	"github.com/MnPutrav2/be-simrs-golang/internal/clients/satu_sehat/models"
	"github.com/MnPutrav2/be-simrs-golang/internal/clients/satu_sehat/services"
	"github.com/MnPutrav2/be-simrs-golang/internal/helper"
	"github.com/MnPutrav2/be-simrs-golang/internal/pkg"
)

func CreateSatuSehatClinicalImpression(w http.ResponseWriter, r *http.Request, db *sql.DB, path string, m string) {
	if r.Method != m {
		helper.ResponseError(w, "method not allowed", "method not allowed : 400", 400, path)
		return
	}

	token, err := pkg.CreateSatuSehatToken(db)
	if err != nil {
		helper.ResponseError(w, "error create token satu sehat", err.Error()+" : 400", 400, path)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		helper.ResponseError(w, "empty request body", err.Error()+" : 400", 400, path)
		return
	}

	var patient models.ClinicalImpressionClientRequest
	_ = json.Unmarshal(body, &patient)

	clinicalImpressionService := services.NewSatuSehatClinicalImpression(db)
	res, err := clinicalImpressionService.CreateClinicalImpression(patient, token)
	if err != nil {
		helper.ResponseError(w, "error fetch data", err.Error()+" : 400", 400, path)
		return
	}

	if res.Code == 201 {
		dt, _ := json.Marshal(models.SatuSehatToClientResponse{Status: "success", Response: res.Data})

		helper.ResponseSuccess(w, "success create data", path, dt, 201)
	} else {
		dt, _ := json.Marshal(models.SatuSehatToClientResponse{Status: "failed", Response: res.Data})

		helper.ResponseSuccess(w, "failed fetch data : 400", path, dt, 400)
	}
}
