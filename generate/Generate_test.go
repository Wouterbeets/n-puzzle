package generate

import (
	"fmt"
	"github.com/Wouterbeets/n-puzzle/board"
)

func TestGetMap() {
	b, err := GetMap()
	if err != nil {
		fmt.Println("Error on GetMap\n")
	} else {
		b.String()
	}
}
