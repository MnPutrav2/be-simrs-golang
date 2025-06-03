package pkg

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func Cors(w http.ResponseWriter, r *http.Request) bool {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("ALLOW_ORIGIN"))
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, DELETE, PATCH")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return true
	}

	return false

}
