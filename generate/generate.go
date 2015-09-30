package generate

import (
	"github.com/Wouterbeets/n-puzzle/board"
	"math/rand"
)

const (
	MAX_SIZE = 4
)

// Return a new slice with all map's Values
// sVal = slice_Value
func Get_slice(size int) (sVal []int) {
	size = size * size
	for i := 0; i < size; i++ {
		sVal = append(sVal, i)
	}
	return sVal
}

// dVal = delete_Value
func Delete_elem(g_slice []int, index int) (n_slice []int) {
	size := len(g_slice)
	for i := 0; i < size; i++ {
		if i != index {
			n_slice = append(n_slice, g_slice[i])
		}
	}
	return n_slice
}

func Get_Value(g_slice []int) ([]int, int) {
	size := len(g_slice)
	index := rand.Intn(size)
	Value := g_slice[index]
	g_slice = Delete_elem(g_slice, index)
	return g_slice, Value
}

func GetMap(b *board.Board) (*board.Board, error) {
	size := rand.Intn(MAX_SIZE + 1)
	if size < 3 {
		size = 3
	}
	b.New(size)
	sVal := Get_slice(size)
	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			sVal, b.Rows[i][j].Val = Get_Value(sVal)
			if b.Rows[i][j].Val == 0 {
				b.BR = i
				b.BC = j
			}
		}
	}
	return b, nil
}
