package input

import (
	"fmt"
	"os"
	"testing"
)

func TestGetInput(t *testing.T) {
	str := "3\n1 2 3\n4#Comment# 5 6\n7 8 0\n"
	fmt.Fprintln(os.Stdin, str)
	GetInput()
}
