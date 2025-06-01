package main

import (
	"net/http"

	"github.com/MnPutrav2/be-simrs-golang/controllers"
	"github.com/MnPutrav2/be-simrs-golang/internal/config"
	"github.com/MnPutrav2/be-simrs-golang/internal/helper"
)

func main() {
	db := config.SqlDb()

	defer db.Close()

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		controllers.LoginUser(w, r, db, "/login")
	})

	helper.LogWorker("[INFO] server runing in port 8080")
	http.ListenAndServe(":8080", nil)
}
