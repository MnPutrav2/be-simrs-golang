package repository

import "github.com/MnPutrav2/be-simrs-golang/internal/models"

type PatientRepository interface {
	CreatePatientData(patient models.PatientData, token string, path string) error
	GetPatientData(limit string, search string, token string, path string) ([]models.PatientData, error)
	DeletePatientData(mr string) error
	UpdatePatientData(patient models.PatientDataUpdate) error
	GetCurrentMedicalRecord() string
}
