// Package jyutping converts chinese characters to jyutping (cantonese romanization).
package jyutping

import (
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/erinok/jyutping/lu"
)

// Convert replaces chinese characters in s with jyutping.
func Convert(s string) string {
	w := strings.Builder{}
	didtrans := false
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
			if w.Len() > 0 && wantsSpace(t) {
				w.WriteByte(' ')
			}
			w.WriteString(t)
			didtrans = true
			s = s[j:]
		} else {
			if didtrans && wantsSpace(s) {
				w.WriteByte(' ')
			}
			w.WriteString(s[:i])
			didtrans = false
			s = s[i:]
		}
	}
	t := w.String()
	for fr, to := range map[string]string{
		`“ `:  `"`,
		"”":   `"`,
		":  ": ": ",
	} {
		t = strings.Replace(t, fr, to, -1)
	}
	return t
}

func wantsSpace(s string) bool {
	r, _ := utf8.DecodeRuneInString(s)
	return unicode.IsLetter(r) || unicode.IsNumber(r) || r == '“'
}
