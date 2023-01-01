package application

import "github.com/robsongomes/todo-go-hexagonal/domain"

type TodoUseCase interface {
	Get(id string) (*domain.Todo, error)
	List() ([]domain.Todo, error)
	Create(title, description string) (*domain.Todo, error)
}
