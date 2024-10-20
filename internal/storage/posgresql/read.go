package posgresql

import (
	"context"
	"simple_crud/internal/models"
)

func (storage *PostgresStorage) ReadId(id int) (*models.Task, error) {
	var task models.Task
	conn := storage.conn
	ctx := context.Background()
	err := conn.QueryRow(ctx,
		"SELECT id, title, description, status, time_create, time_last_update FROM tasks WHERE id = $1;",
		id,
	).Scan(
		&task.Id,
		&task.Title,
		&task.Description,
		&task.Status,
		&task.CreatedAt,
		&task.LastUpdate,
	)
	return &task, err
}

func (storage *PostgresStorage) ReadAll() ([]models.Task, error) {
	tasks := make([]models.Task, 0, 10)
	ctx := context.Background()
	conn := storage.conn
	rows, err := conn.Query(ctx,
		"SELECT id, title, description, status, time_create, time_last_update FROM tasks;",
	)
	if err != nil {
		return tasks, err
	}
	defer rows.Close()
	for rows.Next() {
		var task models.Task
		err := rows.Scan(
			&task.Id,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.CreatedAt,
			&task.LastUpdate,
		)
		if err != nil {
			return tasks, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
