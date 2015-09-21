package solver

import (
	"container/heap"
	//	"github.com/Wouterbeets/n-puzzle/board"
)

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].f < pq[j].f
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	node := x.(*Node)
	node.index = n
	*pq = append(*pq, node)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	node.index = -1 // for safety
	*pq = old[0 : n-1]
	return node
}

// update modifies the f and value of an Node in the queue.
func (pq *PriorityQueue) update(node *Node, g, h, f int) {
	node.g = g
	node.h = h
	node.f = f
	heap.Fix(pq, node.index)
}

type Node struct {
	parent *Node
	f      int
	g      int
	h      int
	state  state
	index  int
}

func (n *Node) Copy() *Node {
	cop := &Node{
		parent: n.parent,
		f:      n.f,
		g:      n.g,
		h:      n.h,
		state:  n.state.Copy(),
	}
	return cop
}

type state interface {
	StateString() string
	GoalString() string
	GetH() int
	Copy() state
	GetMoves() []state
}

type Solver struct {
	BoardStates  map[string]Node
	OpenList     *PriorityQueue
	OpenListBool map[string]bool
	ClosedList   map[string]bool
	Goal         string
}

func abs(num int) int {
	if num < 0 {
		num *= -1
	}
	return num
}

func CalcMD(x, y, fx, fy int) int {
	return abs(x-fx) + abs(y-fy)
}
func New(st state) *Solver {
	s := &Solver{
		BoardStates:  make(map[string]Node),
		OpenListBool: make(map[string]bool),
		ClosedList:   make(map[string]bool),
	}
	heap.Init(s.OpenList)
	currentNode := &Node{parent: nil, state: st.Copy()}
	heap.Push(s.OpenList, currentNode)
	s.OpenListBool[currentNode.state.StateString()] = true
	s.Goal = st.GoalString()
	return s
}

func (s *Solver) Solve() {
	//count := 0
	//var start state
	for len(*s.OpenList) > 0 {
		cNode := heap.Pop(s.OpenList).(*Node)

		//	if count == 0 {
		//		start = cNode.state.Copy()
		//	}
		moves := cNode.state.GetMoves()
		for i := 0; i < len(moves); i++ {
			newNode := &Node{
				g:     cNode.g + 1,
				h:     moves[i].GetH(),
				state: moves[i],
			}
			newNode.f = newNode.g - newNode.h
			//check if newnode was already closed in closed list
			//check if in open
			//	yes: check if g is lower. if so, replace
		}
	}
}
