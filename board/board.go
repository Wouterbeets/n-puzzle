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
)

type tile struct {
	val int
}

func (t tile) String() string {
	return strconv.Itoa(t.val)
}

type row []tile

func (r row) Copy() row {
	ret := make([]tile, len(r))
	copy(ret, r)
	return ret
}

type Rows []row

func (r Rows) Copy() Rows {
	ret := make([]row, len(r))
	for i := 0; i < len(r); i++ {
		ret[i] = r[i].Copy()
	}
	return ret
}

type Board struct {
	size     int
	Rows     Rows
	BR       int
	BC       int
	HeurFun  func() int
	LastMove int
}

func (b *Board) Copy() *Board {
	r := &Board{
		size: b.size,
		BR:   b.BR,
		BC:   b.BC,
		Rows: b.Rows.Copy(),
	}
	r.HeurFun = r.manDist
	return r
}

func New(size int) *Board {
	b := &Board{size: size}
	b.initiate()
	b.SetManDist()
	return b
}

func (b *Board) initiate() {
	b.Rows = make([]row, b.size, b.size)
	for i := 0; i < b.size; i++ {
		r := make([]tile, b.size, b.size)
		b.Rows[i] = r
	}
	plog.Info.Println("board initiated")
}

func (b *Board) StateString() string {
	str := strconv.Itoa(b.size) + " "
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			str += b.Rows[i][j].String() + " "
		}
	}
	return str
}

func (b *Board) GoalString() string {
	str := strconv.Itoa(b.size) + " "
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			if i == b.size-1 && j == b.size-1 {
				str += "0 "
			} else {
				str += strconv.Itoa(i*b.size+j+1) + " "
			}
		}
	}
	return str
}

func (b *Board) String() string {
	str := "\n"
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			str += b.Rows[i][j].String() + " "
		}
		str = strings.Trim(str, " ")
		str += "\n"
	}
	return str
}

func (b *Board) Input(values []int) error {
	err := b.checkInputLen(values)
	if err != nil {
		return err
	}
	err = b.checkNumbers(values)
	if err != nil {
		return err
	}
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			b.Rows[i][j].val = values[i*b.size+j]
			if values[i*b.size+j] == 0 {
				b.BR = i
				b.BC = j
			}
		}
	}
	plog.Info.Println("input values are succesfully imported to board", values)
	return nil
}
func (b *Board) checkInputLen(values []int) error {
	if len(values) != b.size*b.size {
		err := errors.New("length of board values does not match size of board")
		plog.Error.Println(err)
		return err
	}
	plog.Info.Println("length checked")
	return nil
}

func (b *Board) checkNumbers(values []int) error {
	for i := 0; i < len(values); i++ {
		found := false
		j := 0
		for j < len(values) {
			if i == values[j] {
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
	return nil
}

func (b *Board) moveUp() error {
	if b.BR == b.size-1 {
		err := errors.New("cannot slide up with the blanc tile on bottom row")
		plog.Warning.Println(err)
		return err
	}
	b.Rows[b.BR][b.BC], b.Rows[b.BR+1][b.BC] = b.Rows[b.BR+1][b.BC], b.Rows[b.BR][b.BC]
	b.BR++
	return nil
}

func (b *Board) moveDown() error {
	if b.BR == 0 {
		err := errors.New("cannot slide down with the blanc tile on top row")
		plog.Warning.Println(err)
		return err
	}
	b.Rows[b.BR][b.BC], b.Rows[b.BR-1][b.BC] = b.Rows[b.BR-1][b.BC], b.Rows[b.BR][b.BC]
	b.BR--
	return nil
}

func (b *Board) moveLeft() error {
	if b.BC == b.size-1 {
		err := errors.New("cannot slide left  with the blanc tile on right collumn")
		plog.Warning.Println(err)
		return err
	}
	b.Rows[b.BR][b.BC], b.Rows[b.BR][b.BC+1] = b.Rows[b.BR][b.BC+1], b.Rows[b.BR][b.BC]
	b.BC++
	return nil
}

func (b *Board) moveRight() error {
	if b.BC == 0 {
		err := errors.New("cannot slide right  with the blanc tile on left collumn")
		plog.Warning.Println(err)
		return err
	}
	b.Rows[b.BR][b.BC], b.Rows[b.BR][b.BC-1] = b.Rows[b.BR][b.BC-1], b.Rows[b.BR][b.BC]
	b.BC--
	return nil
}

func (b *Board) SetManDist() {
	b.HeurFun = b.manDist
}

func (b *Board) manDist() int {
	h := 0
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			fx := b.Rows[i][j].val % b.size
			fy := b.Rows[i][j].val / b.size
			h += solver.CalcMD(j, i, fx, fy)
			//plog.Info.Println("val =", b.Rows[i][j].val, "fx", fx, "fy", fy, "h", h)
		}
	}
	return h
}

func (b *Board) GetH() int {
	var (
		h, fx, fy int
	)
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			if b.Rows[i][j].val == 0 {
				fx = b.size - 1
				fy = b.size - 1
			} else {
				fx = (b.Rows[i][j].val - 1) % b.size
				fy = (b.Rows[i][j].val - 1) / b.size
			}
			h += solver.CalcMD(j, i, fx, fy)
			//plog.Info.Println("val =", b.Rows[i][j].val, "fx", fx, "fy", fy, "h", h)
		}
	}
	return h
	//return b.HeurFun()
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

func (b *Board) GetLastMove() int {
	return b.LastMove
}

func init() {
	moveMap = make(map[int]func(*Board) error)
	moveMap[Up] = (*Board).moveUp
	moveMap[Down] = (*Board).moveDown
	moveMap[Left] = (*Board).moveLeft
	moveMap[Right] = (*Board).moveRight
}
