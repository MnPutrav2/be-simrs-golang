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
	CreateRecipe(recipe models.RecipeRequest) (string, error)
	CreateRecipeCompound(recipe models.RecipeCompoundRequest) (string, error)
	GetCurrentRecipeNumber(date string) (int, error)
	AddRecipeNumber(recipe string) (string, error)
	GetRecipes(date1 string, date2 string) ([]models.RecipesData, error)
	GetDrugRecipes(recipe string) ([]models.RecipeType, error)
	DeleteDrugRecipes(recipe string, drug string, comname string) error
	ValidateRecipe(recipe string, validate models.ValidateRecipe) error
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

func (q *pharmacyRepository) CreateRecipe(recipe models.RecipeRequest) (string, error) {
	if recipe.Type == "create" {
		var recCheck bool
		err := q.sql.QueryRow("SELECT EXISTS (SELECT 1 FROM recipes WHERE recipe_id = $1)", recipe.RecipeNumber).Scan(&recCheck)
		if err != nil {
			return "", err
		}

		if recCheck {
			return "duplicate", nil
		}

		_, err = q.sql.Exec("INSERT INTO recipes(recipe_id, care_number, date, validate, validate_status, handover) VALUES($1, $2, $3, $4, $5, $6)", recipe.RecipeNumber, recipe.CareNumber, recipe.Date, recipe.Validate, "false", recipe.Handover)
		if err != nil {
			return "", err
		}

		for _, d := range recipe.Drug {
			_, err = q.sql.Exec("INSERT INTO detail_recipes (recipe_id, drug_id, validate_status, compound_name, recipe_type, value, use, embalming, tuslah, total_price) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)", recipe.RecipeNumber, d.DrugID, "false", "com", "common", d.Value, d.Use, d.Embalming, d.Tuslah, d.TotalPrice)
			if err != nil {
				return "", err
			}
		}

		return "", nil
	}

	var check bool
	err := q.sql.QueryRow("SELECT validate_status FROM recipes WHERE recipe_id = $1", recipe.RecipeNumber).Scan(&check)
	if err != nil {
		return "", err
	}

	if check {
		for _, d := range recipe.Drug {
			_, err := q.sql.Exec("INSERT INTO detail_recipes (recipe_id, drug_id, validate_status, compound_name, recipe_type, value, use, embalming, tuslah, total_price) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)", recipe.RecipeNumber, d.DrugID, "true", "com", "common", d.Value, d.Use, d.Embalming, d.Tuslah, d.TotalPrice)
			if err != nil {
				return "", err
			}
		}
	} else {
		for _, d := range recipe.Drug {
			_, err := q.sql.Exec("INSERT INTO detail_recipes (recipe_id, drug_id, validate_status, compound_name, recipe_type, value, use, embalming, tuslah, total_price) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)", recipe.RecipeNumber, d.DrugID, "false", "com", "common", d.Value, d.Use, d.Embalming, d.Tuslah, d.TotalPrice)
			if err != nil {
				return "", err
			}
		}
	}

	return "success", nil
}

func (q *pharmacyRepository) CreateRecipeCompound(recipe models.RecipeCompoundRequest) (string, error) {
	if recipe.Type == "create" {
		var recCheck bool
		err := q.sql.QueryRow("SELECT EXISTS (SELECT 1 FROM recipes WHERE recipe_id = $1)", recipe.RecipeNumber).Scan(&recCheck)
		if err != nil {
			return "", err
		}

		if recCheck {
			return "duplicate", nil
		}

		_, err = q.sql.Exec("INSERT INTO recipes(recipe_id, care_number, date, validate, validate_status, handover) VALUES($1, $2, $3, $4, $5, $6)", recipe.RecipeNumber, recipe.CareNumber, recipe.Date, recipe.Validate, "false", recipe.Handover)
		if err != nil {
			return "", err
		}

		for _, d := range recipe.Recipes {
			for _, x := range d.Drug {
				_, err = q.sql.Exec("INSERT INTO detail_recipes (recipe_id, drug_id, validate_status, compound_name, compound_value, recipe_type, value, use, embalming, tuslah, total_price) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)", recipe.RecipeNumber, x.DrugID, "false", d.RecipeName, d.Value, "compound", x.Value, d.Use, x.Embalming, x.Tuslah, x.Price)
				if err != nil {
					return "", err
				}
			}
		}

		return "success", nil
	}

	var check bool
	err := q.sql.QueryRow("SELECT validate_status FROM recipes WHERE recipe_id = $1", recipe.RecipeNumber).Scan(&check)
	if err != nil {
		return "", err
	}

	if check {
		for _, d := range recipe.Recipes {
			for _, x := range d.Drug {
				_, err = q.sql.Exec("INSERT INTO detail_recipes (recipe_id, drug_id, validate_status, compound_name, compound_value, recipe_type, value, use, embalming, tuslah, total_price) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)", recipe.RecipeNumber, x.DrugID, "true", d.RecipeName, d.Value, "compound", x.Value, d.Use, x.Embalming, x.Tuslah, x.Price)
				if err != nil {
					return "", err
				}
			}
		}
	} else {
		for _, d := range recipe.Recipes {
			for _, x := range d.Drug {
				_, err = q.sql.Exec("INSERT INTO detail_recipes (recipe_id, drug_id, validate_status, compound_name, compound_value, recipe_type, value, use, embalming, tuslah, total_price) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)", recipe.RecipeNumber, x.DrugID, "false", d.RecipeName, d.Value, "compound", x.Value, d.Use, x.Embalming, x.Tuslah, x.Price)
				if err != nil {
					return "", err
				}
			}
		}
	}

	return "success", nil
}

