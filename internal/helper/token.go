package helper

import (
	"database/sql"
	"encoding/base64"
	"strconv"
	"time"

	"github.com/MnPutrav2/be-simrs-golang/models"
)

func createToken(user string, role string) string {
	t := time.Now().Unix()
	c := strconv.FormatInt(t, 10) + "," + user + "," + role

	return base64.StdEncoding.EncodeToString([]byte(c))
}

func SessionToken(sql *sql.DB, us string, pas string) string {
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

func CheckAvailableToken(db *sql.DB) {
	t := time.Now().Format("2006-01-02 15:04:05")

	q, err := db.Query("SELECT session_token.id, session_token.users_id, session_token.expired FROM session_token")
	if err != nil {
		panic(err.Error())
	}

	for q.Next() {
		var exp models.UserTokenData

		err = q.Scan(&exp.ID, &exp.UserID, &exp.Expired)
		if err != nil {
			panic(err.Error())
		}

		if t == exp.Expired {
			_, err := db.Exec("DELETE FROM session_token WHERE session_token.id = ?", exp.ID)
			if err != nil {
				panic(err.Error())
			}

			m := "session token user_id : " + exp.UserID + " Deleted"

			LogWorker(m)
		}
	}
}
