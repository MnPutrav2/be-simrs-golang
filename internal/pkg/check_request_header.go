package pkg

import (
	"database/sql"
	"net/http"

	"github.com/MnPutrav2/be-simrs-golang/internal/helper"
)

func CheckRequestHeader(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string, m string) bool {
	// ---- Needed for every request ---
	// Check Method
	if Cors(w, r) {
		return false
	}

	if r.Method != m {
		helper.ResponseError(w, "", "method not allowed", "method not allowed : 400", 400, path)
		return false
	}
	// Check Method
	return true
}
