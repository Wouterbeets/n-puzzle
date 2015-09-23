package board

import (
	"errors"
	"github.com/Wouterbeets/n-puzzle/plog"
	"github.com/Wouterbeets/n-puzzle/solver"
	"strconv"
	"strings"
)

const (
	Up = iota
	Down
	Left
	Right
	MAX_SIZE = 42
)

type Tile struct {
	Val int
}

func (t Tile) String() string {
	return strconv.Itoa(t.Val)
}

type Row []Tile

func (r Row) Copy() Row {
	ret := make(Row, len(r), len(r))
	copy(ret, r)
	return ret
}

type Rows []Row

func (r Rows) Copy() Rows {
	ret := make(Rows, len(r), len(r))
	copy(ret, r)
	return ret
}

type Board struct {
	Size    int
	Rows    Rows
	BR      int
	BC      int
	HeurFun func() int
}

func (b *Board) Copy() *Board {
	r := &Board{
		Size:    b.Size,
		BR:      b.BR,
		BC:      b.BC,
		HeurFun: b.HeurFun,
		Rows:    b.Rows.Copy(),
	}
	return r
}

func New(size int) *Board {
	b := &Board{Size: size}
	b.initiate()
	b.SetManDist()
	return b
}

func (b *Board) initiate() {
	b.Rows = make([]Row, b.Size, b.Size)
	for i := 0; i < b.Size; i++ {
		r := make([]Tile, b.Size, b.Size)
		b.Rows[i] = r
	}
	plog.Info.Println("board initiated")
}

func (b *Board) StateString() string {
	str := strconv.Itoa(b.Size) + " "
	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			str += b.Rows[i][j].String() + " "
		}
	}
	return str
}

func (b *Board) GoalString() string {
	str := strconv.Itoa(b.Size) + " "
	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			str += strconv.Itoa(i*b.Size+j) + " "
		}
	}
	return str
}

func (b *Board) String() string {
	str := "\n"
	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			str += b.Rows[i][j].String() + " "
		}
		str = strings.Trim(str, " ")
		str += "\n"
	}
	return str
}

func (b *Board) Input(Values []int) error {
	err := b.checkInputLen(Values)
	if err != nil {
		return err
	}
	err = b.checkNumbers(Values)
	if err != nil {
		return err
	}
	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			b.Rows[i][j].Val = Values[i*b.Size+j]
			if Values[i*b.Size+j] == 0 {
				b.BR = i
				b.BC = j
			}
		}
	}
	plog.Info.Println("input Values are succesfully imported to board", Values)
	return nil
}
func (b *Board) checkInputLen(Values []int) error {
	if len(Values) != b.Size*b.Size {
		err := errors.New("length of board Values does not match size of board")
		plog.Error.Println(err)
		return err
	}
	plog.Info.Println("length checked")
	return nil
}

func (b *Board) checkNumbers(Values []int) error {
	for i := 0; i < len(Values); i++ {
		found := false
		j := 0
		for j < len(Values) {
			if i == Values[j] {
				found = true
				break
			}
			j++
		}
		if found == false {
			err := errors.New("not all numbers are present in input")
			plog.Error.Println(err)
			return err
		}
	}
	plog.Info.Println("numbers checked")
	return nil
}

//moveMap is initialised in init()
var moveMap map[int]func(*Board) error

func (b *Board) Move(dir int) error {
	if dir < 4 && dir >= 0 {
		err := moveMap[dir](b)
		if err != nil {
			return err
		}
	}
	//plog.Info.Println(b)
	return nil
}

func (b *Board) moveUp() error {
	if b.BR == b.Size-1 {
		err := errors.New("cannot slide up with the blanc Tile on bottom Row")
		plog.Warning.Println(err)
		return err
	}
	b.Rows[b.BR][b.BC], b.Rows[b.BR+1][b.BC] = b.Rows[b.BR+1][b.BC], b.Rows[b.BR][b.BC]
	b.BR++
	plog.Info.Println("Up")
	return nil
}

func (b *Board) moveDown() error {
	if b.BR == 0 {
		err := errors.New("cannot slide down with the blanc Tile on top Row")
		plog.Warning.Println(err)
		return err
	}
	b.Rows[b.BR][b.BC], b.Rows[b.BR-1][b.BC] = b.Rows[b.BR-1][b.BC], b.Rows[b.BR][b.BC]
	b.BR--
	plog.Info.Println("Down")
	return nil
}

func (b *Board) moveLeft() error {
	if b.BC == b.Size-1 {
		err := errors.New("cannot slide left  with the blanc Tile on right collumn")
		plog.Warning.Println(err)
		return err
	}
	b.Rows[b.BR][b.BC], b.Rows[b.BR][b.BC+1] = b.Rows[b.BR][b.BC+1], b.Rows[b.BR][b.BC]
	b.BC++
	plog.Info.Println("Left")
	return nil
}

func (b *Board) moveRight() error {
	if b.BC == 0 {
		err := errors.New("cannot slide right  with the blanc Tile on left collumn")
		plog.Warning.Println(err)
		return err
	}
	b.Rows[b.BR][b.BC], b.Rows[b.BR][b.BC-1] = b.Rows[b.BR][b.BC-1], b.Rows[b.BR][b.BC]
	b.BC--
	plog.Info.Println("Right")
	return nil
}

func (b *Board) SetManDist() {
	b.HeurFun = b.manDist
}

func (b *Board) manDist() int {
	h := 0
	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			fx := b.Rows[i][j].Val % b.Size
			fy := b.Rows[i][j].Val / b.Size
			h += solver.CalcMD(i, j, fx, fy)
		}
	}
	return h
}

func (b *Board) GetH() int {
	return b.HeurFun()
}

func (b *Board) GetMoves() []solver.State {
	ret := make([]solver.State, 0, 4)
	for i := 0; i < 4; i++ {
		move := b.Copy()
		err := move.Move(i)
		if err == nil {
			ret = append(ret, move)
		} else {
			plog.Warning.Println("move", i, "not possible")
		}
	}
	return ret
}

func init() {
	moveMap = make(map[int]func(*Board) error)
	moveMap[Up] = (*Board).moveUp
	moveMap[Down] = (*Board).moveDown
	moveMap[Left] = (*Board).moveLeft
	moveMap[Right] = (*Board).moveRight
}
