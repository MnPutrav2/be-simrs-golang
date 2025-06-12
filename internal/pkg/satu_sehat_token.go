package pkg

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/MnPutrav2/be-simrs-golang/internal/clients/satu_sehat/models"
	"github.com/MnPutrav2/be-simrs-golang/internal/helper"
	"github.com/joho/godotenv"
)

func CreateSatuSehatToken(db *sql.DB) (string, error) {
	_ = godotenv.Load()
	tm := time.Now()

	var count int
	_ = db.QueryRow("SELECT COUNT(*) FROM satu_sehat_token").Scan(&count)

	endpoint := os.Getenv("SATU_SEHAT_END_POINT_OAUTH") + "/accesstoken?grant_type=client_credentials"

	data := url.Values{}
	data.Set("client_id", os.Getenv("SATU_SEHAT_CLIENT_ID"))
	data.Set("client_secret", os.Getenv("SATU_SEHAT_CLIENT_SECRET"))

	if count == 0 {
		body := tokenCreate(endpoint, data)

		var result models.SatuSehatTokenResponseSucess
		_ = json.Unmarshal(body, &result)

		num, err := strconv.Atoi(result.ExpiresIn)
		if err != nil {
			panic(err.Error())
		}

		h := tm.Add(time.Duration(num) * time.Second).Format("2006-01-02 15:04:05")
		insert, err := db.Query("INSERT INTO satu_sehat_token(token, expired) VALUES(?, ?)", result.AccessToken, h)
		if err != nil {
			panic(err.Error())
		}

		defer insert.Close()

		fmt.Println(helper.Log("satu sehat token created : 201", "/access_token"))
		return result.AccessToken, nil
	}

	var token string
	var dt string
	err := db.QueryRow("SELECT satu_sehat_token.token, satu_sehat_token.expired FROM satu_sehat_token").Scan(&token, &dt)
	if err != nil {
		panic(err.Error())
	}

	exp, err := time.Parse("2006-01-02 15:04:05", dt)
	if err != nil {
		panic(err.Error())
	}

	if tm.After(exp) {
		if _, err := db.Exec("DELETE FROM satu_sehat_token"); err != nil {
			panic(err.Error())
		}

		body := tokenCreate(endpoint, data)

		var result models.SatuSehatTokenResponseSucess
		_ = json.Unmarshal(body, &result)

		num, err := strconv.Atoi(result.ExpiresIn)
		if err != nil {
			panic(err.Error())
		}

		h := tm.Add(time.Duration(num) * time.Second).Format("2006-01-02 15:04:05")
		insert, err := db.Query("INSERT INTO satu_sehat_token(token, expired) VALUES(?, ?)", result.AccessToken, h)
		if err != nil {
			panic(err.Error())
		}

		defer insert.Close()

		fmt.Println(helper.Log("satu sehat token created : 201", "/access_token"))
		return result.AccessToken, nil
	}

	fmt.Println(helper.Log("satu sehat token available : 200", "/access_token"))

	return token, nil
}

func tokenCreate(endpoint string, data url.Values) []byte {
	resp, err := http.Post(endpoint, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		panic(err.Error())
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	return body
}
