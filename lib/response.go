package lib

import (
	"encoding/json"
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
	Log(log, path)
}
