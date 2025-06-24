package models

type RequestBodyDrugData struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Distributor string `json:"distributor"`
	Capacity    int    `json:"capacity"`
	Fill        int    `json:"fill"`
	Unit        string `json:"unit"`
	Composition string `json:"composition"`
	Category    string `json:"category"`
	Price       int    `json:"price"`
	ExpiredDate string `json:"expired_date"`
}

type RequestBodyDrugDataUpdate struct {
	ID   string              `json:"id"`
	Data RequestBodyDrugData `json:"data"`
}

type RecipeRequest struct {
	CareNumber   string       `json:"care_number"`
	RecipeNumber string       `json:"recipe_number"`
	Date         string       `json:"date"`
	Validate     string       `json:"validate"`
	Handover     string       `json:"handover"`
	Type         string       `json:"type"`
	Drug         []RecipeDrug `json:"drug"`
}

type RecipeCompoundRequest struct {
	CareNumber   string           `json:"care_number"`
	RecipeNumber string           `json:"recipe_number"`
	Date         string           `json:"date"`
	Validate     string           `json:"validate"`
	Handover     string           `json:"handover"`
	Type         string           `json:"type"`
	Recipes      []RecipeCompound `json:"recipes"`
}

type RecipeCompound struct {
	RecipeName string               `json:"recipe_name"`
	Use        string               `json:"use"`
	Value      int                  `json:"value"`
	Drug       []RecipeCompoundDrug `json:"drug"`
}

type RecipeCompoundDrug struct {
	Name      string `json:"name"`
	DrugID    string `json:"drug_id"`
	Value     int    `json:"value"`
	Embalming int    `json:"embalming"`
	Tuslah    int    `json:"tuslah"`
	Price     int    `json:"price"`
}

type RecipeDrug struct {
	RecipeID   string `json:"recipe_id"`
	DrugID     string `json:"drug_id"`
	Value      int    `json:"value"`
	Use        string `json:"use"`
	Embalming  int    `json:"embalming"`
	Tuslah     int    `json:"tuslah"`
	TotalPrice int    `json:"total_price"`
}

type ResponseDrugData struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	DistributorID string `json:"distributor_id"`
	Distributor   string `json:"distributor"`
	Capacity      int    `json:"capacity"`
	Fill          int    `json:"fill"`
	Unit          string `json:"unit"`
	Composition   string `json:"composition"`
	Category      string `json:"category"`
	Price         int    `json:"price"`
	ExpiredDate   string `json:"expired_date"`
}

type Distributor struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
