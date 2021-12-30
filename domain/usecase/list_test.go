package usecase

import (
	"reflect"
	"testing"

	"github.com/Vivi3008/todoCLI/domain"
	"github.com/google/uuid"
)

var TdTest = []domain.Todo{
	{
		Id:          uuid.New().String(),
		Description: "Dormir",
		Status:      domain.Pend,
		Priority:    domain.Normal,
	},
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

func TestListAll(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		name       string
		repository domain.Repository
		want       []domain.Todo
	}

	testCases := []TestCase{
		{
			name: "Should list all todos",
			repository: domain.TodoMock{
				OnListAll: func() []domain.Todo {
					return TdTest
				},
			},
			want: TdTest,
		},
		{
			name: "Should list empty",
			repository: domain.TodoMock{
				OnListAll: func() []domain.Todo {
					return []domain.Todo{}
				},
			},
			want: []domain.Todo{},
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			uc := NewTodoUsecase(tt.repository)

			got := uc.ListAll()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}
		})
	}
}
