package main

import (
	"fmt"
	"net/http"

	"github.com/MnPutrav2/be-simrs-golang/controllers"
	"github.com/MnPutrav2/be-simrs-golang/internal/config"
	"github.com/MnPutrav2/be-simrs-golang/internal/helper"
)

func main() {
	db := config.SqlDb()

	defer db.Close()

	http.HandleFunc("/user/auth", func(w http.ResponseWriter, r *http.Request) {
		controllers.AuthUser(w, r, db, "/user/auth")
	})

	http.HandleFunc("/user/status", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetUserStatus(w, r, db, "/user/status")
	})

	http.HandleFunc("/user/logout", func(w http.ResponseWriter, r *http.Request) {
		controllers.UserLogout(w, r, db, "/user/logout")
	})

	fmt.Println(helper.LogWorker("[INFO] server runing in port 8080"))
	http.ListenAndServe(":8080", nil)
}
