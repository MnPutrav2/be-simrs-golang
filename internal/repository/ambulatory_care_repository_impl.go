package repository

import (
	"database/sql"
	"net/http"

	"github.com/MnPutrav2/be-simrs-golang/internal/models"
)

type ambulatoryCareRepository struct {
	sql *sql.DB
	w   http.ResponseWriter
	r   *http.Request
}

func NewAmbulatoryCareRepository(sql *sql.DB, w http.ResponseWriter, r *http.Request) AmbulatoryCareRepository {
	return &ambulatoryCareRepository{sql, w, r}
}

func (q *ambulatoryCareRepository) CreateAmbulatoryCareData(amb models.AmbulatoryCare) error {
	_, err := q.sql.Exec(`INSERT INTO ambulatory_care(care_number, date, body_temperature, tension, pulse, respiration, height, weight, spo2, gcs, awareness, complaint, examination, allergy, followup, assessment, instructions, evaluation, officer) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`, amb.CareNumber, amb.Date, amb.BodyTemperature, amb.Tension, amb.Pulse, amb.Respiration, amb.Height, amb.Weight, amb.Spo2, amb.GCS, amb.Awareness, amb.Complaint, amb.Examination, amb.Allergy, amb.FollowUp, amb.Assessment, amb.Instructions, amb.Evaluation, amb.Officer)
	return err
}

func (q *ambulatoryCareRepository) DeleteAmbulatoryCareData(careNumber string, date string) error {
	_, err := q.sql.Exec("DELETE FROM ambulatory_care WHERE care_number = ? AND date = ?", careNumber, date)
	return err
}
