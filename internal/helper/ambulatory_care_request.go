package helper

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/MnPutrav2/be-simrs-golang/internal/models"
)

func GetAmbulatoryRequest(w http.ResponseWriter, r *http.Request, path string) (models.RequestAmbulatoryCare, error) {
	// get client request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	// encoding client request body
	var care models.RequestAmbulatoryCare
	err = json.Unmarshal(body, &care)
	if err != nil {
		return models.RequestAmbulatoryCare{}, err
	}

	return care, nil
}

func GetAmbulatoryRequestUpdate(w http.ResponseWriter, r *http.Request, path string) (models.RequestUpdateAmbulatorCare, error) {
	// get client request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	// encoding client request body
	var care models.RequestUpdateAmbulatorCare
	err = json.Unmarshal(body, &care)
	if err != nil {
		return models.RequestUpdateAmbulatorCare{}, err
	}

	return care, nil
}
