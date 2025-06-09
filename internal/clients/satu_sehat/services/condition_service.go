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

type SatuSehatConditionService interface {
	CreateSatuSehatCondition(patient models.ConditionClientRequest, token string) ([]byte, error)
}

type satuSehatConditionService struct {
	sql *sql.DB
}

func NewSatuSehatCondition(sql *sql.DB) SatuSehatConditionService {
	return &satuSehatConditionService{sql}
}

func (q *satuSehatConditionService) CreateSatuSehatCondition(patient models.ConditionClientRequest, token string) ([]byte, error) {
	_ = godotenv.Load()

	url := os.Getenv("SATU_SEHAT_END_POINT") + "/Condition"

	c, _ := json.Marshal(models.ConditionRequest{
		ResourceType: "Condition",
		ClinicalStatus: models.ClinicalStatus{
			Coding: []models.Coding{
				{
					System:  "http://terminology.hl7.org/CodeSystem/condition-clinical",
					Code:    "active",
					Display: "Active",
				},
			},
		},
		Category: []models.Category{
			{
				Coding: []models.Coding{
					{
						System:  "http://terminology.hl7.org/CodeSystem/condition-category",
						Code:    "encounter-diagnosis",
						Display: "Encounter Diagnosis",
					},
				},
			},
		},
		Code: models.Code{
			Coding: patient.Diagnosis,
		},
		Subject: models.Subject{
			Reference: "Patient/" + patient.PatientID,
			Display:   patient.PatientName,
		},
		Encounter: models.Encounter{
			Reference: "Encounter/" + patient.Encounter,
		},
		OnSetDateTime: patient.SetDate,
		RecordedDate:  patient.RecorderDate,
	})

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(c))
	if err != nil {
		return nil, err
	}

	// Tambahkan header Content-Type dan Authorization Bearer
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	// Kirim request
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

	return body, nil
}
