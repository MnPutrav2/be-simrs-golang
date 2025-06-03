package pkg

import (
	"database/sql"
	"net/http"
	"strings"
	"time"

	"github.com/MnPutrav2/be-simrs-golang/internal/models"
	"github.com/google/uuid"
)

func createToken() uuid.UUID {
	uid := uuid.New()

	return uid
}

func SessionToken(sql *sql.DB, us string, pas string) uuid.UUID {
	var user models.User
	tm := time.Now()
	h := tm.Add(6 * time.Hour).Format("2006-01-02 15:04:05")
	err := sql.QueryRow("SELECT users.id, users.username, users.role FROM users WHERE users.username = ? AND users.password = ?", us, pas).Scan(&user.ID, &user.Username, &user.Role)
	if err != nil {
		panic(err.Error())
	}

	ut := createToken()
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

func CheckAuthorization(w http.ResponseWriter, path string, db *sql.DB, auth string) bool {
	split := strings.SplitN(auth, " ", 2)

	if len(split) != 2 || split[0] != "Bearer" {
		return false
	}

	if split[0] != "Bearer" {
		return false
	}

	var q int
	err := db.QueryRow("SELECT COUNT(session_token.id) FROM session_token WHERE session_token.token = ?", split[1]).Scan(&q)
	if err != nil {
		panic(err.Error())
	}

	if q == 0 {
		return false
	}

	return true
}
