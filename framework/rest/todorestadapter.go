package rest

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/robsongomes/todo-go-hexagonal/application"
)

type TodoRestAdapter struct {
	Router *chi.Mux
	todoUC application.TodoUseCase
}

func NewTodoRestAdapter(todoUseCase application.TodoUseCase) *TodoRestAdapter {
	router := chi.NewRouter()
	adapter := &TodoRestAdapter{Router: router, todoUC: todoUseCase}

	adapter.Router.Get("/todos", adapter.List)
	adapter.Router.Post("/todos", adapter.Create)
	adapter.Router.Get("/todos/{id}", adapter.Get)

	return adapter
}

func (ta TodoRestAdapter) Create(w http.ResponseWriter, r *http.Request) {
	todo := new(Todo)
	err := json.NewDecoder(r.Body).Decode(todo)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}
	result, err := ta.todoUC.Create(todo.Title, todo.Description)
	if err != nil {
		log.Println(err)
		http.Error(w, "Could not save todo", http.StatusInternalServerError)
		return
	}
	todo.FromDomain(result)
	json.NewEncoder(w).Encode(todo)
}

func (ta TodoRestAdapter) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	result, err := ta.todoUC.Get(id)
	if err != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}
	var todo *Todo = &Todo{}
	todo.FromDomain(result)

	json.NewEncoder(w).Encode(todo)
}

func (ta TodoRestAdapter) List(w http.ResponseWriter, r *http.Request) {
	result, err := ta.todoUC.List()
	if err != nil {
		log.Println(err)
		http.Error(w, "Error listing todos", http.StatusInternalServerError)
		return
	}
	var todos TodoList = TodoList{}
	todos = todos.FromDomain(result)
	json.NewEncoder(w).Encode(todos)
}
