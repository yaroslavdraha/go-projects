package models

type Todo struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Completed   bool   `json:"completed"`
	CompletedAt string `json:"completed_at"`
}

type TodoCreate struct {
	Title string `json:"title" validate:"required"`
}
