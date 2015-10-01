package solver

import (
	"container/heap"
	"fmt"
	"github.com/Wouterbeets/n-puzzle/board"
	"os"
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
	b *board.Board
	n *Node
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
	b      *board.Board
	open   bool
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

type Solver struct {
	BoardStates map[string]*Node
	OpenList    *PriorityQueue
	Goal        string
}

func (s *Solver) checkSolved(cNode *Node) {
	if cNode.key == s.Goal {
		fmt.Println("solition reached")
		fmt.Println(s.OpenList.Len())
		for cNode.parent != nil {
			//fmt.Println(s.BoardStates[cNode.key].b)
			cNode = cNode.parent
		}
		//fmt.Println(s.BoardStates[cNode.key].b)
		os.Exit(1)
	}
}

func (s *Solver) treatCurrentNode(cNode *Node) {
	s.checkSolved(cNode)
	cNode.open = false
}

func (s *Solver) getMoves(cNode *Node) []*board.Board {
	return s.BoardStates[cNode.key].b.GetMoves()
}

func makeNodeFromState(b *board.Board, parentNode *Node) *Node {
	newNode := &Node{
		g:      parentNode.g + 1,
		h:      b.GetH(),
		key:    b.StateString(),
		parent: parentNode,
	}
	newNode.f = newNode.g + newNode.h
	return newNode
}

func New(b *board.Board) *Solver {
	s := &Solver{
		BoardStates: make(map[string]*Node),
		OpenList:    new(PriorityQueue),
	}
	heap.Init(s.OpenList)
	currentNode := &Node{
		parent: nil,
		key:    b.StateString(),
		h:      b.GetH(),
		g:      0,
		b:      b,
		open:   true,
	}
	currentNode.f = currentNode.g + currentNode.h
	s.BoardStates[currentNode.key] = currentNode
	heap.Push(s.OpenList, currentNode)
	s.Goal = b.GoalString()
	return s
}

func (s *Solver) Solve() {
	for len(*s.OpenList) > 0 {
		cNode := heap.Pop(s.OpenList).(*Node)
		s.treatCurrentNode(cNode)
		moves := s.getMoves(cNode)
		for i := 0; i < len(moves); i++ {
			newNode := makeNodeFromState(moves[i], cNode)
			if v, has := s.BoardStates[newNode.key]; has == false || v.open == true {
				if has == false {
					v.open = true
					heap.Push(s.OpenList, newNode)
				} else if v.g > newNode.g {
					s.OpenList.update(v, newNode.g, newNode.h, newNode.f, newNode.parent)
				}
			}
			s.BoardStates[newNode.key] = newNode
		}
	}
}
