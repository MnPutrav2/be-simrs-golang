package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/MnPutrav2/be-simrs-golang/internal/helper"
	"github.com/MnPutrav2/be-simrs-golang/internal/models"
	"github.com/MnPutrav2/be-simrs-golang/internal/pkg"
	"github.com/MnPutrav2/be-simrs-golang/internal/repository"
)

func GetUserStatus(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string, m string) {
	// ---- Needed for every request ---
	if !pkg.CheckRequestHeader(w, r, sql, path, m) {
		return
	}

	// Check Header
	auth := r.Header.Get("Authorization")
	if !pkg.CheckAuthorization(w, path, sql, auth) {
		helper.ResponseWarn(w, "", "unauthorization", "unauthorization : 400", 401, path)
		return
	}

	split := strings.SplitN(auth, " ", 2)

	if len(split) != 2 || split[0] != "Bearer" {
		helper.ResponseWarn(w, "", "unauthorization error format", "unauthorization error format : 400", 400, path)
		return
	}
	// Check Header
	// --- ---
	var id string
	if err := sql.QueryRow("SELECT users.id FROM users INNER JOIN session_token ON users.id = session_token.users_id WHERE session_token.token = $1", split[1]).Scan(&id); err != nil {
		return
	}

	token := split[1]
	userRepo := repository.NewUserRepository(w, r, sql)
	status, err := userRepo.GetUserStatus(token, path)
	if err != nil {
		return
	}

	s, err := json.Marshal(status)
	if err != nil {
		panic(err.Error())
	}

	helper.ResponseSuccess(w, id, "get user status", path, s, 200)
}

func UserLogout(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string, m string) {
	// ---- Needed for every request ---
	if !pkg.CheckRequestHeader(w, r, sql, path, m) {
		return
	}

	// Check Header
	auth := r.Header.Get("Authorization")
	if !pkg.CheckAuthorization(w, path, sql, auth) {
		helper.ResponseWarn(w, "", "unauthorization", "unauthorization", 401, path)
		return
	}

	split := strings.SplitN(auth, " ", 2)

	if len(split) != 2 || split[0] != "Bearer" {
		helper.ResponseWarn(w, "", "unauthorization error format", "unauthorization error format", 400, path)
		return
	}
	// Check Header
	// --- ---

	var id string
	if err := sql.QueryRow("SELECT users.id FROM users INNER JOIN session_token ON users.id = session_token.users_id WHERE session_token.token = $1", split[1]).Scan(&id); err != nil {
		return
	}

	token := split[1]

	userRepo := repository.NewUserRepository(w, r, sql)
	if err := userRepo.UserLogout(token); err != nil {
		panic(err.Error())
	}

	s, err := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "logout"})
	if err != nil {
		panic(err.Error())
	}

	helper.ResponseSuccess(w, id, "client logout", path, s, 200)
}

func GetUserPages(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string, m string) {
	// ---- Needed for every request ---
	if !pkg.CheckRequestHeader(w, r, sql, path, m) {
		return
	}

	// Check Header
	auth := r.Header.Get("Authorization")
	if !pkg.CheckAuthorization(w, path, sql, auth) {
		helper.ResponseWarn(w, "", "unauthorization", "unauthorization", 401, path)
		return
	}

	split := strings.SplitN(auth, " ", 2)

	if len(split) != 2 || split[0] != "Bearer" {
		helper.ResponseWarn(w, "", "unauthorization error format", "unauthorization error format", 400, path)
		return
	}
	// Check Header
	// --- ---

	var id string
	if err := sql.QueryRow("SELECT users.id FROM users INNER JOIN session_token ON users.id = session_token.users_id WHERE session_token.token = $1", split[1]).Scan(&id); err != nil {
		return
	}

	userRepo := repository.NewUserRepository(w, r, sql)

	token := split[1]
	pageList, _ := userRepo.GetUserPagesData(token, path)

	s, err := json.Marshal(pageList)
	if err != nil {
		panic(err.Error())
	}

	helper.ResponseSuccess(w, id, "get pages", path, s, 200)

}
