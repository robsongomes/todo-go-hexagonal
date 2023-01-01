package application

import "github.com/robsongomes/todo-go-hexagonal/domain"

type TodoInputPort struct {
	todoOutputPort TodoOutputPort
}

func NewTodoInputPort(todoOutput TodoOutputPort) *TodoInputPort {
	return &TodoInputPort{todoOutputPort: todoOutput}
}

func (t *TodoInputPort) Get(id string) (*domain.Todo, error) {
	return t.todoOutputPort.Get(id)
}

func (t *TodoInputPort) List() ([]domain.Todo, error) {
	return t.todoOutputPort.List()
}

func (t *TodoInputPort) Create(title, description string) (*domain.Todo, error) {
	id := domain.NewId()
	todo := domain.NewTodo(id, title, description)
	return t.todoOutputPort.Create(todo)
}
