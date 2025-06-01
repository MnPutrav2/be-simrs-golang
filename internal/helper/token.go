package helper

import (
	"database/sql"
	"time"

	"github.com/MnPutrav2/be-simrs-golang/models"
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
