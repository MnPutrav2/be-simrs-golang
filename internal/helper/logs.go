package helper

import (
	"fmt"
	"time"

	"github.com/MnPutrav2/be-simrs-golang/internal/config"
)

func Log(m string, lev string, id string, path string) {

	db := config.SqlDb()
	defer db.Close()

	t := time.Now()
	var ty string

	if id == "" {
		ty = "[System]"
	} else {
		ty = "[User : " + id + "]"
	}

	if lev == "WARN" || lev == "ERROR" {
		if _, err := db.Exec("INSERT INTO logs(users_id, level, message, path) VALUES($1, $2, $3, $4)", id, lev, m, path); err != nil {
			panic(err.Error())
		}
	}

	log := fmt.Sprintf("[%s] %s %s %s", t.Format("02 January 2006, 15:04:05"), ty, m, path)

	fmt.Println(log)
}

func LogWorker(m string) {
	t := time.Now()
	log := fmt.Sprintf("[%s] %s", t.Format("02 January 2006, 15:04:05"), m)

	fmt.Println(log)
}
