package dto

type CreateTodoDto struct {
	Name string `json:"name" binding:"required"`
}
