package input

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"unicode/utf8"
)

func isSpace(r rune) bool {
	if r <= '\u00FF' {
		// Obvious ASCII ones: \t through \r plus space. Plus two Latin-1 oddballs.
		switch r {
		case ' ', '\t', '\n', '\v', '\f', '\r':
			return true
		case '\u0085', '\u00A0':
			return true
		}
		return false
	}
	// High-valued ones.
	if '\u2000' <= r && r <= '\u200a' {
		return true
	}
	switch r {
	case '\u1680', '\u2028', '\u2029', '\u202f', '\u205f', '\u3000':
		return true
	}
	return false
}

//Slightly modified standard lib function, it also terminates words at '#' for comments in de middle of lines
func scanWords(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Skip leading spaces.
	start := 0
	comment := false
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if r == '#' {
			comment = true
			break
		}
		if !isSpace(r) {
			break
		}
	}
	//if state is comment we continue until end of line
	if comment == true {
		for width, i := 0, start; i < len(data); i += width {
			var r rune
			r, width = utf8.DecodeRune(data[i:])
			if r == '\n' {
				return i + width, data[start:i], nil
			}
		}
	} else {
		// Scan until space, marking end of word.
		for width, i := 0, start; i < len(data); i += width {
			var r rune
			r, width = utf8.DecodeRune(data[i:])
			//			also marks end of word at #
			if r == '#' {
				return i, data[start:i], nil
			}
			if isSpace(r) {
				return i + width, data[start:i], nil
			}
		}
	}
	// If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	// Request more data.
	return start, nil, nil
}

func GetInput() {
	scanner := bufio.NewScanner(os.Stdin)
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = scanWords(data, atEOF)
		if err == nil && token != nil && bytes.Contains(token, []byte("#")) == false {
			_, err = strconv.ParseInt(string(token), 10, 32)
		}
		return
	}
	scanner.Split(split)
	i := 0
	for scanner.Scan() {
		fmt.Print(scanner.Text())
		i++
		fmt.Println("\t", i)
	}
	fmt.Println("reading standard input:", scanner.Err())
	if err := scanner.Err(); err != nil {
	}
}
