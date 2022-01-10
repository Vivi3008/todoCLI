package store

import (
	"errors"
	"testing"
)

func TestDeleteTodo(t *testing.T) {
	type TestCase struct {
		name string
		arg  string
		err  error
	}

	todos := ListTodosMock()

	testCases := []TestCase{
		{
			name: "Should delete a todo successfull",
			arg:  todos[0].Id,
			err:  nil,
		},
		{
			name: "Fail if id doesnt exists",
			arg:  "f1ad3f1",
			err:  ErrIdNotExists,
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			_ = todoStore.StoreTodo(todos[0])

			got := todoStore.DeleteTodoId(tt.arg)

			if !errors.Is(got, tt.err) {
				t.Errorf("Expected %s, got %s", tt.err, got)
			}
		})
	}
}
