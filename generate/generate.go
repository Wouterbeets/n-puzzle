package generate

import (
	"github.com/Wouterbeets/n-puzzle/board"
	"math/rand"
)

const (
	MAX_SIZE = 42
)

// Return a new slice with all map's values
// sval = slice_value
func Get_slice(size int) (sval []int) {
	size = size * size
	for i := 0; i < size; i++ {
		sval = append(sval, i)
	}
	return sval
}

// dval = delete_value
func Delete_elem(g_slice []int, index int) (n_slice []int) {
	size := len(g_slice) - 1
	for i := 0; i < size; i++ {
		if i != index {
			n_slice = append(n_slice, g_slice[i])
		}
	}
	return n_slice
}

func Get_value(g_slice []int) ([]int, int) {
	size := len(g_slice)
	r := rand.New(rand.NewSource(int64(size)))
	index := r.Int()
	if index == 0 {
		index = 1
	}
	value := g_slice[index]
	g_slice = Delete_elem(g_slice, index)
	return g_slice, value
}

func GetMap() (*board.Board, error) {
	r := rand.New(rand.NewSource(MAX_SIZE))
	size := r.Int()
	if size < 3 {
		size = 3
	}
	b := board.New(size)
	sval := Get_slice(size)
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			sval, b.Rows[i][j].val = Get_value(sval, size)
			if b.Rows[i][j].val == 0 {
				b.BR = i
				b.BC = j
			}
		}
	}
	return b, nil
}
