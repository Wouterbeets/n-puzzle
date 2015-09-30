package solver_2

import (
	"container/list"
	"fmt"
	"github.com/Wouterbeets/n-puzzle/board"
	"github.com/Wouterbeets/n-puzzle/manhattan"
)

type Node struct {
	relative   *Node
	cout       int
	StateBoard []int
}

type Solver struct {
	size      int
	nbrRow    int
	openList  *list.List
	closeList *list.List
}

func (Svr *Solver) Solve_init(b *board.Board) {
	index := 0
	Svr.size = b.Size * b.Size
	Svr.nbrRow = b.Size
	fNode := new(Node)
	fNode.StateBoard = make([]int, Svr.size)
	fNode.relative = nil
	Svr.openList = list.New()
	Svr.closeList = list.New()

	index = 0
	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			fNode.StateBoard[index] = b.Rows[i][j].Val
			index++
		}
	}
	fNode.cout = manhattan.Get_manhattan_dis(fNode.StateBoard, Svr.nbrRow, Svr.size)
	Svr.openList.PushFront(fNode)
}

func (Svr *Solver) CheckSolved() bool {
	max := Svr.size - 1
	e := Svr.closeList.Back()
	if e == nil {
		return false
	}
	node := e.Value.(*Node)
	for i := 0; i < max; i++ {
		if node.StateBoard[i] != i+1 {
			return false
		}
	}
	return true
}

func (Svr *Solver) AddNode(nNode *Node) {
	fmt.Println("COUCOU")
	nNode.cout = manhattan.Get_manhattan_dis(nNode.StateBoard, Svr.nbrRow, Svr.size)
	eNode := Svr.openList.PushFront(nNode)
	e := eNode.Next()
	for e != nil {
		if nNode.cout < e.Value.(*Node).cout {
			fmt.Println("Add:D")
			Svr.openList.MoveBefore(eNode, e)
			return
		}
		e = e.Next()
	}
}

func (Svr *Solver) Move_top() {
	nNode := new(Node)
	nNode.StateBoard = make([]int, Svr.size)
	nRelative := (Svr.closeList.Back()).Value.(*Node)
	i := 0

	for nRelative.StateBoard[i] != 0 {
		i++
	}
	row := i / Svr.nbrRow
	if row == 0 {
		return
	}
	copy(nNode.StateBoard, nRelative.StateBoard)
	nNode.StateBoard[i] = nNode.StateBoard[i-Svr.nbrRow]
	nNode.StateBoard[i-Svr.nbrRow] = 0
	nNode.relative = nRelative
	Svr.AddNode(nNode)
}

func (Svr *Solver) Move_bot() {
	nNode := new(Node)
	nNode.StateBoard = make([]int, Svr.size)
	nRelative := (Svr.closeList.Back()).Value.(*Node)
	i := 0

	for nRelative.StateBoard[i] != 0 {
		i++
	}
	row := i / Svr.nbrRow
	if row == Svr.nbrRow-1 {
		return
	}
	copy(nNode.StateBoard, nRelative.StateBoard)
	fmt.Print("Value de i: ")
	fmt.Println(i)
	fmt.Print("Value de nbrRow: ")
	fmt.Println(Svr.nbrRow)
	nNode.StateBoard[i] = nNode.StateBoard[i+Svr.nbrRow]
	nNode.StateBoard[i+Svr.nbrRow] = 0
	nNode.relative = nRelative
	Svr.AddNode(nNode)
}

func (Svr *Solver) Move_left() {
	nNode := new(Node)
	nNode.StateBoard = make([]int, Svr.size)
	nRelative := (Svr.closeList.Back()).Value.(*Node)
	i := 0

	for nRelative.StateBoard[i] != 0 {
		i++
	}
	row := i % Svr.nbrRow
	if row == 0 {
		return
	}
	copy(nNode.StateBoard, nRelative.StateBoard)
	nNode.StateBoard[i] = nNode.StateBoard[i-1]
	nNode.StateBoard[i-1] = 0
	nNode.relative = nRelative
	Svr.AddNode(nNode)
}

func (Svr *Solver) Move_right() {
	nNode := new(Node)
	nNode.StateBoard = make([]int, Svr.size)
	nRelative := (Svr.closeList.Back()).Value.(*Node)
	i := 0

	for nRelative.StateBoard[i] != 0 {
		i++
	}
	row := (i + 1) % Svr.nbrRow
	if row == 0 {
		return
	}
	copy(nNode.StateBoard, nRelative.StateBoard)
	nNode.StateBoard[i] = nNode.StateBoard[i+1]
	nNode.StateBoard[i+1] = 0
	nNode.relative = nRelative
	Svr.AddNode(nNode)
}

func (Svr *Solver) Solve() {
	for Svr.CheckSolved() == false {
		e := Svr.openList.Front()
		Svr.closeList.PushBack(e.Value.(*Node))
		Svr.openList.Remove(e)
		Svr.Move_top()
		Svr.Move_bot()
		Svr.Move_left()
		Svr.Move_right()
	}
	fmt.Print("Found with: ")
	fmt.Println(Svr.closeList.Len())
}
