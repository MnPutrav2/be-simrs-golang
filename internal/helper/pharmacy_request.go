package helper

import (
	"encoding/json"
	"net/http"

	"github.com/MnPutrav2/be-simrs-golang/internal/models"
)

func GetRequestBodyDrugData(w http.ResponseWriter, r *http.Request, path string) (models.RequestBodyDrugData, error) {
	var drug models.RequestBodyDrugData

	// Buat decoder dan disallow field yang tidak dikenal
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	// Decode JSON langsung ke struct
	err := decoder.Decode(&drug)
	if err != nil {
		return models.RequestBodyDrugData{}, err
	}

	return drug, nil
}

func GetRequestBodyDrugDataUpdate(w http.ResponseWriter, r *http.Request, path string) (models.RequestBodyDrugDataUpdate, error) {
	var drug models.RequestBodyDrugDataUpdate

	// Buat decoder dan disallow field yang tidak dikenal
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	// Decode JSON langsung ke struct
	err := decoder.Decode(&drug)
	if err != nil {
		return models.RequestBodyDrugDataUpdate{}, err
	}

	return drug, nil
}
