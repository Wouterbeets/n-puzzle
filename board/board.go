package board

import (
	"errors"
	"github.com/Wouterbeets/n-puzzle/plog"
	"math"
	"strconv"
	"strings"
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
	if len(values) != b.size*b.size {
		err := errors.New("length of board values does not match size of board")
		plog.Error.Println(err)
		return err
	}
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			b.Rows[i][j] = tile(values[i*b.size+j])
		}
	}
	plog.Info.Println("input values are:", values, "board is:", b)
	return nil
}
