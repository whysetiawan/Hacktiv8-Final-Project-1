package services

import (
	"final-project-1/httpserver/dto"
	"final-project-1/httpserver/models"
	"final-project-1/httpserver/repositories"
	"time"
)

type TodoService interface {
	CreateTodo(dto dto.CreateTodoDto) (models.TodoModel, error)
	DeleteTodo(id uint) error
}

type todoService struct {
	todoRepository repositories.TodoRepository
}

func NewTodoService(tr repositories.TodoRepository) *todoService {
	return &todoService{
		tr,
	}
}

func (s *todoService) CreateTodo(dto dto.CreateTodoDto) (models.TodoModel, error) {
	todo := models.TodoModel{
		Name:      dto.Name,
		CreatedAt: time.Now(),
	}

	todo, err := s.todoRepository.CreateTodo(todo)
	if err != nil {
		return todo, err
	}
	return todo, nil
}

func (s *todoService) DeleteTodo(id uint) error {
	todo := models.TodoModel{
		TodoId: id,
	}
	todo, err := s.todoRepository.DeleteTodo(todo)

	return err
}
