package repositories

import (
	"final-project-1/httpserver/models"

	"gorm.io/gorm"
)

type TodoRepository interface {
	CreateTodo(todo models.TodoModel) (models.TodoModel, error)
	DeleteTodo(todo models.TodoModel) (models.TodoModel, error)
	UpdateTodo(data models.TodoModel, id int64) (models.TodoModel, error)
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *todoRepository {
	return &todoRepository{db}
}

func (r *todoRepository) CreateTodo(todo models.TodoModel) (models.TodoModel, error) {
	err := r.db.Create(&todo).Error
	if err != nil {
		return todo, err
	}
	return todo, nil
}

func (r *todoRepository) DeleteTodo(todo models.TodoModel) (models.TodoModel, error) {
	err := r.db.Delete(&todo).Error
	if err != nil {
		return todo, err
	}
	return todo, nil
}

func (r *todoRepository) UpdateTodo(data models.TodoModel, id int64) (models.TodoModel, error) {
	err := r.db.Model(&data).Where("todo_id", id).Updates(models.TodoModel{
		Status: true,
	})
	if err != nil {
		return data, err.Error
	}

	return data, nil
}
