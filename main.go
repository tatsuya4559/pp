package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(unindentDiffMarker(line))
	}
}

func isWhitespace(r rune) bool {
	return r == ' ' || r == '\t'
}

func isDiffMarker(r rune) bool {
	return r == '+' || r == '-' || r == '~'
}

func unindentDiffMarker(line string) string {
	var i int
	var r rune
	// Skip leading whitespaces
	for i, r = range line {
		if !isWhitespace(r) {
			break
		}
	}
	if isDiffMarker(r) {
		return fmt.Sprintf("%c%s%s", r, line[:i], line[i+1:])
	}
	return line
}
