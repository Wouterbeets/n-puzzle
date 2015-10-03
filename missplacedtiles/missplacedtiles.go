package missplacedtiles

import ()

func Get_missplacedTiles(tab []int, nbrRow int, max int) int {
	sum := 0
	for i := 0; i < max-1; i++ {
		if tab[i] != i+1 {
			sum++
		}
	}
	return sum
}
