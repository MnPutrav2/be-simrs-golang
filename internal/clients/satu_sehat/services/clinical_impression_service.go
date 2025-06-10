package services

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/MnPutrav2/be-simrs-golang/internal/models"
	"github.com/joho/godotenv"
)

type SatuSehatClinicalImpression interface {
	CreateClinicalImpression(patient models.ClinicalImpressionClientRequest, token string) ([]byte, error)
}

type satuSehatClinicalImpression struct {
	sql *sql.DB
}

func NewSatuSehatClinicalImpression(sql *sql.DB) SatuSehatClinicalImpression {
	return &satuSehatClinicalImpression{sql}
}

func (q *satuSehatClinicalImpression) CreateClinicalImpression(patient models.ClinicalImpressionClientRequest, token string) ([]byte, error) {
	_ = godotenv.Load()

	url := os.Getenv("SATU_SEHAT_END_POINT") + "/ClinicalImpression"

	c, _ := json.Marshal(models.ClinicalImpressionRequest{
		ResourceType: "ClinicalImpression",
		Status:       "completed",
		Description:  patient.Description,
		Subject: models.Encounter2{
			Reference: "Patient/" + patient.PatientID,
			Display:   patient.PatientName,
		},
		Encounter: models.Encounter2{
			Reference: "Encounter/" + patient.Encounter,
			Display:   "Kunjungan " + patient.PatientName + " pada tanggal " + patient.DateDisplay,
		},
		EffectiveDateTime: patient.Date,
		Date:              patient.Date,
		Assessor: models.Assessor{
			Reference: "Practitioner/" + patient.PractitionerID,
		},
		Summary: patient.Summary,
		Finding: []models.Finding{
			{
				ItemCodeableConcept: models.CodeableConcept{
					Coding: patient.Diagnosis,
				},
				ItemReference: models.Assessor{
					Reference: "Condition/" + patient.Condition,
				},
			},
		},
		PrognosisCodeableConcept: []models.CodeableConcept{
			{
				Coding: []models.Coding{
					{
						System:  "http://terminology.kemkes.go.id/CodeSystem/clinical-term",
						Code:    "PR000001",
						Display: "Prognosis",
					},
				},
			},
		},
	})

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(c))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 201 {
		var data models.ConditionBodyResponse
		_ = json.Unmarshal(body, &data)

		return []byte(data.ID), nil
	}

	return body, nil
}
