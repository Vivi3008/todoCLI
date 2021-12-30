package usecase

import (
	"errors"
	"reflect"
	"testing"

	"github.com/Vivi3008/todoCLI/domain"
)

func TestUpdateTodo(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		name       string
		repository domain.Repository
		arg        string
		want       domain.Todo
		err        error
	}

	testCases := []TestCase{
		{
			name: "Should mark an existing todo as done",
			repository: domain.TodoMock{
				OnListById: func(id string) (domain.Todo, error) {
					return TdTest[0], nil
				},
				OnStoreTodo: func(todo domain.Todo) error {
					return nil
				},
			},
			arg: TdTest[0].Id,
			want: domain.Todo{
				Id:          TdTest[0].Id,
				Description: TdTest[0].Description,
				Status:      domain.Done,
				Priority:    TdTest[0].Priority,
				CreatedAt:   TdTest[0].CreatedAt,
			},
			err: nil,
		},
		{
			name: "Should mark an existing todo as pending",
			repository: domain.TodoMock{
				OnListById: func(id string) (domain.Todo, error) {
					return domain.Todo{
						Id:          TdTest[0].Id,
						Description: TdTest[0].Description,
						Status:      domain.Done,
						Priority:    TdTest[0].Priority,
						CreatedAt:   TdTest[0].CreatedAt,
					}, nil
				},
				OnStoreTodo: func(todo domain.Todo) error {
					return nil
				},
			},
			arg: TdTest[0].Id,
			want: domain.Todo{
				Id:          TdTest[0].Id,
				Description: TdTest[0].Description,
				Status:      domain.Pend,
				Priority:    TdTest[0].Priority,
				CreatedAt:   TdTest[0].CreatedAt,
			},
			err: nil,
		},
		{
			name: "Fail if id doens't exists",
			repository: domain.TodoMock{
				OnListById: func(id string) (domain.Todo, error) {
					return domain.Todo{}, ErrIdNotExists
				},
				OnStoreTodo: func(todo domain.Todo) error {
					return nil
				},
			},
			arg:  TdTest[0].Id,
			want: domain.Todo{},
			err:  ErrIdNotExists,
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			uc := NewTodoUsecase(tt.repository)

			got, err := uc.UpdateTodo(tt.arg)

			if !errors.Is(err, tt.err) {
				t.Errorf("Expected %s, got %s", tt.err, err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}
		})
	}
}
