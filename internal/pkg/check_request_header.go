package pkg

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/MnPutrav2/be-simrs-golang/internal/helper"
)

type TokenId struct {
	Token  string
	Id     string
	Status string
}

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

func CheckUserLogin(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) TokenId {
	// Check Header
	auth := r.Header.Get("Authorization")
	if !CheckAuthorization(w, path, sql, auth) {
		return TokenId{Token: "", Id: "", Status: "authorization"}
	}

	split := strings.SplitN(auth, " ", 2)

	if len(split) != 2 || split[0] != "Bearer" {
		return TokenId{Token: "", Id: "", Status: "error_format"}
	}
	// Check Header
	// --- ---

	var id string
	if err := sql.QueryRow("SELECT users.id FROM users INNER JOIN session_token ON users.id = session_token.users_id WHERE session_token.token = $1", split[1]).Scan(&id); err != nil {
		return TokenId{Token: "", Id: "", Status: "authorization"}
	}

	return TokenId{Token: split[1], Id: id, Status: "success"}

}
