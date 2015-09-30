package manhattan

import (
	"fmt"
)

func get_cout(tab []int, value int, nbrRow int) int {
	sum := 0
	indexGood := value - 1
	indexValue := 0
	for tab[indexValue] != value {
		indexValue++
	}
	fmt.Print("Value good: ")
	fmt.Print(indexGood)
	fmt.Print(" | indexValue: ")
	fmt.Println(indexValue)

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
	fmt.Print("Value heur: ")
	fmt.Println(sum)
	fmt.Print("to Value: ")
	fmt.Println(value)

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

/*
	fmt.Print("(")
	fmt.Print(indexGood / nbrRow)
	fmt.Print(" - ")
	fmt.Print(indexValue / nbrRow)
	fmt.Print(") + (")
	fmt.Print(indexGood % nbrRow)
	fmt.Print(" - ")
	fmt.Print(indexValue % nbrRow)
	fmt.Print(")")
*/
/*
	fmt.Print("(")
	fmt.Print(indexValue / nbrRow)
	fmt.Print(" - ")
	fmt.Print(indexGood / nbrRow)
	fmt.Print(") + (")
	fmt.Print(indexValue % nbrRow)
	fmt.Print(" - ")
	fmt.Print(indexGood % nbrRow)
	fmt.Print(")")
*/
