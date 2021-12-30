package store

import (
	"errors"
	"testing"

	"github.com/Vivi3008/todoCLI/domain"
	"github.com/google/uuid"
)

var tdDel = NewTodoStore()

func TestDeleteTodo(t *testing.T) {
	t.Parallel()

	todos := []domain.Todo{
		{
			Id:          uuid.New().String(),
			Description: "Dormir",
			Status:      domain.Pend,
			Priority:    domain.Normal},
		{
			Id:          uuid.New().String(),
			Description: "Comer chocolate",
			Status:      domain.Done,
			Priority:    domain.High,
		},
		{
			Id:          uuid.New().String(),
			Description: "Jogar Game",
			Status:      domain.Pend,
			Priority:    domain.Normal,
		},
	}

	for _, v := range todos {
		_ = tdDel.StoreTodo(v)
	}

	type TestCase struct {
		name string
		arg  string
		err  error
	}

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
			t.Parallel()

			got := tdDel.DeleteTodoId(tt.arg)

			if !errors.Is(got, tt.err) {
				t.Errorf("Expected %s, got %s", tt.err, got)
			}
		})
	}
}
