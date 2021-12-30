package domain

type Repository interface {
	StoreTodo(todo Todo) error
	ListAllTodos() []Todo
	ListTodoById(id string) (Todo, error)
}
