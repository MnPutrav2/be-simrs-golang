package helper

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/MnPutrav2/be-simrs-golang/internal/models"
)

func GetRequestBodyUserAccount(w http.ResponseWriter, r *http.Request, path string) (models.UserAccount, error) {
	// get client request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	// encoding client request body
	var account models.UserAccount
	err = json.Unmarshal(body, &account)
	if err != nil {
		return models.UserAccount{}, err
	}

	return account, nil
}
