package pkg

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/MnPutrav2/be-simrs-golang/internal/repository"
	"github.com/google/uuid"
)

func SessionToken(sql *sql.DB, us string, pas string) uuid.UUID {
	authRepo := repository.NewAuthRepository(sql)
	token := authRepo.CreateSessionToken(us, pas)

	return token
}

func CheckAuthorization(w http.ResponseWriter, path string, db *sql.DB, auth string) bool {
	split := strings.SplitN(auth, " ", 2)

	if len(split) != 2 || split[0] != "Bearer" {
		return false
	}

	if split[0] != "Bearer" {
		return false
	}

	token := split[1]

	authRepo := repository.NewAuthRepository(db)
	q := authRepo.CheckSessionToken(token)

	return q != 0
}
