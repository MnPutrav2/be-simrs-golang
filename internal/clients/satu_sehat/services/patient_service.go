package services

import (
	"database/sql"
	"os"

	// "io/ioutil"
	"net/http"

	"github.com/joho/godotenv"
)

type SatuSehatPatient interface {
	GetDataPatientByNIK(nik string, token string) (string, error)
}

type satuSehatPatient struct {
	sql *sql.DB
	r   *http.Request
}

func NewSatuSehatPatient(sql *sql.DB, r *http.Request) SatuSehatPatient {
	return &satuSehatPatient{sql, r}
}

func (q *satuSehatPatient) GetDataPatientByNIK(nik string, token string) (string, error) {
	_ = godotenv.Load()

	url := os.Getenv("SATU_SEHAT_END_POINT") + "/Patient?identifier=https://fhir.kemkes.go.id/id/nik|" + nik

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return "success", nil
	} else {
		return "failed", err
	}
	// body, _ := ioutil.ReadAll(resp.Body)
}
