package domain

import (
	"fmt"

	"github.com/google/uuid"
)

type Todo struct {
	ID          string //It could be a VO (ID)
	Title       string
	Description string
}

func NewTodo(id, title, description string) *Todo {
	return &Todo{
		ID:          id,
		Title:       title,
		Description: description,
	}
}

func NewId() string {
	return uuid.NewString()
}

func (t *Todo) String() string {
	return fmt.Sprintf("%s - %s", t.Title, t.Description)
}
