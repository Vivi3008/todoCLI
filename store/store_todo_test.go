package store

import (
	"errors"
	"testing"

	"github.com/Vivi3008/todoCLI/domain"
)

func TestStoreTodo(t *testing.T) {
	t.Parallel()

	tdTest := domain.Todo{
		Description: "Dormir",
		Status:      domain.Pend,
		Priority:    domain.Normal,
	}

	type TestCase struct {
		name   string
		fields domain.Todo
		err    error
	}

	var tdStore = NewTodoStore()
	tdField, _ := domain.NewTodo(tdTest)

	testCases := []TestCase{
		{
			name:   "Should store a todo successfull",
			fields: tdField,
			err:    nil,
		},
		{
			name: "Fail if missing id",
			fields: domain.Todo{
				Description: "Dormir",
				Status:      domain.Pend,
				Priority:    domain.Normal,
			},
			err: ErrStore,
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := tdStore.StoreTodo(tt.fields)

			if !errors.Is(err, tt.err) {
				t.Errorf("Expected %s, got %s", tt.err, err)
			}
		})
	}
}
