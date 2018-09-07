// +build ignore

// gen_lu.go generates lu.go, the giant dictionary mapping traditional characters to jyutping

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

var M = map[string]string{}
var lineRE = regexp.MustCompile(`([^ ]+) ([^ ]+).*\{(.*)}`)

func parseLine(s string) (trad1, trad2, jyutping string) {
	matches := lineRE.FindStringSubmatch(s)
	if matches == nil {
		return
	}
	return matches[1], matches[2], matches[3]
}

func insert(k, v string) {
	M[k] = v
	for i := 1; i < len(k); i++ {
		p := k[:i]
		if _, filled := M[p]; !filled {
			M[p] = ""
		}
	}
}

func parseFile(fname string) map[string]string {
	f, err := os.Open(fname)
	if err != nil {
		fatal(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t1, t2, jp := parseLine(scanner.Text())
		if t1 == "" {
			continue
		}
		insert(t1, jp)
		insert(t2, jp)
	}
	if scanner.Err() != nil {
		fatal(scanner.Err())
	}
	return nil
}

func printit() {
	f, err := os.Create("lu.go")
	if err != nil {
		fatal(err)
	}
	kk := make([]string, 0, len(M))
	for k := range M {
		kk = append(kk, k)
	}
	sort.Strings(kk)
	fmt.Fprintln(f, `package lu

var M = map[string]string{`)
	for _, k := range kk {
		v := M[k]
		fmt.Fprintf(f, "	%s: %s,\n", strconv.Quote(k), strconv.Quote(v))
	}
	fmt.Fprintln(f, `}`)
}

func main() {
	parseFile("cccanto-webdist.txt")
	parseFile("cccedict-canto-readings-150923.txt")
	parseFile("extras.txt")
	printit()
}

func fatal(x ...interface{}) {
	fmt.Fprintln(os.Stderr, x...)
	os.Exit(1)
}
