package main

import (
	"flag"
	"github.com/Wouterbeets/n-puzzle/board"
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
	plog.Info.Println("board initailised", b)
	plog.Warning.Println("test")
	plog.Error.Println("test")
}
