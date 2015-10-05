package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/Wouterbeets/n-puzzle/board"
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
			plog.Error.Println(err)
			os.Exit(1)
		}
		size, inp, err = input.GetInput(reader)
	} else {
		err = errors.New("no input from file or stdin")
	}
	return
}

func handleErr(err error) {
	if err != nil {
		plog.Error.Println(err)
		os.Exit(-1)
	}
}

func askHeur() (solv int, heur int) {
	fmt.Printf("Choose your heuristix function\n")
	fmt.Printf("%s\tType 1 for Manhattan Distance\n", GREEN)
	fmt.Printf("%s\tType 2 for MD with linear conflict\n", BLUE)
	fmt.Printf("%s\tType 3 for Misplaced Tiles%s\n", RED, RESET)
	_, err := fmt.Scanf("%d", &heur)
	if err != nil {
		fmt.Printf("Not a valid number: %s\n", err)
		os.Exit(-1)
	}
	fmt.Printf("Choose which memomry method to use:\n")
	fmt.Printf("%s\tType 1 for heap with array's\n", GREEN)
	fmt.Printf("%s\tType 2 for list with array\n", BLUE)
	fmt.Printf("%s\tType 3 or list with double array [][]int%s\n", RED, RESET)
	_, err = fmt.Scanf("%d", &solv)
	if err != nil {
		fmt.Printf("Not a valid number: %s\n", err)
		os.Exit(-1)
	}
	plog.Info.Println("chose number", solv, "as solver and", heur, "as heuristic function")
	if (heur < 1 || heur > 3) && (solv < 1 || solv > 3) {
		fmt.Printf("Invalid choice\n")
		os.Exit(0)
	}
	return
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
	solv, heur := 0, 0

	flag.Parse()
	plog.Activate(showInfo, showWarning, showError, verbose)
	if size, inp, err := chooseInput(file, stdin); err != nil {
		fmt.Println(err)
		plog.Info.Println("generating map")
		rand.Seed(time.Now().Unix())
		solv, heur = askHeur()
		b, err = generate.GetMap(mapSize, heur)
		handleErr(err)
	} else {
		solv, heur = askHeur()
		b = board.New(size, heur)
		err := b.Input(inp)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Printf("Map %d*%d is solvable.\n", b.Size, b.Size)
	fmt.Println(b)
	if solv == 1 {
		s := solver.New(b)
		s.Solve()
	} else if solv == 2 {
		svr2.Solve_init(b, heur)
		svr2.Solve()
	} else if solv == 3 {
		svr3.Solve_init(b, heur)
		svr3.Solve()
	}
	return
}
