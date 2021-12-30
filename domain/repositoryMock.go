package domain

type TodoMock struct {
	OnStoreTodo func(todo Todo) error
	OnListAll   func() []Todo
	OnListById  func(id string) (Todo, error)
}

func (m TodoMock) ListAllTodos() []Todo {
	return m.OnListAll()
}

func (m TodoMock) ListTodoById(id string) (Todo, error) {
	return m.OnListById(id)
}

func (m TodoMock) StoreTodo(todo Todo) error {
	return m.OnStoreTodo(todo)
}
