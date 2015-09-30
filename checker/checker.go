package checker

import (
	"github.com/Wouterbeets/n-puzzle/board"
)

func get_inversion(b *board.Board, ind_i int, ind_j int) int {
	sum := 0
	flag := 1
	j := ind_j

	if b.Rows[ind_i][ind_j].Val == 0 {
		return sum
	}
	for i := ind_i; i < b.Size; i++ {
		if flag == 0 {
			j = 0
		}
		for j < b.Size {
			if b.Rows[ind_i][ind_j].Val > b.Rows[i][j].Val && b.Rows[i][j].Val > 0 {
				sum++
			}
			j++
		}
		flag = 0
	}
	return sum
}

func get_position_zero(b *board.Board) int {
	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			if b.Rows[i][j].Val == 0 {
				return b.Size - i
			}
		}
	}
	return 1
}

func CheckerBoard(b *board.Board) bool {
	sum := 0
	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			sum += get_inversion(b, i, j)
		}
	}
	if b.Size%2 == 0 {
		row := get_position_zero(b)
		if row%2 == 0 && sum%2 != 0 {
			return true
		} else if row%2 != 0 && sum%2 == 0 {
			return true
		}
	} else {
		if sum%2 == 0 {
			return true
		}
	}
	return false
}
