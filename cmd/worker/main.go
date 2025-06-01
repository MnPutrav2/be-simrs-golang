package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/MnPutrav2/be-simrs-golang/internal/config"
	"github.com/MnPutrav2/be-simrs-golang/internal/helper"
	"github.com/MnPutrav2/be-simrs-golang/models"
)

type Exp struct {
	ID      int
	UserID  string
	Expired string
}

func main() {
	db := config.SqlDb()
	defer db.Close()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for t := range ticker.C {
		fmt.Println("time now : ", t.Format("2006-01-02 15:04:05"))
		CheckAvailableToken(db)
	}

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

			fmt.Println(helper.LogWorker(m))
		}
	}
}
