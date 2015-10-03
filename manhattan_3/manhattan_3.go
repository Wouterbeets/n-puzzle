package manhattan_3

import ()

func Get_positionValue(board [][]int, value int, nbrRow int) (int, int) {
	for i := 0; i < nbrRow; i++ {
		for j := 0; j < nbrRow; j++ {
			if board[i][j] == value {
				return i, j
			}
		}
	}
	return -1, -1
}

func Get_GoodIndexValue(board [][]int, value int, nbrRow int) (int, int) {
	start := 1
	for i := 0; i < nbrRow; i++ {
		for j := 0; j < nbrRow; j++ {
			if start == value {
				return i, j
			}
			start++
		}
	}
	return -1, -1
}

/* iV == Index_value I | jV == Index_value J */
/* iG == IndexGood_value I | jG == IndexGood_value J */
func get_cout(tab [][]int, value int, nbrRow int) int {
	sum := 0
	iG, jG := Get_GoodIndexValue(tab, value, nbrRow)
	iV, jV := Get_positionValue(tab, value, nbrRow)

	if iG > iV {
		sum = iG - iV
	} else {
		sum = iV - iG
	}
	if jG > jV {
		sum += jG - jV
	} else {
		sum += jV - jG
	}
	return sum
}

func Get_manhattan_dis(tab [][]int, nbrRow int) int {
	sum := 0
	value := 1
	max := (nbrRow * nbrRow)
	for value < max {
		sum += get_cout(tab, value, nbrRow)
		value++
	}
	return sum
}

func get_conflict(tab [][]int, value int, nbrRow int) int {
	sum := 0

	iG, _ := Get_GoodIndexValue(tab, value, nbrRow)
	iV, jV := Get_positionValue(tab, value, nbrRow)
	if iG == iV {
		for j := jV; j < nbrRow; j++ {
			iG, _ = Get_GoodIndexValue(tab, tab[iV][j], nbrRow)
			if iG == iV && value > tab[iV][j] {
				sum += 2
			}
		}
	}
	return sum
}

func Get_manhattan_dis_linear(tab [][]int, nbrRow int) int {
	sum := 0
	value := 1
	max := (nbrRow * nbrRow)
	for value < max {
		sum += get_cout(tab, value, nbrRow)
		sum += get_conflict(tab, value, nbrRow)
		value++
	}
	return sum
}
