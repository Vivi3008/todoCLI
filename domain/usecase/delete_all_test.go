package usecase

import (
	"errors"
	"testing"

	"github.com/Vivi3008/todoCLI/domain"
)

var ErrToDeleteFile = errors.New("intern error")

func TestDeleteAll(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		name       string
		repository domain.Repository
		err        error
	}

	testCases := []TestCase{
		{
			name: "Should delete all todos",
			repository: domain.TodoMock{
				OnDeleteAll: func() error {
					return nil
				},
			},
			err: nil,
		},
		{
			name: "Should delete all todos",
			repository: domain.TodoMock{
				OnDeleteAll: func() error {
					return ErrToDeleteFile
				},
			},
			err: ErrToDeleteFile,
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			uc := NewTodoUsecase(tt.repository)
			err := uc.DeleteAllTodos()

			if !errors.Is(tt.err, err) {
				t.Errorf("Expected error %s, got %s", tt.err, err)
			}
		})
	}
}
