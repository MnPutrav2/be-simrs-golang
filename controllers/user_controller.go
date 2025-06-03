package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/MnPutrav2/be-simrs-golang/internal/helper"
	"github.com/MnPutrav2/be-simrs-golang/models"
)

func GetUserStatus(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
	// ---- Needed for every request ---
	// Check Method
	if helper.Cors(w, r) {
		return
	}

	if r.Method != "GET" {
		helper.ResponseError(w, "method not allowed", "method not allowed : 400", 400, path)
		return
	}
	// Check Method

	// Check Header
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
	// Check Header
	// --- ---

	var id int

	err := sql.QueryRow("SELECT session_token.users_id FROM session_token WHERE session_token.token = ?", split[1]).Scan(&id)
	if err != nil {
		helper.ResponseError(w, "unauthorization", "unauthorization : 400", 401, path)
		return
	}

	var user models.EmployeeData

	err = sql.QueryRow("SELECT employee.id, employee.name, employee.gender FROM employee INNER JOIN users ON employee.id = users.employee_id WHERE users.id = ?", id).Scan(&user.Employee_ID, &user.Name, &user.Gender)
	if err != nil {
		helper.ResponseError(w, "employee data not found", "employee data not found : 404", 404, path)
		return
	}

	s, err := json.Marshal(models.EmployeeData{Employee_ID: user.Employee_ID, Name: user.Name, Gender: user.Gender})
	if err != nil {
		panic(err.Error())
	}

	helper.ResponseSuccess(w, "get user status : 200", path, s)
}

func UserLogout(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
	// ---- Needed for every request ---
	// Check Method
	if helper.Cors(w, r) {
		return
	}

	if r.Method != "DELETE" {
		helper.ResponseError(w, "method not allowed", "method not allowed : 400", 400, path)
		return
	}
	// Check Method

	// Check Header
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
	// Check Header
	// --- ---

	_, err := sql.Exec("DELETE FROM session_token WHERE session_token.token = ?", split[1])
	if err != nil {
		panic(err.Error())
	}

	s, err := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "logout"})
	if err != nil {
		panic(err.Error())
	}

	helper.ResponseSuccess(w, "client logout : 200", path, s)
}
