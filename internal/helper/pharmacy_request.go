package helper

import (
	"encoding/json"
	"net/http"

	"github.com/MnPutrav2/be-simrs-golang/internal/models"
)

func GetRequestBodyDrugData(w http.ResponseWriter, r *http.Request, path string) (models.DrugData, error) {
	var drug models.DrugData

	// Buat decoder dan disallow field yang tidak dikenal
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	// Decode JSON langsung ke struct
	err := decoder.Decode(&drug)
	if err != nil {
		return models.DrugData{}, err
	}

	return drug, nil
}
