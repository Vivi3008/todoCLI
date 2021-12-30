package domain

import (
	"errors"
	"reflect"
	"testing"
)

func TestCreateTodo(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		name string
		args Todo
		want Todo
		err  error
	}

	testCases := []TestCase{
		{
			name: "Should create a todo succesffull",
			args: Todo{
				Description: "Fazer café",
				Status:      Pend,
				Priority:    High,
			},
			want: Todo{
				Description: "Fazer café",
				Status:      Pend,
				Priority:    High,
			},
			err: nil,
		},
		{
			name: "Fail if miss some values",
			args: Todo{
				Status:   Pend,
				Priority: High,
			},
			want: Todo{},
			err:  ErrEmptyValues,
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := NewTodo(tt.args)

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
