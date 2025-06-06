package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/MnPutrav2/be-simrs-golang/internal/helper"
	"github.com/MnPutrav2/be-simrs-golang/internal/models"
	"github.com/google/uuid"
)

type authRepository struct {
	sql *sql.DB
}

func NewAuthRepository(sql *sql.DB) AuthRepository {
	return &authRepository{sql}
}

func (q *authRepository) CheckUserToken() error {
	t := time.Now().Format("2006-01-02 15:04:05")

	c, err := q.sql.Query("SELECT session_token.id, session_token.users_id, session_token.expired FROM session_token")
	if err != nil {
		panic(err.Error())
	}

	for c.Next() {
		var exp models.UserTokenData

		err := c.Scan(&exp.ID, &exp.UserID, &exp.Expired)
		if err != nil {
			panic(err.Error())
		}

		if t == exp.Expired {
			_, err := q.sql.Exec("DELETE FROM session_token WHERE session_token.id = ?", exp.ID)
			if err != nil {
				panic(err.Error())
			}

			m := "session token user_id : " + exp.UserID + " Deleted"

			fmt.Println(helper.LogWorker(m))
		}
	}

	return nil
}

func (q *authRepository) CreateSessionToken(us string, pass string) uuid.UUID {
	var user models.User
	tm := time.Now()
	h := tm.Add(6 * time.Hour).Format("2006-01-02 15:04:05")
	err := q.sql.QueryRow("SELECT users.id, users.username, users.role FROM users WHERE users.username = ? AND users.password = ?", us, pass).Scan(&user.ID, &user.Username, &user.Role)
	if err != nil {
		panic(err.Error())
	}

	ut := uuid.New()
	var i int
	err = q.sql.QueryRow("SELECT COUNT(id) FROM session_token WHERE session_token.users_id = ?", user.ID).Scan(&i)
	if err != nil {
		panic(err.Error())
	}

	if i == 0 {
		insert, err := q.sql.Query("INSERT INTO session_token(users_id, token, expired) VALUES(?, ?, ?)", user.ID, ut, h)
		if err != nil {
			panic(err.Error())
		}
		defer insert.Close()

		return ut
	}

	_, err = q.sql.Exec("UPDATE session_token SET token = ?, expired = ? WHERE session_token.users_id = ?", ut, h, user.ID)
	if err != nil {
		panic(err.Error())
	}

	return ut
}

func (q *authRepository) CheckSessionToken(token string) int {
	var u int
	err := q.sql.QueryRow("SELECT COUNT(session_token.id) FROM session_token WHERE session_token.token = ?", token).Scan(&u)
	if err != nil {
		panic(err.Error())
	}

	return u
}
