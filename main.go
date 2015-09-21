package main

import (
	"flag"
	"github.com/Wouterbeets/n-puzzle/board"
	"github.com/Wouterbeets/n-puzzle/input"
	"github.com/Wouterbeets/n-puzzle/plog"
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

func main() {
	flag.Parse()
	plog.Activate(ShowInfo, ShowWarning, ShowError, Verbose)
	size, input, err := input.GetInput(os.Stdin)
	if err != nil {
		plog.Error.Println(err)
		return
	}
	b := board.New(size)
	b.Input(input)
	plog.Info.Println("board initailised", b)
	b.Move(board.Up)
	plog.Info.Println(b.GetH())
	b.Move(board.Left)
	plog.Info.Println(b.GetH())
	b.Move(board.Down)
	plog.Info.Println(b.GetH())
	b.Move(board.Right)
}
