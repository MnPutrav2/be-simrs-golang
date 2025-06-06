package helper

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/MnPutrav2/be-simrs-golang/internal/models"
)

func GetRequestBodyPatientData(w http.ResponseWriter, r *http.Request, path string) (models.PatientData, error) {
	// get client request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	// encoding client request body
	var patient models.PatientData
	err = json.Unmarshal(body, &patient)
	if err != nil {
		return models.PatientData{}, err
	}

	return patient, nil
}
