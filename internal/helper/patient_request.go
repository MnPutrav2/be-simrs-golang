package helper

import (
	"encoding/json"
	"net/http"

	"github.com/MnPutrav2/be-simrs-golang/internal/models"
)

func GetRequestBodyPatientData(w http.ResponseWriter, r *http.Request, path string) (models.PatientData, error) {
	// get client request body
	var patient models.PatientData

	// Buat decoder dan disallow field yang tidak dikenal
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	// Decode JSON langsung ke struct
	err := decoder.Decode(&patient)
	if err != nil {
		return models.PatientData{}, err
	}

	return patient, nil
}

func GetRequestBodyPatientDataUpdate(w http.ResponseWriter, r *http.Request, path string) (models.PatientDataUpdate, error) {
	// get client request body
	var patient models.PatientDataUpdate

	// Buat decoder dan disallow field yang tidak dikenal
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	// Decode JSON langsung ke struct
	err := decoder.Decode(&patient)
	if err != nil {
		return models.PatientDataUpdate{}, err
	}

	return patient, nil
}
