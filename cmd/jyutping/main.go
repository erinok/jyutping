// Command jyutping converts chinese characters to jyutping (cantonese romanization).
package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/erinok/jyutping"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(jyutping.Convert(scanner.Text()))
	}
	fmt.Println()
}

func fatal(x ...interface{}) {
	fmt.Fprintln(os.Stderr, x...)
	os.Exit(1)
}
