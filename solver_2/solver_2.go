package solver_2

import (
	"container/list"
	"fmt"
	"github.com/Wouterbeets/n-puzzle/board"
	"github.com/Wouterbeets/n-puzzle/manhattan"
	"github.com/Wouterbeets/n-puzzle/missplacedtiles"
)

const (
	START              = 0
	TOP                = 1
	BOTTOM             = 2
	LEFT               = 3
	RIGHT              = 4
	MANHATTAN          = 1
	MANHATTAN_CONFLICT = 2
	MISSPLACEDTILES    = 3
)

type Node struct {
	relative   *Node
	nbrMove    int
	cout       int
	status     int
	StateBoard []int
}

type Solver struct {
	size      int
	nbrRow    int
	openList  *list.List
	closeList *list.List
	Heuristic func([]int, int, int) int
}

func (Svr *Solver) Solve_init(b *board.Board, Heur int) {
	Svr.size = b.Size * b.Size
	Svr.nbrRow = b.Size
	fNode := new(Node)
	fNode.StateBoard = make([]int, Svr.size)
	fNode.relative = nil
	fNode.nbrMove = 0
	fNode.status = START
	Svr.openList = list.New()
	Svr.closeList = list.New()
	if Heur == MANHATTAN {
		Svr.Heuristic = manhattan.Get_manhattan_dis
	} else if Heur == MANHATTAN_CONFLICT {
		Svr.Heuristic = manhattan.Get_manhattan_dis_linear
	} else if Heur == MISSPLACEDTILES {
		Svr.Heuristic = missplacedtiles.Get_missplacedTiles
	} else {
		Svr.Heuristic = manhattan.Get_manhattan_dis
	}
	for index := 0; index < Svr.size; index++ {
		fNode.StateBoard[index] = b.Tiles[index]
	}
	Svr.openList.PushFront(fNode)
}

func (Svr *Solver) CheckcloseListexist(nNode *Node) *list.Element {
	for e := Svr.closeList.Front(); e != nil; e = e.Next() {
		for i := 0; i < Svr.size; i++ {
			if e.Value.(*Node).StateBoard[i] != nNode.StateBoard[i] {
				break
			}
			if i == Svr.size-1 {
				return e
			}
		}
	}
	return nil
}

func (Svr *Solver) CheckopenListexist(nNode *Node) *list.Element {
	for e := Svr.openList.Front(); e != nil; e = e.Next() {
		for i := 0; i < Svr.size; i++ {
			if e.Value.(*Node).StateBoard[i] != nNode.StateBoard[i] {
				break
			}
			if i == Svr.size-1 {
				return e
			}
		}
	}
	return nil
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

func (Svr *Solver) AddNode(nNode *Node, nRelative *Node) {
	nNode.relative = nRelative
	nNode.nbrMove = nRelative.nbrMove + 1
	nNode.cout = nNode.nbrMove + Svr.Heuristic(nNode.StateBoard, Svr.nbrRow, Svr.size)
	eOpen := Svr.CheckopenListexist(nNode)
	eClose := Svr.CheckcloseListexist(nNode)
	if eClose != nil {
		return
	} else if eOpen != nil && eOpen.Value.(*Node).cout <= nNode.cout {
		return
	} else if eOpen != nil {
		Svr.openList.Remove(eOpen)
	}
	e := Svr.openList.Front()
	for e != nil {
		if nNode.cout < e.Value.(*Node).cout {
			Svr.openList.InsertBefore(nNode, e)
			return
		}
		e = e.Next()
	}
	Svr.openList.PushBack(nNode)
}

func (Svr *Solver) Move_top() {
	nRelative := (Svr.closeList.Back()).Value.(*Node)

	if nRelative.status == BOTTOM {
		return
	}
	i := 0
	nNode := new(Node)
	nNode.StateBoard = make([]int, Svr.size)

	nNode.status = TOP
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
	Svr.AddNode(nNode, nRelative)
}

func (Svr *Solver) Move_bot() {
	nRelative := (Svr.closeList.Back()).Value.(*Node)

	if nRelative.status == TOP {
		return
	}
	i := 0
	nNode := new(Node)
	nNode.StateBoard = make([]int, Svr.size)

	nNode.status = BOTTOM
	for nRelative.StateBoard[i] != 0 {
		i++
	}
	row := i / Svr.nbrRow
	if row == Svr.nbrRow-1 {
		return
	}
	copy(nNode.StateBoard, nRelative.StateBoard)
	nNode.StateBoard[i] = nNode.StateBoard[i+Svr.nbrRow]
	nNode.StateBoard[i+Svr.nbrRow] = 0
	Svr.AddNode(nNode, nRelative)
}

func (Svr *Solver) Move_left() {
	nRelative := (Svr.closeList.Back()).Value.(*Node)

	if nRelative.status == RIGHT {
		return
	}
	i := 0
	nNode := new(Node)
	nNode.StateBoard = make([]int, Svr.size)

	nNode.status = LEFT
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
	Svr.AddNode(nNode, nRelative)
}

func (Svr *Solver) Move_right() {
	nRelative := (Svr.closeList.Back()).Value.(*Node)

	if nRelative.status == LEFT {
		return
	}
	i := 0
	nNode := new(Node)
	nNode.StateBoard = make([]int, Svr.size)
	nNode.status = RIGHT
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
	Svr.AddNode(nNode, nRelative)
}

func (Svr *Solver) PrintResult() {
	e := Svr.closeList.Back()
	nRelative := e.Value.(*Node)
	fmt.Printf("Result found in %d move\n", nRelative.nbrMove)
	for i := nRelative.nbrMove; nRelative != nil; i-- {
		fmt.Printf("Move: %d, with heuristic value: %d\n", i, nRelative.cout)
		fmt.Println(nRelative.StateBoard)
		nRelative = nRelative.relative
	}
}

func (Svr *Solver) Solve() {
	for Svr.CheckSolved() == false {
		e := Svr.openList.Front()
		if e == nil {
			return
		} else {
			Svr.closeList.PushBack(e.Value.(*Node))
			Svr.openList.Remove(e)
			Svr.Move_top()
			Svr.Move_bot()
			Svr.Move_left()
			Svr.Move_right()
		}
	}

	valClosed := Svr.openList.Len()
	valOpen := Svr.closeList.Len()
	fmt.Printf("Result Found !\nOpen list contains %d elements.\nClose list contains %d elements (complexitiy in time)\n", valOpen, valClosed)
	value := valOpen + valClosed
	fmt.Printf("Complexity in size: %d\n", value)
	Svr.PrintResult()
}
