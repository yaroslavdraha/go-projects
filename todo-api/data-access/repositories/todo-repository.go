package repositories

import (
	"database/sql"
	"todo-api/data-access"
	"todo-api/models"
)

type TodoRepository struct{}

var db = data_access.GetInstance()

func (t *TodoRepository) Create(todo models.TodoCreate) (id int, err error) {
	result, err := db.Exec(`INSERT INTO todos (title) VALUES (?)`, todo.Title)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(lastInsertID), nil
}

func (t *TodoRepository) Update() {}

func (t *TodoRepository) Delete() {}

func (t *TodoRepository) List() {}

func (t *TodoRepository) Get(id int) (*models.Todo, error) {
	var todo models.Todo

	err := db.QueryRow(`SELECT * FROM todos WHERE id = ?`, id).Scan(&todo.Id, &todo.Title, &todo.Completed, &todo.CompletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &todo, nil
}
