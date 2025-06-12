package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/MnPutrav2/be-simrs-golang/internal/helper"
	"github.com/MnPutrav2/be-simrs-golang/internal/models"
	"github.com/MnPutrav2/be-simrs-golang/internal/pkg"
)

func AuthUser(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string, m string) {
	if !pkg.CheckRequestHeader(w, r, sql, path, m) {
		return
	}

	// get client request body
	account, err := helper.GetRequestBodyUserAccount(w, r, path)
	if err != nil {
		helper.ResponseWarn(w, 0, "empty request body", "empty request body", 400, path)
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
		helper.ResponseWarn(w, 0, "Login failed : Check your username or password", "username or password error", 400, path)
		return
	}

	// success
	s, err := json.Marshal(models.AuthResponse{Status: "success", Token: pkg.SessionToken(sql, account.Username, account.Password)})
	if err != nil {
		panic(err.Error())
	}

	helper.ResponseSuccess(w, id, "client login", path, s, 201)
}
