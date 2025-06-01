package controllers

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/MnPutrav2/be-simrs-golang/lib"
	"github.com/MnPutrav2/be-simrs-golang/models"
)

type token struct {
	Status string `json:"status"`
	Token  string `json:"token"`
}

func LoginUser(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
	// check client request method
	if !lib.RequestNotAllowed(w, r, "POST") {
		lib.Log("method not allowed : 400", path)
		return
	}

	// get client request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		lib.ResponseError(w, "Error encoding JSON", "failed encoding data : 500", 500, path)
		return
	}

	// encoding client request body
	var account models.UserAccount
	err = json.Unmarshal(body, &account)
	if err != nil {
		lib.ResponseError(w, "No JSON data", "empty client request body : 400", 400, path)
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
		lib.ResponseError(w, "Login failed : Check your username or password", "failed login : 400", 400, path)
		return
	}

	// success
	s, err := json.Marshal(token{Status: "success", Token: sessionToken(sql, account.Username, account.Password)})
	if err != nil {
		panic(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	w.Write(s)
	lib.Log("success login : 200", path)
}

func createToken(user string, role string) string {
	t := time.Now().Unix()
	c := strconv.FormatInt(t, 10) + "," + user + "," + role

	return base64.StdEncoding.EncodeToString([]byte(c))
}

func sessionToken(sql *sql.DB, us string, pas string) string {
	var user models.User
	tm := time.Now()
	h := tm.Add(6 * time.Hour).Format("2006-01-02 15:04:05")
	err := sql.QueryRow("SELECT users.id, users.username, users.role FROM users WHERE users.username = ? AND users.password = ?", us, pas).Scan(&user.ID, &user.Username, &user.Role)
	if err != nil {
		panic(err.Error())
	}

	ut := createToken(user.Username, user.Role)
	var i int
	err = sql.QueryRow("SELECT COUNT(id) FROM session_token WHERE session_token.users_id = ?", user.ID).Scan(&i)
	if err != nil {
		panic(err.Error())
	}

	if i == 0 {
		insert, err := sql.Query("INSERT INTO session_token(users_id, token, expired) VALUES(?, ?, ?)", user.ID, ut, h)
		if err != nil {
			panic(err.Error())
		}
		defer insert.Close()

		return ut
	}

	_, err = sql.Exec("UPDATE session_token SET token = ?, expired = ? WHERE session_token.users_id = ?", ut, h, user.ID)
	if err != nil {
		panic(err.Error())
	}

	return ut
}
