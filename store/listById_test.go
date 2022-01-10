package store

import (
	"errors"
	"reflect"
	"testing"

	"github.com/Vivi3008/todoCLI/domain"
)

func TestListTodoById(t *testing.T) {
	todos := ListTodosMock()

	type TestCase struct {
		name string
		arg  string
		want domain.Todo
		err  error
	}

	testCases := []TestCase{
		{
			name: "Should list todo by id",
			arg:  todos[0].Id,
			want: todos[0],
			err:  nil,
		},
		{
			name: "Fail to list todo if id doesn't exists",
			arg:  "1fdsaf5",
			want: domain.Todo{},
			err:  ErrIdNotExists,
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			_ = todoStore.StoreTodo(todos[0])
			got, err := todoStore.ListTodoById(tt.arg)

			if !errors.Is(err, tt.err) {
				t.Errorf("Expected %s, got %s", tt.err, err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}
		})
	}

}
