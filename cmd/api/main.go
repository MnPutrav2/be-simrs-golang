package main

import (
	"database/sql"
	"net/http"

	"github.com/MnPutrav2/be-simrs-golang/internal/clients/satu_sehat/handlers"
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
	handler("GET", "/patient/get-current-medical-record", controllers.GetCurrentMR, db)

	// Register API
	handler("POST", "/register/create", controllers.CreateRegistrationPatient, db)
	handler("DELETE", "/register/delete", controllers.DeleteRegistrationPatient, db)
	handler("GET", "/register/get", controllers.GetRegistrationPatient, db)
	handler("GET", "/register/get-current-register", controllers.GetCurrentRegisterNum, db)
	handler("GET", "/register/get-current-care-number", controllers.GetCurrentCareNum, db)

	// Ambulatory Care API
	handler("POST", "/ambulatory-care/create", controllers.CreateAmbulatoryCarePatient, db)
	handler("DELETE", "/ambulatory-care/delete", controllers.DeleteAmbulatoryCarePatient, db)
	handler("GET", "/ambulatory-care/get", controllers.GetAmbulatoryCarePatient, db)
	handler("PUT", "/ambulatory-care/update", controllers.UpdateAmbulatoryCarePatient, db)

	// Pharmacy API
	handler("POST", "/pharmacy/create-drug-data", controllers.CreateDrugDatas, db)
	handler("GET", "/pharmacy/get-drug-data", controllers.GetDrugDatas, db)
	handler("PUT", "/pharmacy/update-drug-data", controllers.UpdateDrugDatas, db)
	handler("DELETE", "/pharmacy/delete-drug-data", controllers.DeleteDrugDatas, db)
	handler("GET", "/pharmacy/get-distributor", controllers.GetDistributor, db)
	handler("POST", "/pharmacy/create-recipe", controllers.CreateRecipe, db)
	handler("POST", "/pharmacy/create-recipe-compound", controllers.CreateRecipeCompound, db)

	// Satu Sehat
	handler("GET", "/satu-sehat/get-patient-by-nik", handlers.GetSatuSehatPatient, db)
	handler("POST", "/satu-sehat/create-encounter", handlers.CreateSatuSehatEncounter, db)
	handler("POST", "/satu-sehat/create-condition", handlers.CreateSatuSehatCondition, db)
	handler("POST", "/satu-sehat/create-observation", handlers.CreateSatuSehatObservation, db)
	handler("POST", "/satu-sehat/create-clinical-impression", handlers.CreateSatuSehatClinicalImpression, db)
	handler("POST", "/satu-sehat/create-care-plan", handlers.CreateSatuSehatCarePlan, db)

	// Logs
	handler("GET", "/logs", controllers.GetLogs, db)

	helper.LogWorker("[INFO] server runing in port 8080")
	http.ListenAndServe(":8080", nil)
}

func handler(m string, path string, handler func(w http.ResponseWriter, r *http.Request, db *sql.DB, path string, m string), db *sql.DB) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, db, path, m)
	})
}
