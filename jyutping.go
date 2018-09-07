// Package jyutping converts chinese characters to jyutping (cantonese romanization).
package jyutping

import (
	"strings"

	"github.com/erinok/jyutping/lu"
)

// Convert replaces chinese characters in s with jyutping.
func Convert(s string) string {
	w := strings.Builder{}
	sp := false
	for s != "" {
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
			if sp {
				w.WriteByte(' ')
			}
			w.WriteString(t)
			s = s[j:]
		} else {
			w.WriteString(s[:i])
			s = s[i:]
		}
		sp = true
	}
	return w.String()
}
