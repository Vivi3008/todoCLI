package usecase

import (
	"errors"
	"reflect"
	"testing"

	"github.com/Vivi3008/todoCLI/domain"
	"github.com/google/uuid"
)

var tdTestId = []domain.Todo{
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

func TestListOneTodo(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		name       string
		arg        string
		repository domain.Repository
		want       domain.Todo
		err        error
	}

	testCases := []TestCase{
		{
			name: "Should list todo by id",
			repository: domain.TodoMock{
				OnListById: func(id string) (domain.Todo, error) {
					return tdTestId[0], nil
				},
				OnListAll: func() ([]domain.Todo, error) {
					return tdTestId, nil
				},
			},
			arg:  tdTestId[0].Id,
			want: tdTestId[0],
			err:  nil,
		},
		{
			name: "Fail if id doesnt exists",
			arg:  "fa1df3sa",
			repository: domain.TodoMock{
				OnListById: func(id string) (domain.Todo, error) {
					if id == tdTestId[0].Id {
						return tdTestId[0], nil
					}
					return domain.Todo{}, domain.ErrEmptyValues
				},
				OnListAll: func() ([]domain.Todo, error) {
					return tdTestId, nil
				},
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

			got, err := uc.ListTodoId(tt.arg)

			if !errors.Is(err, tt.err) {
				t.Errorf("Expected %s, got %s", tt.err, err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Expected %s, got %s", tt.want, got)
			}
		})
	}
}
