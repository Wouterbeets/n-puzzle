package solver_3

import (
	"container/list"
	"fmt"
	"github.com/Wouterbeets/n-puzzle/board"
	"github.com/Wouterbeets/n-puzzle/manhattan_3"
	"github.com/Wouterbeets/n-puzzle/missplacedtiles_3"
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
	StateBoard [][]int
}

type Solver struct {
	nbrRow    int
	openList  *list.List
	closeList *list.List
	Heuristic func([][]int, int) int
}

func Get_StateBoard(size int) [][]int {
	nSlice := make([][]int, size)
	for i := 0; i < size; i++ {
		nSlice[i] = make([]int, size)
	}
	return nSlice
}

func (Svr *Solver) Solve_init(b *board.Board, Heur int) {
	Svr.nbrRow = b.Size
	fNode := new(Node)
	fNode.StateBoard = Get_StateBoard(Svr.nbrRow)
	fNode.relative = nil
	fNode.nbrMove = 0
	fNode.status = START
	Svr.openList = list.New()
	Svr.closeList = list.New()
	if Heur == MANHATTAN {
		Svr.Heuristic = manhattan_3.Get_manhattan_dis
	} else if Heur == MANHATTAN_CONFLICT {
		Svr.Heuristic = manhattan_3.Get_manhattan_dis_linear
	} else if Heur == MISSPLACEDTILES {
		Svr.Heuristic = missplacedtiles_3.Get_missplacedTiles
	} else {
		Svr.Heuristic = manhattan_3.Get_manhattan_dis
	}

	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			fNode.StateBoard[i][j] = b.Tiles[(i*b.Size)+j]
		}
	}
	Svr.openList.PushFront(fNode)
}

func Compare_Node(tab1 [][]int, tab2 [][]int, nbrRow int) bool {
	for i := 0; i < nbrRow; i++ {
		for j := 0; j < nbrRow; j++ {
			if tab1[i][j] != tab2[i][j] {
				return false
			}
		}
	}
	return true
}

func (Svr *Solver) CheckcloseListexist(nNode *Node) *list.Element {
	for e := Svr.closeList.Front(); e != nil; e = e.Next() {
		if Compare_Node(e.Value.(*Node).StateBoard, nNode.StateBoard, Svr.nbrRow) == true {
			return e
		}
	}
	return nil
}

func (Svr *Solver) CheckopenListexist(nNode *Node) *list.Element {
	for e := Svr.openList.Front(); e != nil; e = e.Next() {
		if Compare_Node(e.Value.(*Node).StateBoard, nNode.StateBoard, Svr.nbrRow) == true {
			return e
		}
	}
	return nil
}

func (Svr *Solver) CheckSolved() bool {
	check := 1
	e := Svr.closeList.Back()
	if e == nil {
		return false
	}
	node := e.Value.(*Node)
	for i := 0; i < Svr.nbrRow; i++ {
		for j := 0; j < Svr.nbrRow; j++ {
			if node.StateBoard[i][j] != check {
				return false
			}
			if i == Svr.nbrRow-1 && j == Svr.nbrRow-2 {
				check = 0
			} else {
				check++
			}
		}
	}
	return true
}

func (Svr *Solver) AddNode(nNode *Node, nRelative *Node) {
	nNode.relative = nRelative
	nNode.nbrMove = nRelative.nbrMove + 1
	nNode.cout = nNode.nbrMove + Svr.Heuristic(nNode.StateBoard, Svr.nbrRow)

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

func Copy_Tab(dest [][]int, src [][]int, nbrRow int) {
	for i := 0; i < nbrRow; i++ {
		for j := 0; j < nbrRow; j++ {
			dest[i][j] = src[i][j]
		}
	}
}

func (Svr *Solver) Move_top() {
	nRelative := (Svr.closeList.Back()).Value.(*Node)

	if nRelative.status == BOTTOM {
		return
	}
	nNode := new(Node)
	nNode.StateBoard = Get_StateBoard(Svr.nbrRow)
	nNode.status = TOP
	i, j := manhattan_3.Get_positionValue(nRelative.StateBoard, 0, Svr.nbrRow)
	if i == 0 {
		return
	}
	Copy_Tab(nNode.StateBoard, nRelative.StateBoard, Svr.nbrRow)
	nNode.StateBoard[i][j] = nRelative.StateBoard[i-1][j]
	nNode.StateBoard[i-1][j] = 0
	Svr.AddNode(nNode, nRelative)
}

func (Svr *Solver) Move_bot() {
	nRelative := (Svr.closeList.Back()).Value.(*Node)

	if nRelative.status == TOP {
		return
	}
	nNode := new(Node)
	nNode.StateBoard = Get_StateBoard(Svr.nbrRow)

	nNode.status = BOTTOM
	i, j := manhattan_3.Get_positionValue(nRelative.StateBoard, 0, Svr.nbrRow)
	if i == Svr.nbrRow-1 {
		return
	}
	Copy_Tab(nNode.StateBoard, nRelative.StateBoard, Svr.nbrRow)
	nNode.StateBoard[i][j] = nRelative.StateBoard[i+1][j]
	nNode.StateBoard[i+1][j] = 0
	Svr.AddNode(nNode, nRelative)
}

func (Svr *Solver) Move_left() {
	nRelative := (Svr.closeList.Back()).Value.(*Node)

	if nRelative.status == RIGHT {
		return
	}
	nNode := new(Node)
	nNode.StateBoard = Get_StateBoard(Svr.nbrRow)

	nNode.status = LEFT
	i, j := manhattan_3.Get_positionValue(nRelative.StateBoard, 0, Svr.nbrRow)
	if j == 0 {
		return
	}
	Copy_Tab(nNode.StateBoard, nRelative.StateBoard, Svr.nbrRow)
	nNode.StateBoard[i][j] = nRelative.StateBoard[i][j-1]
	nNode.StateBoard[i][j-1] = 0
	Svr.AddNode(nNode, nRelative)
}

func (Svr *Solver) Move_right() {
	nRelative := (Svr.closeList.Back()).Value.(*Node)

	if nRelative.status == LEFT {
		return
	}
	nNode := new(Node)
	nNode.StateBoard = Get_StateBoard(Svr.nbrRow)
	nNode.status = RIGHT
	i, j := manhattan_3.Get_positionValue(nRelative.StateBoard, 0, Svr.nbrRow)
	if j == Svr.nbrRow-1 {
		return
	}
	Copy_Tab(nNode.StateBoard, nRelative.StateBoard, Svr.nbrRow)
	nNode.StateBoard[i][j] = nRelative.StateBoard[i][j+1]
	nNode.StateBoard[i][j+1] = 0
	Svr.AddNode(nNode, nRelative)
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
	fmt.Print("Found with: ")
	fmt.Print(Svr.closeList.Len())
	fmt.Print(" in close list and: ")
	fmt.Print(Svr.openList.Len())
	fmt.Println(" in open list")
	fmt.Println("Result:")
	fmt.Println((Svr.closeList.Back()).Value.(*Node).StateBoard)
}
