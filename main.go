package main

import (
	"flag"
	"github.com/Wouterbeets/n-puzzle/board"
	"github.com/Wouterbeets/n-puzzle/generate"
	"github.com/Wouterbeets/n-puzzle/input"
	"github.com/Wouterbeets/n-puzzle/plog"
	"github.com/Wouterbeets/n-puzzle/solver"
	"os"
)

var (
	ShowInfo    bool
	ShowWarning bool
	ShowError   bool
	Verbose     bool
)

func init() {
	flag.BoolVar(&ShowWarning, "w", false, "show warnings")
	flag.BoolVar(&ShowError, "e", true, "show error")
	flag.BoolVar(&ShowInfo, "i", false, "show info")
	flag.BoolVar(&Verbose, "v", false, "show everything")
}

// Where you call init function ?
func main() {
	var b *board.Board
	flag.Parse()
	plog.Activate(ShowInfo, ShowWarning, ShowError, Verbose)
	size, input, err := input.GetInput(os.Stdin)
	if err != nil {
		b, err = generate.GetMap()
		if err != nil {
			plog.Error.Println(err)
			return
		}
	} else {
		b = board.New(size)
		b.Input(input)
	}
	plog.Info.Println("board initailised", b)
	s := solver.New(b)
	s.Solve()
}
