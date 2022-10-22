package services

type TodoService interface {
	CreateTodo()
	DeleteTodo()
}

type todoService struct {
}

func NewTodoService() *todoService {
	return &todoService{}
}

func (service *todoService) CreateTodo() {}

func (service *todoService) DeleteTodo() {}
