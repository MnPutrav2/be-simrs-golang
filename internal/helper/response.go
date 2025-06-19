package helper

import (
	"encoding/json"
	"net/http"

	"github.com/MnPutrav2/be-simrs-golang/internal/models"
)

func ResponseError(w http.ResponseWriter, id string, m string, log string, c int, path string) {

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
	Log(log, "ERROR", id, path)
}

func ResponseWarn(w http.ResponseWriter, id string, m string, log string, c int, path string) {

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
	Log(log, "WARN", id, path)
}

func ResponseSuccess(w http.ResponseWriter, id string, m string, path string, s []byte, c int) {
	w.WriteHeader(c)
	w.Write(s)
	Log(m, "INFO", id, path)
}
