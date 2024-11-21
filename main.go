package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"unicode"
)

func printUsage() {
	usage := `pp is a tool to prettify diff output.

Usage: pp < diff.txt
`
	fmt.Fprintln(os.Stderr, usage)
}

func main() {
	flag.Usage = printUsage
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(unindentDiffMarker(line))
	}
}

func isDiffMarker(r rune) bool {
	return r == '+' || r == '-' || r == '~'
}

func unindentDiffMarker(line string) string {
	var i int
	var r rune
	// Skip leading whitespaces
	for i, r = range line {
		if !unicode.IsSpace(r) {
			break
		}
	}
	if isDiffMarker(r) {
		return fmt.Sprintf("%c%s%s", r, line[:i], line[i+1:])
	}
	return line
}
