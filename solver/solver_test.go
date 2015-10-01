package solver

import (
	"container/heap"
	"github.com/Wouterbeets/n-puzzle/board"
	"testing"
)

func BenchmarkSolve(bench *testing.B) {
	inp := []int{
		6, 3, 0,
		4, 2, 7,
		1, 5, 8,
	}
	b := board.New(3)
	b.Input(inp)
	s := New(b)
	for i := 0; i < bench.N; i++ {
		s.Solve()
	}

}

func TestPq(t *testing.T) {
	var tests = []struct {
		add        []*Node
		wantPop    []*Node
		updatef    int
		updateItem int
	}{
		{
			add: []*Node{
				{
					f: 4,
				},
				{
					f: 2,
				},
				{
					f: 3,
				},
				{
					f: 1,
				},
			},
			wantPop: []*Node{
				{
					f: 0,
				},
				{
					f: 1,
				},
				{
					f: 2,
				},
				{
					f: 3,
				},
			},
			updatef:    0,
			updateItem: 0,
		},
	}

	pq := new(PriorityQueue)
	heap.Init(pq)
	for _, test := range tests {
		got := make([]*Node, 0, len(test.add))
		for _, v := range test.add {
			heap.Push(pq, v)
		}
		pq.update(test.add[test.updateItem], 0, 0, test.updatef, nil)
		for pq.Len() > 0 {
			n := heap.Pop(pq).(*Node)
			got = append(got, n)

		}
		for k, v := range got {
			if test.wantPop[k].f != v.f {
				t.Error(test.wantPop[k].f, v.f)
			}
		}
	}
}
