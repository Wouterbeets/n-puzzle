package board

import (
	"errors"
	"github.com/Wouterbeets/n-puzzle/plog"
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

type Board struct {
	LastMove int
	Size     int
	Tiles    []int
	BR       int
	BC       int
	HeurFun  func(int, int, int, int) int
}

func (b *Board) Copy() *Board {
	r := &Board{
		Size: b.Size,
		BR:   b.BR,
		BC:   b.BC,
	}
	copy(r.Tiles, b.Tiles)
	r.HeurFun = b.HeurFun
	return r
}

func (b *Board) initiate() {
	b.Tiles = make([]int, b.Size*b.Size, b.Size*b.Size)
	plog.Info.Println("board initiated")
}

func New(size int) *Board {
	b := new(Board)
	b.Size = size
	b.initiate()
	b.SetOutOfPlace()
	return b
}

func (b *Board) StateString() string {
	str := strconv.Itoa(b.Size) + " "
	for i := 0; i < b.Size*b.Size; i++ {
		str += strconv.Itoa(b.Tiles[i]) + " "
	}
	return str
}

func (b *Board) GoalString() string {
	str := strconv.Itoa(b.Size) + " "
	for i := 0; i < b.Size*b.Size-1; i++ {
		str += strconv.Itoa(i+1) + " "
	}
	str += "0 "
	return str
}

func (b *Board) String() string {
	str := "\n"
	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			str += strconv.Itoa(b.Tiles[b.Size*i+j]) + " "
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
	b.Tiles = Values
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
	return nil
}

func (b *Board) SetManDist() {
	b.HeurFun = CalcMD
}

func (b *Board) SetOutOfPlace() {
	b.HeurFun = OutOfPlace
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

func OutOfPlace(x, y, fx, fy int) int {
	if x == fx && y == fy {
		return 0
	}
	return 1
}

func (b *Board) GetH() int {
	var (
		h, fx, fy int
	)
	for y := 0; y < b.Size; y++ {
		for x := 0; x < b.Size; x++ {
			if b.Rows[y][x].Val == 0 {
				fx = b.Size - 1
				fy = b.Size - 1
			} else {
				fx = (b.Rows[y][x].Val - 1) % b.Size
				fy = (b.Rows[y][x].Val - 1) / b.Size
			}
			h += b.HeurFun(x, y, fx, fy)
		}
	}
	return h
}

func (b *Board) GetMoves() []*Board {
	ret := make([]*Board, 0, 4)
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
