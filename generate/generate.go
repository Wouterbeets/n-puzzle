package generate

import (
	"github.com/Wouterbeets/n-puzzle/board"
	"math/rand"
)

const (
	MAX_SIZE = 42
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
	size := len(g_slice) - 1
	for i := 0; i < size; i++ {
		if i != index {
			n_slice = append(n_slice, g_slice[i])
		}
	}
	return n_slice
}

func Get_Value(g_slice []int) ([]int, int) {
	size := len(g_slice)
	r := rand.New(rand.NewSource(int64(size)))
	index := r.Int()
	if index == 0 {
		index = 1
	}
	Value := g_slice[index]
	g_slice = Delete_elem(g_slice, index)
	return g_slice, Value
}

func GetMap() (*board.Board, error) {
	r := rand.New(rand.NewSource(MAX_SIZE))
	size := r.Int()
	if size < 3 {
		size = 3
	}
	b := board.New(size)
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
