package repository

import "github.com/MnPutrav2/be-simrs-golang/internal/models"

type PharmacyRepository interface {
	CreateDrugData(drug models.DrugData) error
}
