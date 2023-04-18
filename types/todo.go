package types

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `gorm:"primaryKey" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"primaryKey" json:"updated_at"`
	DeleteAt    gorm.DeletedAt `gorm:"index" json:"delete_at"`
	Done        bool           `json:"done"`
}

type CreateTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func NewTodo(t, d string) *Todo {
	return &Todo{
		Title:       t,
		Description: d,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		Done:        false,
	}
}
