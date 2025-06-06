package repository

import "github.com/google/uuid"

type AuthRepository interface {
	CheckUserToken() error
	CreateSessionToken(user string, pass string) uuid.UUID
	CheckSessionToken(token string) int
}
