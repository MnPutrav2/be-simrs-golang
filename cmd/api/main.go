package main

import (
	"fmt"
	"net/http"

	"github.com/MnPutrav2/be-simrs-golang/internal/config"
	"github.com/MnPutrav2/be-simrs-golang/internal/controllers"
	"github.com/MnPutrav2/be-simrs-golang/internal/helper"
)

func main() {
	db := config.SqlDb()

	defer db.Close()

	// User API

	http.HandleFunc("/user/auth", func(w http.ResponseWriter, r *http.Request) {
		controllers.AuthUser(w, r, db, "/user/auth")
	})

	http.HandleFunc("/user/status", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetUserStatus(w, r, db, "/user/status")
	})

	http.HandleFunc("/user/pages", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetUserPages(w, r, db, "/user/pages")
	})

	http.HandleFunc("/user/logout", func(w http.ResponseWriter, r *http.Request) {
		controllers.UserLogout(w, r, db, "/user/logout")
	})

	// User API

	http.HandleFunc("/patient/create", func(w http.ResponseWriter, r *http.Request) {
		controllers.CreatePatient(w, r, db, "/patient/create")
	})

	http.HandleFunc("/patient/get", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetPatient(w, r, db, "/patient/get")
	})

	http.HandleFunc("/patient/delete", func(w http.ResponseWriter, r *http.Request) {
		controllers.DeletePatient(w, r, db, "/patient/delete")
	})

	fmt.Println(helper.LogWorker("[INFO] server runing in port 8080"))
	http.ListenAndServe(":8080", nil)
}
