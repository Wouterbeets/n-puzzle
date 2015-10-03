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
		}, {
			input: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 0},
			want:  "\n0 0 0\n0 0 0\n0 0 0\n",
			err:   errors.New("error"),
		}, {
			input: []int{},
			want:  "\n0 0 0\n0 0 0\n0 0 0\n",
			err:   errors.New("error"),
		}, {
			input: []int{0, 1, 2, 3, 4, 5, 6, 7, 8},
			want:  "\n0 1 2\n3 4 5\n6 7 8\n",
			err:   nil,
		},
	}

	for _, test := range tests {
		b := New(3)
		err := b.Input(test.input)
		plog.Activate(false, false, false, false)
		if b.String() != test.want || err != nil && test.err == nil || err != nil && test.err == nil {
			t.Error("input doesn't match value", test.input, "expected", test.want, "\ngot\n", b.String(), "err:", err, "err want", test.err)
		}
	}
}

func TestCheckNumbers(t *testing.T) {
	var tests = []struct {
		size  int
		input []int
		err   error
	}{
		{
			size:  3,
			input: []int{1, 0, 3, 4, 5, 6, 7, 2, 0},
			err:   errors.New("error"),
		},
		{
			size:  3,
			input: []int{1, 2, 3, 4, 5, 6, 7, 8, 0},
			err:   nil,
		},
	}
	for _, test := range tests {
		b := New(test.size)
		err := b.checkNumbers(test.input)
		if err == nil && test.err != nil || err != nil && test.err == nil {
			t.Error("input produces wrong error return", test.input, "expected", test.err, "\ngot\n", err)
		}
	}
}

func TestMove(t *testing.T) {
	var tests = []struct {
		size  int
		input []int
		want  string
		move  int
		err   error
	}{
		{
			size:  3,
			input: []int{1, 2, 3, 4, 5, 6, 7, 8, 0},
			want:  "\n1 2 3\n4 5 6\n7 8 0\n",
			move:  Up,
			err:   errors.New("error"),
		}, {
			size:  3,
			input: []int{1, 2, 3, 4, 5, 0, 7, 8, 6},
			want:  "\n1 2 3\n4 0 5\n7 8 6\n",
			move:  Right,
			err:   nil,
		}, {
			size:  3,
			input: []int{1, 2, 3, 0, 5, 6, 7, 8, 4},
			want:  "\n1 2 3\n5 0 6\n7 8 4\n",
			move:  Left,
			err:   nil,
		}, {
			size:  3,
			input: []int{0, 2, 3, 4, 5, 6, 7, 8, 1},
			want:  "\n4 2 3\n0 5 6\n7 8 1\n",
			move:  Up,
			err:   nil,
		}, {
			size:  3,
			input: []int{1, 2, 3, 4, 5, 6, 7, 8, 0},
			want:  "\n1 2 3\n4 5 0\n7 8 6\n",
			move:  Down,
			err:   nil,
		},
	}
	for _, test := range tests {
		b := New(test.size)
		b.Input(test.input)
		err := b.Move(test.move)
		if b.String() != test.want || err != nil && test.err == nil || err != nil && test.err == nil {
			t.Error("Move fail", test.input, "expected", test.want, "\ngot\n", b.String(), "err:", err, "errWant", test.err, "move", test.move, "\nboard.Tiles =", b.Tiles)
		}
	}

}
func TestOutOfPlace(t *testing.T) {
	var tests = []struct {
		size  int
		input []int
		want  int
	}{
		{
			size: 3,
			input: []int{
				1, 2, 3,
				4, 5, 6,
				7, 0, 8,
			},
			want: 2,
		},
		{
			size: 3,
			input: []int{
				2, 3, 1,
				4, 5, 6,
				7, 0, 8,
			},
			want: 5,
		},
		{
			size: 4,
			input: []int{
				5, 2, 3, 0,
				1, 6, 7, 8,
				9, 10, 11, 12,
				13, 14, 15, 4,
			},
			want: 4,
		},
	}
	for _, test := range tests {
		b := New(test.size)
		b.HeurFun = OutOfPlace
		b.Input(test.input)
		result := b.GetH()
		if result != test.want {
			t.Error("heuristic function returns wrong h value", result, test.want)
		}
	}
}

func TestMD(t *testing.T) {
	var tests = []struct {
		size  int
		input []int
		want  int
	}{
		{
			size: 3,
			input: []int{
				1, 2, 3,
				4, 5, 6,
				7, 0, 8,
			},
			want: 2,
		},
		{
			size: 3,
			input: []int{
				2, 3, 1,
				4, 5, 6,
				7, 0, 8,
			},
			want: 6,
		},
		{
			size: 4,
			input: []int{
				5, 2, 3, 0,
				1, 6, 7, 8,
				9, 10, 11, 12,
				13, 14, 15, 4,
			},
			want: 8,
		},
	}
	for _, test := range tests {
		b := New(test.size)
		b.HeurFun = CalcMD
		b.Input(test.input)
		result := b.GetH()
		if result != test.want {
			t.Error("heuristic function returns wrong h value", result, test.want)
		}
	}
}

func TestGoalStateString(t *testing.T) {
	var tests = []struct {
		size      int
		input     []int
		wantGoal  string
		wantState string
	}{
		{
			size:      3,
			input:     []int{1, 2, 4, 3, 5, 6, 7, 8, 0},
			wantState: "3 1 2 4 3 5 6 7 8 0 ",
			wantGoal:  "3 1 2 3 4 5 6 7 8 0 ",
		},
		{
			size:      4,
			input:     []int{1, 2, 4, 3, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 0},
			wantState: "4 1 2 4 3 5 6 7 8 9 10 11 12 13 14 15 0 ",
			wantGoal:  "4 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 0 ",
		},
	}
	for _, test := range tests {
		b := New(test.size)
		b.Input(test.input)
		if b.StateString() != test.wantState || b.GoalString() != test.wantGoal {
			t.Error("state =", b.StateString(), "wanted", test.wantState, "goal =", b.GoalString(), "want", test.wantGoal)

		}
	}

}
