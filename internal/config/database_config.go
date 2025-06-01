package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func SqlDb() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	addr := os.Getenv("DB_ADDR")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")

	var dsn string

	if pass == "" {
		dsn = fmt.Sprintf("%s:@tcp(%s)/%s", user, addr, name)
	} else {
		dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s", user, pass, addr, name)
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	return db
}
