package main

import "fmt"

func main() {

}

func isWhitespace(r rune) bool {
	return r == ' '
}

func isDiffMarker(r rune) bool {
	return r == '+'
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
