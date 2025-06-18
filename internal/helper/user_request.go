package helper

import (
	"encoding/json"
	"net/http"

	"github.com/MnPutrav2/be-simrs-golang/internal/models"
)

func GetRequestBodyUserAccount(w http.ResponseWriter, r *http.Request, path string) (models.UserAccount, error) {
	// get client request body
	var user models.UserAccount

	// Buat decoder dan disallow field yang tidak dikenal
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	// Decode JSON langsung ke struct
	err := decoder.Decode(&user)
	if err != nil {
		return models.UserAccount{}, err
	}

	return user, nil
}
