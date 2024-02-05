// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wordCount := make(map[string]int)
	inputArgs := os.Args[1:]
	if len(inputArgs) == 0 {
		countLine(os.Stdin, wordCount)
	}
	for _, arg := range inputArgs {
		file, err := os.Open(arg)
		if err != nil {
			continue
		}
		countLine(file, wordCount)
		file.Close()
	}
	for line, n := range wordCount {
		fmt.Printf("%d\t%s\n", n, line)
	}
}

func countLine(f *os.File, counts map[string]int) {
	lineReader := bufio.NewScanner(f)
	for lineReader.Scan() {
		counts[lineReader.Text()]++
	}
}
