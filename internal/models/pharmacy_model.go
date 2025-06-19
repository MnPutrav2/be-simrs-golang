package models

type RequestBodyDrugData struct {
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

type RequestBodyDrugDataUpdate struct {
	ID   string              `json:"id"`
	Data RequestBodyDrugData `json:"data"`
}

type ResponseDrugData struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	DistributorID string `json:"distributor_id"`
	Distributor   string `json:"distributor"`
	Capacity      int    `json:"capacity"`
	Fill          int    `json:"fill"`
	Unit          string `json:"unit"`
	Category      string `json:"category"`
	Price         int    `json:"price"`
	ExpiredDate   string `json:"expired_date"`
}

type Distributor struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
