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
	//	"github.com/Wouterbeets/n-puzzle/solver"
	"github.com/Wouterbeets/n-puzzle/solver_2"
	//"github.com/Wouterbeets/n-puzzle/solver_3"
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
		fmt.Println("Use stdin")
		size, inp, err = input.GetInput(os.Stdin)
	} else if file != "" {
		fmt.Println("Use file")
		reader, err := os.Open(file)
		if err != nil {
			os.Exit(1)
		}
		size, inp, err = input.GetInput(reader)
	} else {
		fmt.Println("Call generate map, empty argument or map invalid")
		err = errors.New("no input")
	}
	return
}

func main() {
	b := new(board.Board)
	svr := new(solver_2.Solver)
	//svr := new(solver_3.Solver)
	var err error

	flag.Parse()
	plog.Activate(showInfo, showWarning, showError, verbose)
	size, inp, err := chooseInput(file, stdin)
	if err != nil {
		rand.Seed(time.Now().Unix())
		b, err = generate.GetMap(b)
		if err != nil {
			plog.Error.Println(err)
			return
		}
	} else {
		b.New(size)
		b.Input(inp)
	}
	fmt.Println(b)
	if checker.CheckerBoard(b) == true {
		fmt.Println("Map is solvent")
		svr.Solve_init(b)
		svr.Solve()
		//		s := solver.New(b)
		//		s.Solve()
	} else {
		fmt.Println("Map isn't insolvent")
	}
}
