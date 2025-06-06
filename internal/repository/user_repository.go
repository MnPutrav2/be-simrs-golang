package repository

import "github.com/MnPutrav2/be-simrs-golang/internal/models"

type UserRepository interface {
	GetUserPagesData(token string, path string) ([]models.UserPages, error)
	GetUserStatus(token string, path string) (models.EmployeeData, error)
	UserLogout(token string) error
}
