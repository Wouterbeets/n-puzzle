package board

import (
	"strconv"
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
			b.Rows[i][j] = 1
		}
	}
}

func (b *Board) String() string {
	str := ""
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			str += b.Rows[i][j].String() + " "
		}
	}
	return str
}
