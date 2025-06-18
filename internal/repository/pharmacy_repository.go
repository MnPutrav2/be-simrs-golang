package repository

import (
	"database/sql"

	"github.com/MnPutrav2/be-simrs-golang/internal/models"
)

type PharmacyRepository interface {
	CreateDrugData(drug models.DrugData) error
}

type pharmacyRepository struct {
	sql *sql.DB
}

func NewPharmacyRepository(sql *sql.DB) PharmacyRepository {
	return &pharmacyRepository{sql}
}

func (q *pharmacyRepository) CreateDrugData(drug models.DrugData) error {
	_, err := q.sql.Exec("INSERT INTO drug_datas(id, name, distributor, capacity, fill, unit, price, category, expired_date) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)", drug.ID, drug.Name, drug.Distributor, drug.Capacity, drug.Fill, drug.Unit, drug.Price, drug.Category, drug.ExpiredDate)
	return err
}
