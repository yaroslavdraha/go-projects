package main

import "net/http"

func defineApiRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /todos", todoHandlers.Create)
	mux.HandleFunc("GET /todos", todoHandlers.List)
	mux.HandleFunc("GET /todos/{id}", todoHandlers.Get)
	mux.HandleFunc("DELETE /todos/{id}", todoHandlers.Delete)
}
