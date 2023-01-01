package framework

import (
	"fmt"

	"github.com/robsongomes/todo-go-hexagonal/domain"
)

type todoMemory struct {
	ID          string
	Title       string
	Description string
}

type todoListMemory []todoMemory

type TodoMemoryAdapter struct {
	todos todoListMemory
}

func (m *todoMemory) ToDomain() *domain.Todo {
	return &domain.Todo{
		ID:          m.ID,
		Title:       m.Title,
		Description: m.Description,
	}
}
func (m *todoMemory) FromDomain(todo *domain.Todo) {
	if m == nil {
		m = &todoMemory{}
	}

	m.ID = todo.ID
	m.Title = todo.Title
	m.Description = todo.Description
}

func (m todoListMemory) ToDomain() []domain.Todo {
	todos := make([]domain.Todo, len(m))
	for k, td := range m {
		todo := td.ToDomain()
		todos[k] = *todo
	}

	return todos
}

func NewMemoryAdapter() *TodoMemoryAdapter {
	return &TodoMemoryAdapter{todos: todoListMemory{}}
}

func (a *TodoMemoryAdapter) Get(id string) (*domain.Todo, error) {
	var todo todoMemory
	for _, t := range a.todos {
		if t.ID == id {
			todo = t
		}
	}

	if todo.ID == "" {
		return nil, fmt.Errorf("todo %s not found", id)
	}

	return todo.ToDomain(), nil
}

func (a *TodoMemoryAdapter) List() ([]domain.Todo, error) {
	return a.todos.ToDomain(), nil
}

func (a *TodoMemoryAdapter) Create(todo *domain.Todo) (*domain.Todo, error) {
	t := todoMemory{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
	}

	a.todos = append(a.todos, t)

	return t.ToDomain(), nil
}
