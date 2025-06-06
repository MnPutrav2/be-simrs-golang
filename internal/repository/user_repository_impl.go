package repository

import (
	"database/sql"
	"net/http"

	"github.com/MnPutrav2/be-simrs-golang/internal/helper"
	"github.com/MnPutrav2/be-simrs-golang/internal/models"
)

type userRepository struct {
	w   http.ResponseWriter
	r   *http.Request
	sql *sql.DB
}

func NewUserRepository(w http.ResponseWriter, r *http.Request, sql *sql.DB) UserRepository {
	return &userRepository{w, r, sql}
}

func (q *userRepository) GetUserPagesData(token string, path string) ([]models.UserPages, error) {
	var id int

	err := q.sql.QueryRow("SELECT session_token.users_id FROM session_token WHERE session_token.token = ?", token).Scan(&id)
	if err != nil {
		helper.ResponseError(q.w, "unauthorization", "unauthorization : 400", 401, path)
		return nil, err
	}

	result, err := q.sql.Query("SELECT user_pages.name, user_pages.path FROM user_pages WHERE user_pages.users_id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	var pageList []models.UserPages

	for result.Next() {
		var p models.UserPages

		err := result.Scan(&p.Name, &p.Path)
		if err != nil {
			panic(err.Error())
		}

		pageList = append(pageList, p)
	}

	return pageList, err
}

func (q *userRepository) GetUserStatus(token string, path string) (models.EmployeeData, error) {
	var id int

	err := q.sql.QueryRow("SELECT session_token.users_id FROM session_token WHERE session_token.token = ?", token).Scan(&id)
	if err != nil {
		helper.ResponseError(q.w, "unauthorization", "unauthorization : 400", 401, path)
		return models.EmployeeData{}, err
	}

	var user models.EmployeeData

	err = q.sql.QueryRow("SELECT employee.id, employee.name, employee.gender FROM employee INNER JOIN users ON employee.id = users.employee_id WHERE users.id = ?", id).Scan(&user.Employee_ID, &user.Name, &user.Gender)
	if err != nil {
		helper.ResponseError(q.w, "employee data not found", "employee data not found : 404", 404, path)
		return models.EmployeeData{}, err
	}

	return user, err
}

func (q *userRepository) UserLogout(token string) error {
	_, err := q.sql.Exec("DELETE FROM session_token WHERE session_token.token = ?", token)
	return err
}
