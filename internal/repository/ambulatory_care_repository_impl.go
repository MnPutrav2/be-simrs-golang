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

func (q *ambulatoryCareRepository) CreateAmbulatoryCareData(amb models.RequestAmbulatoryCare) error {
	_, err := q.sql.Exec(`INSERT INTO ambulatory_care(care_number, medical_record, date, body_temperature, tension, pulse, respiration, height, weight, spo2, gcs, awareness, complaint, examination, allergy, followup, assessment, instructions, evaluation, officer) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`, amb.CareNumber, amb.MedicalRecord, amb.Date, amb.BodyTemperature, amb.Tension, amb.Pulse, amb.Respiration, amb.Height, amb.Weight, amb.Spo2, amb.GCS, amb.Awareness, amb.Complaint, amb.Examination, amb.Allergy, amb.FollowUp, amb.Assessment, amb.Instructions, amb.Evaluation, amb.Officer)
	return err
}

func (q *ambulatoryCareRepository) DeleteAmbulatoryCareData(careNumber string, date string) error {
	_, err := q.sql.Exec("DELETE FROM ambulatory_care WHERE care_number = ? AND date = ?", careNumber, date)
	return err
}

func (q *ambulatoryCareRepository) GetAmbulatoryCareData(careNumber string, date1 string, date2 string) ([]models.ResponseAmbulatoryCare, error) {
	result, err := q.sql.Query("SELECT ambulatory_care.care_number, ambulatory_care.medical_record, patients.name, ambulatory_care.date, ambulatory_care.body_temperature, ambulatory_care.tension, ambulatory_care.pulse, ambulatory_care.respiration, ambulatory_care.height, ambulatory_care.weight, ambulatory_care.spo2, ambulatory_care.gcs, ambulatory_care.awareness, ambulatory_care.complaint, ambulatory_care.examination, ambulatory_care.allergy, ambulatory_care.followup, ambulatory_care.assessment, ambulatory_care.instructions, ambulatory_care.evaluation, ambulatory_care.officer, employee.name FROM ambulatory_care INNER JOIN employee ON ambulatory_care.officer = employee.id INNER JOIN patients ON ambulatory_care.medical_record = patients.medical_record WHERE (care_number = ? OR ? = '') AND date BETWEEN ? AND ? ORDER BY date DESC", careNumber, careNumber, date1, date2)
	if err != nil {
		panic(err.Error())
	}

	var datas []models.ResponseAmbulatoryCare

	for result.Next() {
		var amb models.ResponseAmbulatoryCare

		err = result.Scan(&amb.CareNumber, &amb.MedicalRecord, &amb.Name, &amb.Date, &amb.BodyTemperature, &amb.Tension, &amb.Pulse, &amb.Respiration, &amb.Height, &amb.Weight, &amb.Spo2, &amb.GCS, &amb.Awareness, &amb.Complaint, &amb.Examination, &amb.Allergy, &amb.FollowUp, &amb.Assessment, &amb.Instructions, &amb.Evaluation, &amb.Officer, &amb.OfficerName)
		if err != nil {
			panic(err.Error())
		}

		datas = append(datas, amb)
	}

	return datas, nil
}
