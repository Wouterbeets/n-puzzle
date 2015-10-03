package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/Wouterbeets/n-puzzle/board"
	"github.com/Wouterbeets/n-puzzle/checker"
	"github.com/Wouterbeets/n-puzzle/generate"
	"github.com/Wouterbeets/n-puzzle/input"
	"github.com/Wouterbeets/n-puzzle/plog"
	"github.com/Wouterbeets/n-puzzle/solver"
	"github.com/Wouterbeets/n-puzzle/solver_2"
	"github.com/Wouterbeets/n-puzzle/solver_3"
	"math/rand"
	"os"
	"time"
)

var (
	showInfo    bool
	showWarning bool
	showError   bool
	verbose     bool
	stdin       bool
	file        string
	mapSize     int
)

func init() {
	flag.BoolVar(&showWarning, "w", false, "show warnings")
	flag.BoolVar(&showError, "e", true, "show error")
	flag.BoolVar(&showInfo, "i", false, "show info")
	flag.BoolVar(&verbose, "v", false, "show everything")
	flag.BoolVar(&stdin, "s", false, "expecting input from stdin")
	flag.StringVar(&file, "f", "", "input from file, usage:\"n-puzzle -f=filename.txt\"")
	flag.IntVar(&mapSize, "size", 3, "size of generated map")
}

func chooseInput(filename string, stdin bool) (size int, inp []int, err error) {
	if stdin == true {
		plog.Info.Println("Using stdin")
		size, inp, err = input.GetInput(os.Stdin)
	} else if file != "" {
		plog.Info.Println("Using file")
		reader, err := os.Open(file)
		if err != nil {
			os.Exit(1)
		}
		size, inp, err = input.GetInput(reader)
	} else {
		err = errors.New("no input from file or stdin")
	}
	return
}

func main() {
	var b *board.Board
	svr2 := new(solver_2.Solver)
	svr3 := new(solver_3.Solver)
	var err error

	flag.Parse()
	plog.Activate(showInfo, showWarning, showError, verbose)
	size, inp, err := chooseInput(file, stdin)
	if err != nil {
		fmt.Println(err)
		plog.Info.Println("generating map")
		rand.Seed(time.Now().Unix())
		b, err = generate.GetMap(mapSize)
		if err != nil {
			plog.Error.Println(err)
			return
		}
	} else {
		b = board.New(size)
		b.Input(inp)
	}
	fmt.Println(b)
	if checker.CheckerBoard(b) == true {
		fmt.Println("Map is solvable")
		s := solver.New(b)
		s.Solve()
		svr2.Solve_init(b)
		svr2.Solve()
		svr3.Solve_init(b)
		svr3.Solve()
	} else {
		fmt.Println("Map is unsolvable")
	}
}
