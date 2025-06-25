package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/MnPutrav2/be-simrs-golang/internal/helper"
	"github.com/MnPutrav2/be-simrs-golang/internal/pkg"
	"github.com/MnPutrav2/be-simrs-golang/internal/repository"
)

func GetLogs(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string, m string) {
	// ---- Needed for every request ---
	if !pkg.CheckRequestHeader(w, r, sql, path, m) {
		return
	}

	val := pkg.CheckUserLogin(w, r, sql, path)
	if val.Status == "authorization" {
		helper.ResponseWarn(w, "", "unauthorization", "unauthorization", 401, path)
		return
	} else if val.Status == "error_format" {
		helper.ResponseWarn(w, "", "unauthorization error format", "unauthorization error format", 400, path)
		return
	}

	param := r.URL.Query()
	date1 := param.Get("date1")
	date2 := param.Get("date2")

	logRepo := repository.NewLogsRepository(sql)
	dt, err := logRepo.GetLogsData(date1, date2)
	if err != nil {
		panic(err.Error())
	}

	s, _ := json.Marshal(dt)

	w.Write(s)
}
