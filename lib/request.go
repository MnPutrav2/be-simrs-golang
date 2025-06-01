package lib

import (
	"encoding/json"
	"net/http"

	"github.com/MnPutrav2/be-simrs-golang/models"
)

func RequestNotAllowed(w http.ResponseWriter, r *http.Request, m string) bool {
	if r.Method != m {
		s, err := json.Marshal(models.ResponseDataError{Status: "failed", Errors: "method not allowed"})
		if err != nil {
			panic(err)
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Write(s)
		return false
	}

	return true
}