func (q *pharmacyRepository) GetCurrentRecipeNumber(date string) (int, error) {
	var current int
	dt := date + "%"
	err := q.sql.QueryRow("SELECT COUNT(*) FROM recipes WHERE date::date = $1", dt).Scan(&current)
	if err != nil {
		return 0, err
	}

	return current, nil
}

func (q *pharmacyRepository) AddRecipeNumber(care string) (string, error) {
	var rec string
	err := q.sql.QueryRow("SELECT recipe_id FROM recipes WHERE care_number = $1", care).Scan(&rec)
	if err != nil {
		return "", err
	}

	return rec, nil
}

func (q *pharmacyRepository) GetRecipes(date1 string, date2 string) ([]models.RecipesData, error) {
	result, err := q.sql.Query("SELECT recipes.recipe_id, recipes.care_number, patients.name, recipes.date, recipes.validate, recipes.validate_status, recipes.handover FROM recipes INNER JOIN registration ON recipes.care_number = registration.care_number INNER JOIN patients ON registration.medical_record = patients.medical_record WHERE recipes.date::date BETWEEN $1 AND $2", date1, date2)
	if err != nil {
		return nil, err
	}

	var recipes []models.RecipesData
	for result.Next() {
		var rec models.RecipesData

		err := result.Scan(&rec.RecipeId, &rec.CareNumber, &rec.Name, &rec.Date, &rec.Validate, &rec.ValidateStatus, &rec.Handover)
		if err != nil {
			return nil, err
		}

		recipes = append(recipes, rec)
	}

	return recipes, nil
}

func (q *pharmacyRepository) GetDrugRecipes(recipe string) ([]models.RecipeType, error) {
	var recipeDatas []models.RecipeType

	check, err := q.sql.Query("SELECT DISTINCT compound_name, recipe_type FROM detail_recipes WHERE recipe_id = $1", recipe)
	if err != nil {
		return nil, err
	}

	for check.Next() {
		var typ models.RecipeType

		err := check.Scan(&typ.CompoundName, &typ.RecipeType)
		if err != nil {
			return nil, err
		}

		result, err := q.sql.Query("SELECT detail_recipes.recipe_id, detail_recipes.drug_id, drug_datas.name, detail_recipes.validate_status, detail_recipes.compound_name,  COALESCE(detail_recipes.compound_value, 0) AS compoundvalue, detail_recipes.recipe_type, detail_recipes.value, detail_recipes.use, detail_recipes.embalming, detail_recipes.tuslah, detail_recipes.total_price FROM detail_recipes INNER JOIN drug_datas ON detail_recipes.drug_id = drug_datas.id WHERE detail_recipes.recipe_id = $1 AND detail_recipes.compound_name = $2", recipe, typ.CompoundName)
		if err != nil {
			return nil, err
		}

		var drugs []models.DetailRecipe
		for result.Next() {
			var drug models.DetailRecipe

			err := result.Scan(&drug.RecipeID, &drug.DrugID, &drug.DrugName, &drug.ValidateStatus, &drug.CompoundName, &drug.CompoundValue, &drug.RecipeType, &drug.Value, &drug.Use, &drug.Embalming, &drug.Tuslah, &drug.TotalPrice)
			if err != nil {
				return nil, err
			}

			drugs = append(drugs, drug)
		}

		typ.Data = append(typ.Data, drugs...)

		recipeDatas = append(recipeDatas, typ)
	}

	return recipeDatas, nil
}

func (q *pharmacyRepository) DeleteDrugRecipes(recipe string, drug string, comname string) error {
	_, err := q.sql.Exec("DELETE FROM detail_recipes WHERE recipe_id = $1 AND drug_id = $2 AND compound_name = $3", recipe, drug, comname)
	if err != nil {
		return err
	}

	var length int
	err = q.sql.QueryRow("SELECT COUNT(*) FROM detail_recipes WHERE recipe_id = $1", recipe).Scan(&length)
	if err != nil {
		return err
	}

	if length == 0 {
		_, err := q.sql.Exec("DELETE FROM recipes WHERE recipe_id = $1", recipe)
		if err != nil {
			return err
		}

		return nil
	}

	return nil
}

func (q *pharmacyRepository) ValidateRecipe(recipe string, validate models.ValidateRecipe) error {
	_, err := q.sql.Exec("UPDATE recipes SET validate_status = $1, validate = $2 WHERE recipe_id = $3", validate.ValidateStatus, validate.ValidateDate, recipe)
	if err != nil {
		return err
	}

	_, err = q.sql.Exec("UPDATE detail_recipes SET validate_status = $1 WHERE recipe_id = $2", validate.ValidateStatus, recipe)
	if err != nil {
		return err
	}

	return nil
}
