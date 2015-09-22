package input

import (
	"bufio"
	"github.com/Wouterbeets/n-puzzle/plog"
	"io"
	//"os"
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

func skipComment(start int, data []byte) int {
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if r == '\n' {
			return i + width
		}
	}
	return 0
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
		start = skipComment(start, data)
	}
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
	// If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	// Request more data.
	return start, nil, nil
}

func GetInput(r io.Reader) (int, []int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(scanWords)
	ret := make([]int, 0, 0)
	size := -1
	for scanner.Scan() {
		token := scanner.Text()
		if token == "" {
			continue
		}
		t, err := strconv.Atoi(token)
		if err != nil {
			plog.Error.Println(err)
			return size, ret, err
		}
		if size == -1 {
			size = t
		} else {
			ret = append(ret, t)
		}
	}
	if err := scanner.Err(); err != nil {
		plog.Error.Println(err)
		return size, ret, err
	}
	return size, ret, nil
}
