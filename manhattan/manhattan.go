package manhattan

import ()

func get_cout(tab []int, value int, nbrRow int) int {
	sum := 0
	indexGood := value - 1
	indexValue := 0
	for tab[indexValue] != value {
		indexValue++
	}
	if indexGood > indexValue {
		sum = ((indexGood / nbrRow) - (indexValue / nbrRow))
		if (indexGood % nbrRow) > (indexValue % nbrRow) {
			sum += ((indexGood % nbrRow) - (indexValue % nbrRow))
		} else {
			sum += ((indexValue % nbrRow) - (indexGood % nbrRow))
		}
	} else {
		sum = ((indexValue / nbrRow) - (indexGood / nbrRow))
		if (indexGood % nbrRow) > (indexValue % nbrRow) {
			sum += ((indexGood % nbrRow) - (indexValue % nbrRow))
		} else {
			sum += ((indexValue % nbrRow) - (indexGood % nbrRow))
		}
	}
	return sum
}

func Get_manhattan_dis(tab []int, nbrRow int, max int) int {
	sum := 0
	value := 1
	for value < max {
		sum += get_cout(tab, value, nbrRow)
		value++
	}
	return sum
}

func get_conflict(tab []int, idxValue int, nbrRow int) int {
	sum := 0
	maxConflict := ((idxValue % nbrRow) * nbrRow) + nbrRow
	value := tab[idxValue]
	if value >= maxConflict-nbrRow && value < maxConflict {
		for i := idxValue; i < maxConflict; i++ {
			if tab[i] >= maxConflict-nbrRow && value > tab[i] {
				sum += 2
			}
		}
	}
	return sum
}

func get_cout_linear(tab []int, value int, nbrRow int) int {
	sum := 0
	indexGood := value - 1
	indexValue := 0
	for tab[indexValue] != value {
		indexValue++
	}
	if indexGood > indexValue {
		sum = ((indexGood / nbrRow) - (indexValue / nbrRow))
		if (indexGood % nbrRow) > (indexValue % nbrRow) {
			sum += ((indexGood % nbrRow) - (indexValue % nbrRow))
		} else {
			sum += ((indexValue % nbrRow) - (indexGood % nbrRow))
		}
	} else {
		sum = ((indexValue / nbrRow) - (indexGood / nbrRow))
		if (indexGood % nbrRow) > (indexValue % nbrRow) {
			sum += ((indexGood % nbrRow) - (indexValue % nbrRow))
		} else {
			sum += ((indexValue % nbrRow) - (indexGood % nbrRow))
		}
	}
	sum += get_conflict(tab, indexValue, nbrRow)
	return sum
}

func Get_manhattan_dis_linear(tab []int, nbrRow int, max int) int {
	sum := 0
	value := 1
	for value < max {
		sum += get_cout_linear(tab, value, nbrRow)
		value++
	}
	return sum
}
