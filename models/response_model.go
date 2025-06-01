package models

import "github.com/google/uuid"

type ResponseDataSuccess struct {
	Status   string `json:"status"`
	Response string `json:"response"`
}

type ResponseDataError struct {
	Status string `json:"status"`
	Errors string `json:"errors"`
}

type AuthResponse struct {
	Status string    `json:"status"`
	Token  uuid.UUID `json:"token"`
}
