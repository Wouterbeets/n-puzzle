package solver

import (
	"container/heap"
	"github.com/Wouterbeets/n-puzzle/plog"
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

type StateNode struct {
	st State
	n  *Node
}

// update modifies the f and value of an Node in the queue.
func (pq *PriorityQueue) update(node *Node, g, h, f int, p *Node) {
	node.g = g
	node.h = h
	node.f = f
	node.parent = p
	heap.Fix(pq, node.index)
}

type Node struct {
	parent *Node
	f      int
	g      int
	h      int
	key    string
	index  int
}

func (n *Node) Copy() *Node {
	cop := &Node{
		parent: n.parent,
		f:      n.f,
		g:      n.g,
		h:      n.h,
		key:    n.key,
	}
	return cop
}

type State interface {
	StateString() string
	GoalString() string
	GetH() int
	GetMoves() []State
}

type Solver struct {
	BoardStates  map[string]StateNode
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
func New(st State) *Solver {
	s := &Solver{
		BoardStates:  make(map[string]StateNode),
		OpenListBool: make(map[string]bool),
		ClosedList:   make(map[string]bool),
	}
	heap.Init(s.OpenList)
	currentNode := &Node{
		parent: nil,
		key:    st.StateString(),
		h:      st.GetH(),
		g:      0,
	}
	currentNode.f = currentNode.g + currentNode.h
	s.BoardStates[st.StateString()] = StateNode{st: st, n: currentNode}
	heap.Push(s.OpenList, currentNode)
	s.OpenListBool[currentNode.key] = true
	s.Goal = st.GoalString()
	return s
}

func (s *Solver) Solve() {
	//count := 0
	//var start State
	for len(*s.OpenList) > 0 {
		cNode := heap.Pop(s.OpenList).(*Node)
		//	if count == 0 {
		//		start = cNode.State.Copy()
		//	}
		if cNode.key == s.Goal {
			plog.Info.Println("solition reached")
		}
		//add to closed list and remove from open list
		s.OpenListBool[cNode.key] = false
		s.ClosedList[cNode.key] = true

		moves := s.BoardStates[cNode.key].st.GetMoves()
		for i := 0; i < len(moves); i++ {
			newNode := &Node{
				g:      cNode.g + 1,
				h:      moves[i].GetH(),
				key:    moves[i].StateString(),
				parent: cNode,
			}
			newNode.f = newNode.g + newNode.h
			s.BoardStates[newNode.key] = StateNode{
				st: moves[i],
				n:  newNode,
			}
			if _, ok := s.ClosedList[newNode.key]; !ok {

				if _, ok := s.OpenListBool[newNode.key]; !ok {
					s.OpenListBool[newNode.key] = true
					heap.Push(s.OpenList, newNode)
				}
				if _, ok := s.OpenListBool[newNode.key]; ok && s.BoardStates[newNode.key].n.g > newNode.g {
					s.OpenList.update(s.BoardStates[newNode.key].n, newNode.g, newNode.h, newNode.f, newNode.parent)
				}
			}
		}
	}
}
