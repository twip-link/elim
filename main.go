// Package elim is a command line utility that parrots standard input by way of a bufio.Scanner.
// To limit the number of lines returned, set flag `-l` to an integer other than zero.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s < input.txt -l 0\n", os.Args[0])
		flag.PrintDefaults()
	}

	lastLineIndex := flag.Int("l", 0, "The index of the last line that should run; zero runs all lines.")
	flag.Parse()

	if flag.NFlag() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	var n int
	lines := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			if scanner.Err() != nil {
				fmt.Println(scanner.Err().Error())
			}
			if *lastLineIndex == 0 {
				break
			}
			os.Exit(1)
		}
		if n != 0 && n == *lastLineIndex {
			break
		}

		text := scanner.Text()
		if len(text) > 0 {
			lines = append(lines, text)
		}

		n++
	}

	sow := bufio.NewWriter(os.Stdout)
	defer sow.Flush()
	for i := 0; i < len(lines); i++ {
		fmt.Fprintf(sow, "%v\n", strings.TrimSpace(lines[i]))
	}
}
