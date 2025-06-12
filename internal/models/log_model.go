package models

type Log struct {
	ID      string `json:"id"`
	User    string `json:"user"`
	Level   string `json:"level"`
	Message string `json:"message"`
	Path    string `json:"path"`
	Date    string `json:"date"`
}
