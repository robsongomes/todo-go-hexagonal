package application

import "github.com/robsongomes/todo-go-hexagonal/domain"

type TodoOutputPort interface {
	Get(id string) (*domain.Todo, error)
	List() ([]domain.Todo, error)
	Create(todo *domain.Todo) (*domain.Todo, error)
}
