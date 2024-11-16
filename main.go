package main

import "fmt"

func main() {

}

func isWhiteSpace(r rune) bool {
	return r == ' '
}

func isDiffMarker(r rune) bool {
	return r == '+'
}

func unindentDiffMarker(line string) string {
	for i, r := range line {
		if isWhiteSpace(r) {
			continue
		}
		if isDiffMarker(r) {
			return fmt.Sprintf("%c%s%s", r, line[:i], line[i+1:])
		} else {
			return line
		}
	}
	return line
}
