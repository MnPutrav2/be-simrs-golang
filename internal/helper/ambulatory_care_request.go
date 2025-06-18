package helper

import (
	"encoding/json"
	"net/http"

	"github.com/MnPutrav2/be-simrs-golang/internal/models"
)

func GetAmbulatoryRequest(w http.ResponseWriter, r *http.Request, path string) (models.RequestAmbulatoryCare, error) {
	// get client request body
	var patient models.RequestAmbulatoryCare

	// Buat decoder dan disallow field yang tidak dikenal
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	// Decode JSON langsung ke struct
	err := decoder.Decode(&patient)
	if err != nil {
		return models.RequestAmbulatoryCare{}, err
	}

	return patient, nil
}

func GetAmbulatoryRequestUpdate(w http.ResponseWriter, r *http.Request, path string) (models.RequestUpdateAmbulatorCare, error) {
	// get client request body
	var patient models.RequestUpdateAmbulatorCare

	// Buat decoder dan disallow field yang tidak dikenal
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	// Decode JSON langsung ke struct
	err := decoder.Decode(&patient)
	if err != nil {
		return models.RequestUpdateAmbulatorCare{}, err
	}

	return patient, nil
}
