package domain

type Repository interface {
	StoreTodo(todo Todo) error
	ListAllTodos() ([]Todo, error)
	ListTodoById(id string) (Todo, error)
	DeleteTodoId(id string) error
	DeleteAll() error
}
