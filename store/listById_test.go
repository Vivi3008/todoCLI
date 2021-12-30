package store

import (
	"errors"
	"reflect"
	"testing"

	"github.com/Vivi3008/todoCLI/domain"
	"github.com/google/uuid"
)

var tdStore = NewTodoStore()

func TestListTodoById(t *testing.T) {
	t.Parallel()

	tdTestId := []domain.Todo{
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

	for _, v := range tdTestId {
		_ = tdStore.StoreTodo(v)
	}

	type TestCase struct {
		name string
		arg  string
		want domain.Todo
		err  error
	}

	testCases := []TestCase{
		{
			name: "Should list todo by id",
			arg:  tdTestId[0].Id,
			want: tdTestId[0],
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
			t.Parallel()

			got, err := tdStore.ListTodoById(tt.arg)

			if !errors.Is(err, tt.err) {
				t.Errorf("Expected %s, got %s", tt.err, err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}
		})
	}

}
