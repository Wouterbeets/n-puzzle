package main

import (
	"flag"
	"github.com/Wouterbeets/n-puzzle/board"
	"github.com/Wouterbeets/n-puzzle/input"
	"github.com/Wouterbeets/n-puzzle/plog"
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

func main() {
	flag.Parse()
	plog.Activate(ShowInfo, ShowWarning, ShowError, Verbose)

	b := board.New(3)
	b.Input([]int{0, 1, 2, 3, 4, 5, 6, 7, 8})
	plog.Info.Println("board initailised", b)
	input.GetInput()
	b.Move(board.Up)
	b.Move(board.Left)
	b.Move(board.Down)
	b.Move(board.Right)
}
