package handlers

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"

	"github.com/MnPutrav2/be-simrs-golang/internal/clients/satu_sehat/services"
	"github.com/MnPutrav2/be-simrs-golang/internal/helper"
	"github.com/MnPutrav2/be-simrs-golang/internal/models"
	"github.com/MnPutrav2/be-simrs-golang/internal/pkg"
)

func CreateSatuSehatCondition(w http.ResponseWriter, r *http.Request, db *sql.DB, path string, m string) {
	if r.Method != m {
		helper.ResponseError(w, "method not allowed", "method not allowed : 400", 400, path)
		return
	}

	token, err := pkg.CreateToken(db)
	if err != nil {
		helper.ResponseError(w, "error create token satu sehat", err.Error()+" : 400", 400, path)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	var patient models.ConditionClientRequest
	_ = json.Unmarshal(body, &patient)

	conditionService := services.NewSatuSehatCondition(db)
	res, err := conditionService.CreateSatuSehatCondition(patient, token)
	if err != nil {
		helper.ResponseError(w, "error fetch data", err.Error()+" : 400", 400, path)
		return
	}

	helper.ResponseSuccess(w, "success fetch data", path, res, 200)
}
