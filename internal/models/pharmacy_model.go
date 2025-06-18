package models

import "time"

type DrugData struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Distributor string    `json:"distributor"`
	Capacity    string    `json:"capacity"`
	Unit        string    `json:"unit"`
	Category    string    `json:"category"`
	Price       int       `json:"price"`
	ExpiredDate time.Time `json:"expired_date"`
}
