package store

import (
	"errors"
	"reflect"
	"testing"

	"github.com/Vivi3008/todoCLI/domain"
)

/* var tdTest []domain.Todo
var tdSt = NewTodoStore()

*/

func TestListTodos(t *testing.T) {
	type TestCase struct {
		name      string
		runBefore bool
		want      []domain.Todo
		err       error
	}
	todos := ListTodosMock()
	testCases := []TestCase{
		{
			name:      "Should list all todos",
			runBefore: true,
			want:      todos,
			err:       nil,
		},
		{
			name:      "Should list empty",
			runBefore: false,
			want:      []domain.Todo{},
			err:       nil,
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			if tt.runBefore {
				RunBeforeTests()
			}
			got, err := todoStore.ListAllTodos()

			if !errors.Is(tt.err, err) {
				t.Errorf("Expeceted %s, got %s", tt.err, err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}
		})
	}
}
