package usecase

import (
	"errors"
	"reflect"
	"testing"

	"github.com/Vivi3008/todoCLI/domain"
)

func TestCreateTodo(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		name       string
		repository domain.Repository
		args       domain.Todo
		want       domain.Todo
		err        error
	}

	testCases := []TestCase{
		{
			name: "Should create and store todo successfull",
			repository: domain.TodoMock{
				OnStoreTodo: func(todo domain.Todo) error {
					return nil
				},
			},
			args: domain.Todo{
				Description: "Make sex",
				Status:      domain.Pend,
				Priority:    domain.High,
			},
			want: domain.Todo{
				Description: "Make sex",
				Status:      domain.Pend,
				Priority:    domain.High,
			},
			err: nil,
		},
		{
			name: "Fail if got store error",
			repository: domain.TodoMock{
				OnStoreTodo: func(todo domain.Todo) error {
					return domain.ErrEmptyValues
				},
			},
			args: domain.Todo{
				Description: "Make sex again",
				Status:      domain.Pend,
				Priority:    domain.High,
			},
			want: domain.Todo{},
			err:  domain.ErrEmptyValues,
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			uc := NewTodoUsecase(tt.repository)

			got, err := uc.CreateTodo(tt.args)

			if !errors.Is(tt.err, err) {
				t.Errorf("Expected %s, got %s", tt.err, err)
			}
			tt.want.Id = got.Id
			tt.want.CreatedAt = got.CreatedAt

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}
		})
	}
}
