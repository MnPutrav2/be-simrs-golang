package helper

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MnPutrav2/be-simrs-golang/models"
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
