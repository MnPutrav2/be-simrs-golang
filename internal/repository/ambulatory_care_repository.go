package repository

import "github.com/MnPutrav2/be-simrs-golang/internal/models"

type AmbulatoryCareRepository interface {
	CreateAmbulatoryCareData(amb models.AmbulatoryCare) error
	DeleteAmbulatoryCareData(careNumber string, date string) error
}
