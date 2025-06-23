package repository

import (
	"database/sql"

	"github.com/MnPutrav2/be-simrs-golang/internal/models"
)

type PharmacyRepository interface {
	CreateDrugData(drug models.RequestBodyDrugData) error
	GetDrugData(search string, limit int) ([]models.ResponseDrugData, error)
	UpdateDrugData(drug models.RequestBodyDrugDataUpdate) error
	DeleteDrugData(id string) error
	GetDistributor() ([]models.Distributor, error)
}

type pharmacyRepository struct {
	sql *sql.DB
}

func NewPharmacyRepository(sql *sql.DB) PharmacyRepository {
	return &pharmacyRepository{sql}
}

func (q *pharmacyRepository) CreateDrugData(drug models.RequestBodyDrugData) error {
	_, err := q.sql.Exec("INSERT INTO drug_datas(id, name, distributor, capacity, fill, unit, price, category, expired_date) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)", drug.ID, drug.Name, drug.Distributor, drug.Capacity, drug.Fill, drug.Unit, drug.Price, drug.Category, drug.ExpiredDate)
	return err
}

func (q *pharmacyRepository) GetDrugData(search string, limit int) ([]models.ResponseDrugData, error) {
	result, err := q.sql.Query("SELECT drug_datas.id, drug_datas.name, distributor.id, distributor.name, drug_datas.capacity, drug_datas.fill, drug_datas.unit, drug_datas.composition, drug_datas.price, drug_datas.category, drug_datas.expired_date FROM drug_datas INNER JOIN distributor ON drug_datas.distributor = distributor.id WHERE drug_datas.name ILIKE $1 OR drug_datas.composition ILIKE $2 ORDER BY drug_datas.id DESC LIMIT $3", search, search, limit)
	if err != nil {
		return nil, err
	}

	var data []models.ResponseDrugData
	for result.Next() {
		var dt models.ResponseDrugData

		err := result.Scan(&dt.ID, &dt.Name, &dt.DistributorID, &dt.Distributor, &dt.Capacity, &dt.Fill, &dt.Unit, &dt.Composition, &dt.Price, &dt.Category, &dt.ExpiredDate)
		if err != nil {
			return nil, err
		}

		data = append(data, dt)
	}

	return data, nil
}

func (q *pharmacyRepository) UpdateDrugData(drug models.RequestBodyDrugDataUpdate) error {
	_, err := q.sql.Exec("UPDATE drug_datas SET id = $1, name = $2, distributor = $3, capacity = $4, fill = $5, unit = $6, composition = $7, price = $8, category = $9, expired_date = $10 WHERE id = $11;", drug.Data.ID, drug.Data.Name, drug.Data.Distributor, drug.Data.Capacity, drug.Data.Fill, drug.Data.Unit, drug.Data.Composition, drug.Data.Price, drug.Data.Category, drug.Data.ExpiredDate, drug.ID)
	return err
}

func (q *pharmacyRepository) DeleteDrugData(id string) error {
	_, err := q.sql.Exec("DELETE FROM drug_datas WHERE drug_datas.id = $1", id)
	return err
}

func (q *pharmacyRepository) GetDistributor() ([]models.Distributor, error) {
	result, err := q.sql.Query("SELECT id, name, address FROM distributor")
	if err != nil {
		return nil, err
	}

	var dis []models.Distributor
	for result.Next() {
		var dt models.Distributor

		err := result.Scan(&dt.ID, &dt.Name, &dt.Address)
		if err != nil {
			return nil, err
		}

		dis = append(dis, dt)
	}

	return dis, nil
}
