package repository

import (
	"database/sql"
	"net/http"

	"github.com/MnPutrav2/be-simrs-golang/internal/models"
)

type UserRepository interface {
	GetUserPagesData(token string, path string) ([]models.UserPages, error)
	GetUserStatus(token string, path string) (models.EmployeeData, error)
	UserLogout(token string) error
}

type userRepository struct {
	w   http.ResponseWriter
	r   *http.Request
	sql *sql.DB
}

func NewUserRepository(w http.ResponseWriter, r *http.Request, sql *sql.DB) UserRepository {
	return &userRepository{w, r, sql}
}

func (q *userRepository) GetUserPagesData(token string, path string) ([]models.UserPages, error) {
	var id string

	err := q.sql.QueryRow("SELECT session_token.users_id FROM session_token WHERE session_token.token = $1", token).Scan(&id)
	if err != nil {
		return nil, err
	}

	result, err := q.sql.Query("SELECT user_pages.name, user_pages.path FROM user_pages WHERE user_pages.users_id = $1", id)
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
	var id string

	err := q.sql.QueryRow("SELECT session_token.users_id FROM session_token WHERE session_token.token = $1", token).Scan(&id)
	if err != nil {
		return models.EmployeeData{}, err
	}

	var user models.EmployeeData

	err = q.sql.QueryRow("SELECT employees.id, employees.name, employees.gender, employees.birth_place, employees.birth_date, employees.address, employees.village, employees.district, employees.regencie, employees.province, employees.nik, employees.bpjs, employees.npwp, employees.phone_number, employees.email FROM employees INNER JOIN users ON employees.id = users.employee_id WHERE users.id = $1", id).Scan(&user.Employee_ID, &user.Name, &user.Gender, &user.BirthPlace, &user.BirthDate, &user.Address, &user.Village, &user.District, &user.Regencie, &user.Province, &user.NIK, &user.BPJS, &user.NPWP, &user.PhoneNumber, &user.Email)
	if err != nil {
		return models.EmployeeData{}, err
	}

	return user, err
}

func (q *userRepository) UserLogout(token string) error {
	_, err := q.sql.Exec("DELETE FROM session_token WHERE session_token.token = $1", token)
	return err
}
