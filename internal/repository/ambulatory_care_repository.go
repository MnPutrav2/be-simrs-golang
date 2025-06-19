package repository

import (
	"database/sql"
	"net/http"

	"github.com/MnPutrav2/be-simrs-golang/internal/models"
)

type AmbulatoryCareRepository interface {
	CreateAmbulatoryCareData(amb models.RequestAmbulatoryCare) error
	DeleteAmbulatoryCareData(careNumber string, date string) error
	GetAmbulatoryCareData(careNumber string, date1 string, date2 string) ([]models.ResponseAmbulatoryCare, error)
	UpdateAmbulatoryCareData(data models.RequestUpdateAmbulatorCare) error
}

type ambulatoryCareRepository struct {
	sql *sql.DB
	w   http.ResponseWriter
	r   *http.Request
}

func NewAmbulatoryCareRepository(sql *sql.DB, w http.ResponseWriter, r *http.Request) AmbulatoryCareRepository {
	return &ambulatoryCareRepository{sql, w, r}
}

func (q *ambulatoryCareRepository) CreateAmbulatoryCareData(amb models.RequestAmbulatoryCare) error {
	_, err := q.sql.Exec("INSERT INTO ambulatory_care(care_number, medical_record, date, body_temperature, tension, pulse, respiration, height, weight, spo2, gcs, awareness, complaint, examination, allergy, followup, assessment, instructions, evaluation, officer) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20);", amb.CareNumber, amb.MedicalRecord, amb.Date, amb.BodyTemperature, amb.Tension, amb.Pulse, amb.Respiration, amb.Height, amb.Weight, amb.Spo2, amb.GCS, amb.Awareness, amb.Complaint, amb.Examination, amb.Allergy, amb.FollowUp, amb.Assessment, amb.Instructions, amb.Evaluation, amb.Officer)
	return err
}

func (q *ambulatoryCareRepository) DeleteAmbulatoryCareData(careNumber string, date string) error {
	_, err := q.sql.Exec("DELETE FROM ambulatory_care WHERE care_number = $1 AND date = $2", careNumber, date)
	return err
}

func (q *ambulatoryCareRepository) GetAmbulatoryCareData(careNumber string, date1 string, date2 string) ([]models.ResponseAmbulatoryCare, error) {
	result, err := q.sql.Query("SELECT ambulatory_care.care_number, ambulatory_care.medical_record, patients.name, ambulatory_care.date, ambulatory_care.body_temperature, ambulatory_care.tension, ambulatory_care.pulse, ambulatory_care.respiration, ambulatory_care.height, ambulatory_care.weight, ambulatory_care.spo2, ambulatory_care.gcs, ambulatory_care.awareness, ambulatory_care.complaint, ambulatory_care.examination, ambulatory_care.allergy, ambulatory_care.followup, ambulatory_care.assessment, ambulatory_care.instructions, ambulatory_care.evaluation, ambulatory_care.officer, employees.name FROM ambulatory_care INNER JOIN employees ON ambulatory_care.officer = employees.id INNER JOIN patients ON ambulatory_care.medical_record = patients.medical_record WHERE (care_number = $1 OR $2 = '') AND date BETWEEN $3 AND $4 ORDER BY date DESC", careNumber, careNumber, date1, date2)
	if err != nil {
		return nil, err
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

func (q *ambulatoryCareRepository) UpdateAmbulatoryCareData(patient models.RequestUpdateAmbulatorCare) error {
	_, err := q.sql.Exec("UPDATE ambulatory_care SET medical_record = $1, date = $2, body_temperature = $3, tension = $4, pulse = $5, respiration = $6, height = $7, weight = $8, spo2 = $9, gcs = $10, awareness = $11, complaint = $12, examination = $13, allergy = $14, followup = $15, assessment = $16, instructions = $17, evaluation = $18, officer = $19 WHERE care_number = $20 AND date = $21;", patient.Data.MedicalRecord, patient.Data.Date, patient.Data.BodyTemperature, patient.Data.Tension, patient.Data.Pulse, patient.Data.Respiration, patient.Data.Height, patient.Data.Weight, patient.Data.Spo2, patient.Data.GCS, patient.Data.Awareness, patient.Data.Complaint, patient.Data.Examination, patient.Data.Allergy, patient.Data.FollowUp, patient.Data.Assessment, patient.Data.Instructions, patient.Data.Evaluation, patient.Data.Officer, patient.CareNumber, patient.Date)
	return err
}
