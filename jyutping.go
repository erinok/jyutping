// Package jyutping converts chinese characters to jyutping (cantonese romanization).
package jyutping

import (
	"fmt"
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
		`“`:   `"`,
		"”":   `"`,
		":  ": ": ",
	} {
		t = strings.Replace(t, fr, to, -1)
	}
	return t
}

// ConvertRuby annotates chinese characters w/ color jyutping ruby (in the format used by github.com/luoliyan/chinese-support-redux anki plugin).
//
// Result is html.
func ConvertRuby(s string) string {
	w := strings.Builder{}
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
			w.WriteString(s[:j])
			w.WriteRune('[')
			w.WriteString(colorize(sanitizeForHtml(t)))
			w.WriteRune(']')
			s = s[j:]
		} else {
			w.WriteString(s[:i])
			s = s[i:]
		}
	}
	t := w.String()
	for fr, to := range map[string]string{
		`“ `:  `"`,
		`“`:   `"`,
		"”":   `"`,
		":  ": ": ",
	} {
		t = strings.Replace(t, fr, to, -1)
	}
	return t
}

// jat1 go3 jyut6 => <span class="tone1">jat</span> <span class="tone3">go</span> <span class="tone6">jyut</span>
func colorize(s string) string {
	ff := strings.Fields(s)
	for i := range ff {
		ff[i] = colorize1(ff[i])
	}
	return strings.Join(ff, " ")
}

// jat1 => <span class="tone1">jat</span>
func colorize1(s string) string {
	b := strings.Trim(s, "123456")
	n := s[len(b):]
	return fmt.Sprintf(`<span class="tone%s">%s</span>`, n, b)
}

func sanitizeForHtml(s string) string {
	s = strings.Join(strings.Fields(s), " ")
	s = strings.Replace(s, `"`, "&quot", -1)
	return s
}

func wantsSpace(s string) bool {
	r, _ := utf8.DecodeRuneInString(s)
	return unicode.IsLetter(r) || unicode.IsNumber(r) || r == '“'
}
