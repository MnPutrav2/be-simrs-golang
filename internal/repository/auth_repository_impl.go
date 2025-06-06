package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/MnPutrav2/be-simrs-golang/internal/helper"
	"github.com/MnPutrav2/be-simrs-golang/internal/models"
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
