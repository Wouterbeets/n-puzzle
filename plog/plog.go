package plog

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	Info = log.New(ioutil.Discard, "", 0)
	Warning = log.New(ioutil.Discard, "", 0)
	Error = log.New(ioutil.Discard, "", 0)
}

func Activate(ShowInfo bool, ShowWarning bool, ShowError bool, Verbose bool) {
	var iOut io.Writer
	var eOut io.Writer
	var wOut io.Writer
	var iFile io.Writer
	var wFile io.Writer
	var eFile io.Writer

	if Verbose == true {
		ShowInfo, ShowError, ShowWarning = true, true, true
	}
	iFile, err := os.Create("infoLog")
	if err != nil {
		fmt.Println(err)
		iFile = ioutil.Discard
	}
	wFile, err = os.Create("warningLog")
	if err != nil {
		fmt.Println(err)
		eFile = ioutil.Discard
	}
	eFile, err = os.Create("errLog")
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
	Info = log.New(iOut, "INFO : ", log.Ltime|log.Lshortfile)
	Warning = log.New(wOut, "WARNING : ", log.Ltime|log.Lshortfile)
	Error = log.New(eOut, "ERROR : ", log.Ltime|log.Lshortfile)
}
