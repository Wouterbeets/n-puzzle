package main

import (
	"flag"
	"fmt"
	"github.com/Wouterbeets/n-puzzle/board"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	ShowInfo    bool
	ShowWarning bool
	ShowError   bool
	Verbose     bool
)

func init() {
	flag.BoolVar(&ShowWarning, "Warning", false, "Should the program show warning messages during execution")
	flag.BoolVar(&ShowError, "Error", true, "Should the program show warning messages during execution")
	flag.BoolVar(&ShowInfo, "Info", false, "Should the program show info messages during execution")
	flag.BoolVar(&Verbose, "Verbose", false, "Should the program show all messages during execution")
}

func main() {
	flag.Parse()
	initLoggers()
	b := board.New(10)
	board.Info.Println("board initailised " + b.String())
}

func initLoggers() {
	var iOut io.Writer
	var eOut io.Writer
	var wOut io.Writer
	var iFile io.Writer
	var wFile io.Writer
	var eFile io.Writer

	iFile, err := os.Create("infoLog")
	if err != nil {
		fmt.Println(err)
		iFile = ioutil.Discard
	}
	eFile, err = os.Create("warningLog")
	if err != nil {
		fmt.Println(err)
		eFile = ioutil.Discard
	}
	wFile, err = os.Create("errLog")
	if err != nil {
		fmt.Println(err)
		wFile = ioutil.Discard
	}

	if ShowInfo == true {
		iOut = os.Stdout
	} else {
		iOut = ioutil.Discard
	}
	if ShowWarning == true {
		wOut = os.Stdout
	} else {
		wOut = ioutil.Discard
	}
	if ShowError == true {
		eOut = os.Stdout
	} else {
		eOut = ioutil.Discard
	}

	iOut = io.MultiWriter(iOut, iFile)
	eOut = io.MultiWriter(eOut, eFile)
	wOut = io.MultiWriter(wOut, wFile)
	board.Info = log.New(iOut, "INFO : ", log.Ltime|log.Lshortfile)
	board.Warning = log.New(wOut, "WARNING : ", log.Ltime|log.Lshortfile)
	board.Error = log.New(eOut, "ERROR : ", log.Ltime|log.Lshortfile)
}
