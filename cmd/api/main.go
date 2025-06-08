package main

import (
	"database/sql"
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
	handler("POST", "/user/auth", controllers.AuthUser, db)
	handler("GET", "/user/status", controllers.GetUserStatus, db)
	handler("GET", "/user/pages", controllers.GetUserPages, db)
	handler("DELETE", "/user/logout", controllers.UserLogout, db)

	// Patient API
	handler("POST", "/patient/create", controllers.CreatePatient, db)
	handler("GET", "/patient/get", controllers.GetPatient, db)
	handler("PUT", "/patient/update", controllers.UpdatePatientData, db)
	handler("DELETE", "/patient/delete", controllers.DeletePatient, db)
	handler("GET", "/patient/getCurrentMedicalRecord", controllers.GetCurrentMR, db)

	// Register API
	handler("POST", "/register/create", controllers.CreateRegistrationPatient, db)
	handler("DELETE", "/register/delete", controllers.DeleteRegistrationPatient, db)
	handler("GET", "/register/get", controllers.GetRegistrationPatient, db)

	// Ambulatory Care API
	handler("POST", "/ambulatory-care/create", controllers.CreateAmbulatoryCarePatient, db)
	handler("DELETE", "/ambulatory-care/delete", controllers.DeleteAmbulatoryCarePatient, db)

	fmt.Println(helper.LogWorker("[INFO] server runing in port 8080"))
	http.ListenAndServe(":8080", nil)
}

func handler(m string, path string, handler func(w http.ResponseWriter, r *http.Request, db *sql.DB, path string, m string), db *sql.DB) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, db, path, m)
	})
}
