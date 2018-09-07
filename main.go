// Command jyutping converts traditional chinese characters to jyutping (cantonese romanization).
package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/erinok/jyutping/lu"
)

func convert(s string) {
	if s == "" {
		return
	}
	t := ""
	j := 0
	i := 1
	for {
		t_, prefix := lu.M[s[:i]]
		if !prefix {
			break
		}
		if t_ != "" {
			t = t_
			j = i
		}
		if i == len(s) {
			break
		}
		i++
	}
	if t != "" {
		fmt.Print(t)
		fmt.Print(" ")
		convert(s[j:])
	} else {
		fmt.Print(s[:i])
		fmt.Print(" ")
		convert(s[i:])
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		convert(scanner.Text())
	}
	fmt.Println()
}

func fatal(x ...interface{}) {
	fmt.Fprintln(os.Stderr, x...)
	os.Exit(1)
}
