package main

import (
	"net/http"

	"github.com/robsongomes/todo-go-hexagonal/application"
	"github.com/robsongomes/todo-go-hexagonal/framework"
	"github.com/robsongomes/todo-go-hexagonal/framework/rest"
)

func main() {
	// var db *sql.DB
	var todoUseCase application.TodoUseCase

	// todoOutput = framework.NewMysqlAdapter(db)
	todoOutput := framework.NewMemoryAdapter()
	todoUseCase = application.NewTodoInputPort(todoOutput)

	_, err := todoUseCase.Create("Buy cheese", "Buy some mussarela")
	if err != nil {
		panic(err)
	}
	adapter := rest.NewTodoRestAdapter(todoUseCase)

	http.ListenAndServe(":3000", adapter.Router)
}
