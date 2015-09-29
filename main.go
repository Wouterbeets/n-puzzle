package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/Wouterbeets/n-puzzle/board"
	//	"github.com/Wouterbeets/n-puzzle/generate"
	"github.com/Wouterbeets/n-puzzle/input"
	"github.com/Wouterbeets/n-puzzle/plog"
	"github.com/Wouterbeets/n-puzzle/solver"
	//	"math/rand"
	"os"
	//	"time"
)

var (
	showInfo    bool
	showWarning bool
	showError   bool
	verbose     bool
	stdin       bool
	file        string
)

func init() {
	flag.BoolVar(&showWarning, "w", false, "show warnings")
	flag.BoolVar(&showError, "e", true, "show error")
	flag.BoolVar(&showInfo, "i", false, "show info")
	flag.BoolVar(&verbose, "v", false, "show everything")
	flag.BoolVar(&stdin, "s", false, "expecting input from stdin")
	flag.StringVar(&file, "f", "", "input from file, usage:\"n-puzzle -f=filename.txt\"")
}

func chooseInput(filename string, stdin bool) (size int, inp []int, err error) {
	if stdin == true {
		size, inp, err = input.GetInput(os.Stdin)
	} else if file != "" {
		reader, err := os.Open(file)
		if err != nil {
			os.Exit(1)
		}
		size, inp, err = input.GetInput(reader)
	} else {
		err = errors.New("no input")
	}
	return
}

// Where you call init function ?
func main() {
	var b *board.Board
	//	var err error

	flag.Parse()
	plog.Activate(showInfo, showWarning, showError, verbose)
	size, inp, _ := chooseInput(file, stdin)
	//	if err != nil {
	//		rand.Seed(time.Now().Unix())
	//		b, err = generate.GetMap()
	//		if err != nil {
	//			plog.Error.Println(err)
	//			return
	//		}
	//	} else {
	b = board.New(size)
	b.Input(inp)
	//	}
	fmt.Println(b)
	s := solver.New(b)
	s.Solve()
}
