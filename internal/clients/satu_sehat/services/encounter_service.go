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

type SatuSehatEncounter interface {
	CreateEncounterData(patient models.EncounterResponse, token string) ([]byte, error)
}

type satuSehatEncounter struct {
	sql *sql.DB
}

func NewSatuSehatEncounter(sql *sql.DB) SatuSehatEncounter {
	return &satuSehatEncounter{sql}
}

type Polic struct {
	ID          int
	Name        string
	SatuSehatID string
}

func (q *satuSehatEncounter) CreateEncounterData(patient models.EncounterResponse, token string) ([]byte, error) {
	_ = godotenv.Load()

	url := os.Getenv("SATU_SEHAT_END_POINT") + "/Encounter"

	var polic Polic
	_ = q.sql.QueryRow("SELECT name, satu_sehat_id FROM policlinics WHERE id = ?", patient.LocationID).Scan(&polic.Name, &polic.SatuSehatID)

	c, err := json.Marshal(models.EncounterBody{
		ResourceType: "Encounter",
		Identifier: []models.Identifier{
			{
				System: "http://sys-ids.kemkes.go.id/encounter/" + os.Getenv("SATU_SEHAT_ORGANIZATION"),
				Value:  patient.CareNumber,
			},
		},
		Status: "arrived",
		Class: models.Class{
			System:  "http://terminology.hl7.org/CodeSystem/v3-ActCode",
			Code:    "AMB",
			Display: "ambulatory",
		},
		Subject: models.Subject{
			Reference: "Patient/" + patient.PatientID,
			Display:   patient.PatientName,
		},
		Participant: []models.Participan{
			{
				Individual: models.Subject{
					Reference: "Practitioner/" + patient.PractitionerID,
					Display:   patient.PractitionerName,
				},
				Type: []models.Type{
					{
						Coding: []models.Class{
							{
								Code:    "ATND",
								Display: "attender",
								System:  "http://terminology.hl7.org/CodeSystem/v3-ParticipationType",
							},
						},
					},
				},
			},
		},
		Period: models.Period{
			Start: patient.Start,
		},
		Location: []models.Location{
			{
				Location: models.Subject{
					Reference: "Location/" + polic.SatuSehatID,
					Display:   polic.Name,
				},
				Period: models.Period{
					Start: patient.Start,
				},
				Extension: []models.Extension{
					{
						URL: "https://fhir.kemkes.go.id/r4/StructureDefinition/ServiceClass",
						Extension: []models.Extension2{
							{
								URL: "value",
								ValueCodeableConcept: models.Type{
									Coding: []models.Class{
										{
											System:  "http://terminology.kemkes.go.id/CodeSystem/locationServiceClass-Outpatient",
											Code:    "reguler",
											Display: "Kelas Reguler",
										},
									},
								},
							},
							{
								URL: "upgradeClassIndicator",
								ValueCodeableConcept: models.Type{
									Coding: []models.Class{
										{
											System:  "http://terminology.kemkes.go.id/CodeSystem/locationUpgradeClass",
											Code:    "kelas-tetap",
											Display: "Kelas Tetap Perawatan",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		StatusHistory: []models.StatusHistory{
			{
				Status: "arrived",
				Period: models.Period{
					Start: patient.Start,
				},
			},
		},
		ServiceProvider: models.ServiceProvider{
			Reference: "Organization/" + os.Getenv("SATU_SEHAT_ORGANIZATION"),
		},
	})

	if err != nil {
		return nil, err
	}

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

	if resp.StatusCode == 201 {
		var data models.EncounterBodyResponse
		_ = json.Unmarshal(body, &data)

		return []byte(data.ID), nil
	}

	return body, nil
}
