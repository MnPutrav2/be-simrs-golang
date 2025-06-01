package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/MnPutrav2/be-simrs-golang/internal/helper"
	"github.com/MnPutrav2/be-simrs-golang/models"
)

func UserLogout(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
	if !helper.RequestNotAllowed(w, r, "DELETE") {
		fmt.Println(helper.Log("method not allowed : 400", path))
		return
	}

	auth := r.Header.Get("Authorization")
	if !helper.CheckAuthorization(w, path, sql, auth) {
		return
	}

	split := strings.Split(auth, " ")

	_, err := sql.Exec("DELETE FROM session_token WHERE session_token.token = ?", split[1])
	if err != nil {
		panic(err.Error())
	}

	s, err := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "Success logout"})
	if err != nil {
		panic(err.Error())
	}

	helper.ResponseSuccess(w, path, s)
}
