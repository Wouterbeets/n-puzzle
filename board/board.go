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
)

type tile int

func (t tile) String() string {
	return strconv.Itoa(int(t))
}

type row []tile

type Board struct {
	size     int
	Rows     []row
	BlankRow int
	BlankCol int
}

func New(size int) *Board {
	b := &Board{size: size}
	b.initiate()
	return b
}

func (b *Board) initiate() {
	b.Rows = make([]row, b.size, b.size)
	for i := 0; i < b.size; i++ {
		r := make([]tile, b.size, b.size)
		b.Rows[i] = r
		for j := 0; j < b.size; j++ {
			b.Rows[i][j] = 0
		}
	}
	plog.Info.Println("board initiated")
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
			b.Rows[i][j] = tile(values[i*b.size+j])
			if values[i*b.size+j] == 0 {
				b.BlankRow = i
				b.BlankCol = j
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
	plog.Info.Println(b)
	return nil
}

func (b *Board) moveUp() error {
	if b.BlankRow == b.size-1 {
		err := errors.New("cannot slide up with the blanc tile on bottom row")
		plog.Warning.Println(err)
		return err
	}
	b.Rows[b.BlankRow][b.BlankCol] = b.Rows[b.BlankRow+1][b.BlankCol]
	b.Rows[b.BlankRow+1][b.BlankCol] = 0
	b.BlankRow++
	plog.Info.Println("Up")
	return nil
}
func (b *Board) moveDown() error {
	if b.BlankRow == 0 {
		err := errors.New("cannot slide down with the blanc tile on top row")
		plog.Warning.Println(err)
		return err
	}
	b.Rows[b.BlankRow][b.BlankCol] = b.Rows[b.BlankRow-1][b.BlankCol]
	b.Rows[b.BlankRow-1][b.BlankCol] = 0
	b.BlankRow--
	plog.Info.Println("Down")
	return nil
}
func (b *Board) moveLeft() error {
	if b.BlankCol == b.size {
		err := errors.New("cannot slide left  with the blanc tile on right collumn")
		plog.Warning.Println(err)
		return err
	}
	b.Rows[b.BlankRow][b.BlankCol] = b.Rows[b.BlankRow][b.BlankCol+1]
	b.Rows[b.BlankRow][b.BlankCol+1] = 0
	b.BlankCol++
	plog.Info.Println("Left")
	return nil
}
func (b *Board) moveRight() error {
	if b.BlankCol == 0 {
		err := errors.New("cannot slide right  with the blanc tile on left collumn")
		plog.Warning.Println(err)
		return err
	}
	b.Rows[b.BlankRow][b.BlankCol] = b.Rows[b.BlankRow][b.BlankCol-1]
	b.Rows[b.BlankRow][b.BlankCol-1] = 0
	b.BlankCol--
	plog.Info.Println("Right")
	return nil
}

func init() {
	moveMap = make(map[int]func(*Board) error)
	moveMap[Up] = (*Board).moveUp
	moveMap[Down] = (*Board).moveDown
	moveMap[Left] = (*Board).moveLeft
	moveMap[Right] = (*Board).moveRight
}
