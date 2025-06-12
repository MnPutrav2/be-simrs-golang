package repository

import (
	"database/sql"

	"github.com/MnPutrav2/be-simrs-golang/internal/models"
)

type LogsRepository interface {
	GetLogsData(date1 string, date2 string) ([]models.Log, error)
}

type logsRepository struct {
	sql *sql.DB
}

func NewLogsRepository(sql *sql.DB) LogsRepository {
	return &logsRepository{sql}
}

func (q *logsRepository) GetLogsData(date1 string, date2 string) ([]models.Log, error) {
	result, err := q.sql.Query("SELECT id, users_id, level, message, path, create_at FROM logs WHERE create_at BETWEEN ? AND ? ORDER BY create_at DESC", date1, date2)
	if err != nil {
		return []models.Log{}, err
	}

	var data []models.Log
	for result.Next() {
		var dt models.Log

		err := result.Scan(&dt.ID, &dt.User, &dt.Level, &dt.Message, &dt.Path, &dt.Date)
		if err != nil {
			panic(err.Error())
		}

		data = append(data, dt)
	}

	return data, nil
}
