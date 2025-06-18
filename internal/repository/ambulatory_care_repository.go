package repository

import "github.com/MnPutrav2/be-simrs-golang/internal/models"

type AmbulatoryCareRepository interface {
	CreateAmbulatoryCareData(amb models.RequestAmbulatoryCare) error
	DeleteAmbulatoryCareData(careNumber string, date string) error
	GetAmbulatoryCareData(careNumber string, date1 string, date2 string) ([]models.ResponseAmbulatoryCare, error)
	UpdateAmbulatoryCareData(data models.RequestUpdateAmbulatorCare) error
}
