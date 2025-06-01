package controllers

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"

	"github.com/MnPutrav2/be-simrs-golang/internal/helper"
	"github.com/MnPutrav2/be-simrs-golang/models"
)

type token struct {
	Status string `json:"status"`
	Token  string `json:"token"`
}

func LoginUser(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
	// check client request method
	if !helper.RequestNotAllowed(w, r, "POST") {
		helper.Log("method not allowed : 400", path)
		return
	}

	// get client request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		helper.ResponseError(w, "Error encoding JSON", "failed encoding data : 500", 500, path)
		return
	}

	// encoding client request body
	var account models.UserAccount
	err = json.Unmarshal(body, &account)
	if err != nil {
		helper.ResponseError(w, "No JSON data", "empty client request body : 400", 400, path)
		return
	}

	// check user
	// check user available
	var id int
	err = sql.QueryRow("SELECT COUNT(id) FROM users WHERE users.username = ? AND users.password = ?", account.Username, account.Password).Scan(&id)
	if err != nil {
		panic(err.Error())
	}

	// if account not available
	if id != 1 {
		helper.ResponseError(w, "Login failed : Check your username or password", "failed login : 400", 400, path)
		return
	}

	// success
	s, err := json.Marshal(token{Status: "success", Token: helper.SessionToken(sql, account.Username, account.Password)})
	if err != nil {
		panic(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	w.Write(s)
	helper.Log("success login : 200", path)
}
