package main

import (
	"log"
	"net/http"
	"todo-api/data-access"
	"todo-api/handlers"
)

var todoHandlers = handlers.TodoHandlers{}

func main() {
	data_access.InitDB()

	mux := http.NewServeMux()

	defineApiRoutes(mux)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
