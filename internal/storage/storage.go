package storage

import "simple_crud/internal/models"

type Storage interface {
	Create(task *models.Task) (int, error)
	Delete(id int) error
	ReadId(id int) (*models.Task, error)
	ReadAll() ([]models.Task, error)
	Update(id int, title string, description string, status string) error
	Close() error
}
