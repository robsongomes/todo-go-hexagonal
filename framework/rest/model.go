package rest

import "github.com/robsongomes/todo-go-hexagonal/domain"

type Todo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TodoList []Todo

func (td *Todo) FromDomain(todo *domain.Todo) {
	if td == nil {
		td = &Todo{}
	}

	td.ID = todo.ID
	td.Title = todo.Title
	td.Description = todo.Description
}

func (td *Todo) ToDomain() *domain.Todo {
	if td == nil {
		td = &Todo{}
	}

	return &domain.Todo{
		ID:          td.ID,
		Title:       td.Title,
		Description: td.Description,
	}
}

func (td TodoList) FromDomain(tdms []domain.Todo) TodoList {
	for _, t := range tdms {
		todo := Todo{}
		todo.FromDomain(&t)
		td = append(td, todo)
	}

	return td
}
