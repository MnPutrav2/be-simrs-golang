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
	_, err := q.sql.Exec("INSERT INTO drug_datas(id, name, distributor, capacity, fill, unit, price, category, expired_date) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)", drug.ID, drug.Name, drug.Distributor, drug.Capacity, drug.Fill, drug.Unit, drug.Price, drug.Category, drug.ExpiredDate)
	return err
}

func (q *pharmacyRepository) GetDrugData(search string, limit int) ([]models.ResponseDrugData, error) {
	result, err := q.sql.Query("SELECT drug_datas.id, drug_datas.name, distributor.id, distributor.name, drug_datas.capacity, drug_datas.fill, drug_datas.unit, drug_datas.price, drug_datas.category, drug_datas.expired_date FROM drug_datas INNER JOIN distributor ON drug_datas.distributor = distributor.id WHERE drug_datas.name LIKE ? ORDER BY drug_datas.id DESC LIMIT ?", search, limit)
	if err != nil {
		return nil, err
	}

	var data []models.ResponseDrugData
	for result.Next() {
		var dt models.ResponseDrugData

		err := result.Scan(&dt.ID, &dt.Name, &dt.DistributorID, &dt.Distributor, &dt.Capacity, &dt.Fill, &dt.Unit, &dt.Price, &dt.Category, &dt.ExpiredDate)
		if err != nil {
			return nil, err
		}

		data = append(data, dt)
	}

	return data, nil
}

func (q *pharmacyRepository) UpdateDrugData(drug models.RequestBodyDrugDataUpdate) error {
	_, err := q.sql.Exec("UPDATE drug_datas SET id = ?, name = ?, distributor = ?, capacity = ?, fill = ?, unit = ?, price = ?, category = ?, expired_date = ? WHERE id = ?", drug.Data.ID, drug.Data.Name, drug.Data.Distributor, drug.Data.Capacity, drug.Data.Fill, drug.Data.Unit, drug.Data.Price, drug.Data.Category, drug.Data.ExpiredDate, drug.ID)
	return err
}

func (q *pharmacyRepository) DeleteDrugData(id string) error {
	_, err := q.sql.Exec("DELETE FROM drug_datas WHERE drug_datas.id = ?", id)
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
