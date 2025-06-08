package helper

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/MnPutrav2/be-simrs-golang/internal/models"
)

func GetAmbulatoryRequest(w http.ResponseWriter, r *http.Request, path string) (models.AmbulatoryCare, error) {
	// get client request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	// encoding client request body
	var care models.AmbulatoryCare
	err = json.Unmarshal(body, &care)
	if err != nil {
		return models.AmbulatoryCare{}, err
	}

	return care, nil
}
