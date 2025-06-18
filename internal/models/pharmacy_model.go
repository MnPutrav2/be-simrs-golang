package models

type DrugData struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Distributor string `json:"distributor"`
	Capacity    int    `json:"capacity"`
	Fill        int    `json:"fill"`
	Unit        string `json:"unit"`
	Category    string `json:"category"`
	Price       int    `json:"price"`
	ExpiredDate string `json:"expired_date"`
}
