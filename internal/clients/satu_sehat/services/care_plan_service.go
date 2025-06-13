package services

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/MnPutrav2/be-simrs-golang/internal/clients/satu_sehat/models"
	"github.com/joho/godotenv"
)

type SatuSehatCarePlan interface {
	CreateCarePlan(patient models.CarePlantRequest, token string) (models.SatuSehatResponseReturn, error)
}

type satuSehatCarePlan struct {
	sql *sql.DB
}

func NewSatuSehatCarePlan(sql *sql.DB) SatuSehatCarePlan {
	return &satuSehatCarePlan{sql}
}

func (q *satuSehatCarePlan) CreateCarePlan(patient models.CarePlantRequest, token string) (models.SatuSehatResponseReturn, error) {
	_ = godotenv.Load()

	url := os.Getenv("SATU_SEHAT_END_POINT") + "/CarePlan"

	c, _ := json.Marshal(models.CarePlan{
		ResourceType: "CarePlan",
		Identifier: models.Identifier{
			System: "http://sys-ids.kemkes.go.id/careplan/" + os.Getenv("SATU_SEHAT_ORGANIZATION"),
			Value:  patient.CareNumber,
		},
		Intent: "plan",
		Title:  "Instruksi Medik dan Keperawatan Pasien",
		Status: "active",
		Category: []models.Category{
			{
				Coding: []models.Coding{
					{
						System:  "http://snomed.info/sct",
						Code:    "736271009",
						Display: "Outpatient care plan",
					},
				},
			},
		},
		Description: patient.Description,
		Subject: models.Reference{
			Reference: "Patient/" + patient.PatientID,
			Display:   patient.PatientName,
		},
		Encounter: models.Encounter2{
			Reference: "Encounter/" + patient.Encounter,
			Display:   "Kunjungan " + patient.PatientName + " pada tanggal " + patient.DateDisplay,
		},
		Created: patient.Date,
		Author: models.Reference{
			Reference: "Practitioner/" + patient.PractitionerID,
			Display:   patient.PractitionerName,
		},
	})

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(c))
	if err != nil {
		return models.SatuSehatResponseReturn{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return models.SatuSehatResponseReturn{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.SatuSehatResponseReturn{}, err
	}

	if resp.StatusCode == 201 {
		var data models.ConditionBodyResponse
		_ = json.Unmarshal(body, &data)

		return models.SatuSehatResponseReturn{Data: data.ID, Code: 201}, nil
	}

	var response models.SatuSehatErrorResponse
	_ = json.Unmarshal(body, &response)

	in := response.Issue[0].Details.Text

	return models.SatuSehatResponseReturn{Data: in, Code: 400}, nil
}
