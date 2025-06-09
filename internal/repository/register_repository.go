package repository

import (
	"github.com/MnPutrav2/be-simrs-golang/internal/models"
)

type RegisterPatient interface {
	CreateRegistrationData(reg models.RequestRegisterPatient, path string) error
	DeleteRegistrationData(treatmentNumber string) error
	GetRegistrationData(date1 string, date2 string, limit string, search string) ([]models.ResponseRegisterPatient, error)
	GetCurrentRegisterNumber(date string, policlinic string) (int, error)
	GetCurrentCareNumber(date string) (int, error)
}
