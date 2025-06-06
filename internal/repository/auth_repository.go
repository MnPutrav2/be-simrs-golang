package repository

type AuthRepository interface {
	CheckUserToken() error
}
