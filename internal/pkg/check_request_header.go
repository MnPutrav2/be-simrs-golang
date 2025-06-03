package pkg

import (
	"database/sql"
	"net/http"

	"github.com/MnPutrav2/be-simrs-golang/internal/helper"
)

func CheckRequestHeader(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string, m string) {
	// ---- Needed for every request ---
	// Check Method
	if Cors(w, r) {
		return
	}

	if r.Method != m {
		helper.ResponseError(w, "method not allowed", "method not allowed : 400", 400, path)
		return
	}
	// Check Method
}
