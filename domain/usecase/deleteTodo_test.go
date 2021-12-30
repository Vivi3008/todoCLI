package usecase

import (
	"errors"
	"testing"

	"github.com/Vivi3008/todoCLI/domain"
)

func TestDeleteTodo(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		name       string
		repository domain.Repository
		arg        string
		err        error
	}

	testCases := []TestCase{
		{
			name: "Should delete todo by id",
			repository: domain.TodoMock{
				OnDeleteById: func(id string) error {
					return nil
				},
			},
			arg: "1fdas3",
			err: nil,
		},
		{
			name: "Fail if id doesnt exists",
			repository: domain.TodoMock{
				OnDeleteById: func(id string) error {
					return ErrIdNotExists
				},
			},
			arg: "1fdas3",
			err: ErrIdNotExists,
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			uc := NewTodoUsecase(tt.repository)

			err := uc.DeleteTodo(tt.arg)

			if !errors.Is(tt.err, err) {
				t.Errorf("Expected %s, got %s", tt.err, err)
			}
		})
	}
}
