package models

type ResponseDataSuccess struct {
	Status   string `json:"status"`
	Response string `json:"response"`
}

type ResponseDataError struct {
	Status string `json:"status"`
	Errors string `json:"errors"`
}
