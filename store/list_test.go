package store

import (
	"errors"
	"reflect"
	"testing"

	"github.com/Vivi3008/todoCLI/domain"
	"github.com/google/uuid"
)

var tdTest []domain.Todo
var tdSt = NewTodoStore()

func runBefore() {
	tdTest = []domain.Todo{
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

	for _, v := range tdTest {
		_ = tdSt.StoreTodo(v)
	}
}

func TestListTodos(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		name      string
		runBefore bool
		want      []domain.Todo
		err       error
	}

	testCases := []TestCase{
		{
			name:      "Should list all todos",
			runBefore: true,
			want:      tdTest,
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
			t.Parallel()

			if tt.runBefore {
				runBefore()
			}
			got, err := tdSt.ListAllTodos()

			if !errors.Is(tt.err, err) {
				t.Errorf("Expeceted %s, got %s", tt.err, err)
			}
			tt.want = got

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}
		})
	}
}
