package board

import (
	"errors"
	"github.com/Wouterbeets/n-puzzle/plog"
	"testing"
)

func TestInput(t *testing.T) {
	var tests = []struct {
		input []int
		want  string
		err   error
	}{
		{
			input: []int{1, 2, 3, 4, 5, 6, 7, 8, 0},
			want:  "\n1 2 3\n4 5 6\n7 8 0\n",
			err:   nil,
		},
		{
			input: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 0},
			want:  "\n0 0 0\n0 0 0\n0 0 0\n",
			err:   errors.New("error"),
		},
		{
			input: []int{0, 1, 2, 3, 4, 5, 6, 7, 8},
			want:  "\n0 1 2\n3 4 5\n6 7 8\n",
			err:   nil,
		},
	}

	for _, test := range tests {
		b := New(3)
		err := b.Input(test.input)
		plog.Activate(true, true, true, true)
		if b.String() != test.want || err != nil && test.err == nil || err != nil && test.err == nil {
			t.Error("input doesn't match value", test.input, "expected", test.want, "\ngot\n", b.String(), "err:", err, "err want", test.err)
		}
	}
}
