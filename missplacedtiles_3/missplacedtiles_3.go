package missplacedtiles_3

import ()

func Get_missplacedTiles(tab [][]int, nbrRow int) int {
	sum := 0
	for i := 0; i < nbrRow; i++ {
		for j := 0; j < nbrRow; j++ {
			if tab[i][j] != (i*nbrRow)+(j+1) {
				sum++
			}
		}
	}
	return sum
}
