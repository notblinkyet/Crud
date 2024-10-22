package models

import "time"

type Task struct {
	Id          int
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   time.Time
	LastUpdate  time.Time
}
