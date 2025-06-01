package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/MnPutrav2/be-simrs-golang/internal/helper"
	"github.com/MnPutrav2/be-simrs-golang/models"
)

func UserLogout(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
	if helper.Cors(w, r) {
		return
	}

	if r.Method != "DELETE" {
		helper.ResponseError(w, "method not allowed", "method not allowed : 400", 400, path)
		return
	}

	auth := r.Header.Get("Authorization")
	if !helper.CheckAuthorization(w, path, sql, auth) {
		helper.ResponseError(w, "unauthorization", "unauthorization : 400", 401, path)
		return
	}

	split := strings.SplitN(auth, " ", 2)

	if len(split) != 2 || split[0] != "Bearer" {
		helper.ResponseError(w, "unauthorization error format", "unauthorization error format : 400", 400, path)
		return
	}

	_, err := sql.Exec("DELETE FROM session_token WHERE session_token.token = ?", split[1])
	if err != nil {
		panic(err.Error())
	}

	s, err := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "logout"})
	if err != nil {
		panic(err.Error())
	}

	helper.ResponseSuccess(w, path, s)
}
