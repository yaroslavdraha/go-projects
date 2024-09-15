package handlers

import (
	"net/http"
	"strconv"
	"todo-api/data-access/repositories"
	"todo-api/libs/api"
	"todo-api/models"
)

var todoRepo = repositories.TodoRepository{}

type TodoHandlers struct{}

func (t *TodoHandlers) Create(w http.ResponseWriter, r *http.Request) {
	var body models.TodoCreate

	if err := api.ParseAndValidate(r, &body); err != nil {
		api.RespondWithError(w, http.StatusBadRequest, err.Error())
	}

	id, err := todoRepo.Create(body)
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	todo, err := todoRepo.Get(id)
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	api.RespondWithJSON(w, http.StatusCreated, todo)
}

func (t *TodoHandlers) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		api.RespondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	todo, err := todoRepo.Get(id)
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	status := http.StatusOK
	if todo == nil {
		status = http.StatusNotFound
	}

	api.RespondWithJSON(w, status, todo)
}

func (t *TodoHandlers) Update(w http.ResponseWriter, r *http.Request) {}

func (t *TodoHandlers) Delete(w http.ResponseWriter, r *http.Request) {}

func (t *TodoHandlers) List(w http.ResponseWriter, r *http.Request) {}
