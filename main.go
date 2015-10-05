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

const (
	GREEN = "\033[32m"
	BLUE  = "\033[34m"
	RED   = "\033[31m"
	RESET = "\033[0m"
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

type choose struct {
	solv int
	heur int
}

func check_Value(value int) bool {
	if value < 1 || value > 3 {
		return false
	}
	return true
}

func main() {
	var b *board.Board
	svr2 := new(solver_2.Solver)
	svr3 := new(solver_3.Solver)
	var err error

	flag.Parse()
	plog.Activate(showInfo, showWarning, showError, verbose)
	fmt.Printf("Choose your heuristix function\n %sType 1 for Manhattan Distance\n %sType 2 for MD with linear conflict\n %sType 3 for Misplaced Tiles%s\n", GREEN, BLUE, RED, RESET)

	c := choose{}
	fmt.Scanf("%d", &c.heur)
	if check_Value(c.heur) == false {
		fmt.Println("Invalid value to Heuristic\n")
		return
	}
	fmt.Printf("Choose which memomry method to use:\n %sType 1 for heap with array's\n %sType 2 for list with array\n %sType 3 or list with double array [][]int%s\n", GREEN, BLUE, RED, RESET)
	fmt.Scanf("%d", &c.solv)
	if check_Value(c.solv) == false {
		fmt.Println("Invalid value to solver map\n")
		return
	}
	size, inp, err := chooseInput(file, stdin)
	if err != nil {
		fmt.Println(err)
		plog.Info.Println("generating map")
		rand.Seed(time.Now().Unix())
		b, err = generate.GetMap(mapSize, c.heur)
		for checker.CheckerBoard(b) == false {
			b, err = generate.GetMap(mapSize, c.heur)
		}
		if err != nil {
			plog.Error.Println(err)
			return
		}
	} else {
		b = board.New(size, 1)
		b.Input(inp)
	}
	if checker.CheckerBoard(b) == true {
		fmt.Printf("Map %d*%d is solvable.\n", b.Size, b.Size)
		fmt.Println(b.Tiles)
		if c.solv == 1 {
			s := solver.New(b)
			s.Solve()
		} else if c.solv == 2 {
			svr2.Solve_init(b, c.heur)
			svr2.Solve()
		} else {
			svr3.Solve_init(b, c.heur)
			svr3.Solve()
		}
	} else {
		fmt.Println("Map is unsolvable")
	}
}
