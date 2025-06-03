package helper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/MnPutrav2/be-simrs-golang/internal/models"
)

func ResponseError(w http.ResponseWriter, m string, log string, c int, path string) {

	res := models.ResponseDataError{
		Status: "failed",
		Errors: m,
	}

	data, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(c)
	w.Write(data)
	fmt.Println(Log(log, path))
}

func ResponseSuccess(w http.ResponseWriter, m string, path string, s []byte) {
	w.WriteHeader(http.StatusOK)
	w.Write(s)
	fmt.Println(Log(m, path))
}
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
