package posgresql

import (
	"context"

	"github.com/notblinkyet/Crud/pkg/models"
)

func (storage *PostgresStorage) Create(task *models.Task) (int, error) {
	var id int
	conn := storage.conn
	ctx := context.Background()
	err := conn.QueryRow(ctx,
		"INSERT INTO tasks (title, description, status) VALUES($1, $2, $3) RETURNING id;",
		task.Title,
		task.Description,
		task.Status,
	).Scan(&id)

	if err != nil {
		return -1, err
	}

	return id, nil
}
