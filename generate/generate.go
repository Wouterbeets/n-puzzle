package generate

import (
	"github.com/Wouterbeets/n-puzzle/board"
	"math/rand"
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

func GetMap(size int, heur int) (*board.Board, error) {
	var err error

	if size < 3 {
		size = 3
	}
	b := board.New(size, heur)
	sVal := Get_slice(size)
	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			sVal, b.Tiles[i*b.Size+j] = Get_Value(sVal)
			if b.Tiles[i*b.Size+j] == 0 {
				b.BR = i
				b.BC = j
			}
		}
	}
	if b.CheckBoard() == false {
		b, err = GetMap(size, heur)
		if err != nil {
			return nil, err
		}
	}
	return b, nil
}
