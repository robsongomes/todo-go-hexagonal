package framework

import (
	"database/sql"
	"fmt"

	"github.com/robsongomes/todo-go-hexagonal/domain"
)

type todoMysql struct {
	ID          string
	Title       string
	Description string
}

type todoListMysql []todoMysql

type TodoMysqlAdapter struct {
	db *sql.DB
}

func (m *todoMysql) ToDomain() *domain.Todo {
	return &domain.Todo{
		ID:          m.ID,
		Title:       m.Title,
		Description: m.Description,
	}
}
func (m *todoMysql) FromDomain(todo *domain.Todo) {
	if m == nil {
		m = &todoMysql{}
	}

	m.ID = todo.ID
	m.Title = todo.Title
	m.Description = todo.Description
}

func (m todoListMysql) ToDomain() []domain.Todo {
	todos := make([]domain.Todo, len(m))
	for k, td := range m {
		todo := td.ToDomain()
		todos[k] = *todo
	}

	return todos
}

func NewMysqlAdapter(db *sql.DB) *TodoMysqlAdapter {
	return &TodoMysqlAdapter{db: db}
}

func (a *TodoMysqlAdapter) Get(id string) (*domain.Todo, error) {
	var todo todoMysql = todoMysql{}
	sqsS := fmt.Sprintf("SELECT id, title, description FROM todo WHERE id = '%s'", id)

	result := a.db.QueryRow(sqsS)
	if result.Err() != nil {
		return nil, result.Err()
	}

	if err := result.Scan(&todo.ID, &todo.Title, &todo.Description); err != nil {
		return nil, err
	}

	return todo.ToDomain(), nil
}

func (a *TodoMysqlAdapter) List() ([]domain.Todo, error) {
	var todos todoListMysql
	sqsS := "SELECT id, title, description FROM todo"

	result, err := a.db.Query(sqsS)
	if err != nil {
		return nil, err
	}

	if result.Err() != nil {
		return nil, result.Err()
	}

	for result.Next() {
		todo := todoMysql{}

		if err := result.Scan(&todo.ID, &todo.Title, &todo.Description); err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	return todos.ToDomain(), nil
}

func (a *TodoMysqlAdapter) Create(todo *domain.Todo) (*domain.Todo, error) {
	sqlS := "INSERT INTO todo (id, title, description) VALUES (?, ?, ?)"

	_, err := a.db.Exec(sqlS, todo.ID, todo.Title, todo.Description)

	if err != nil {
		return nil, err
	}

	return todo, nil
}
