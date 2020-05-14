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
			w.WriteString(colorizeJP(sanitizeForHtml(t)))
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

// ColorizeChars wraps chinese characters in <span class="toneX">, per the tones.
func ColorizeChars(s string) string {
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
			d := strings.Fields(t)
			c := strings.Split(s[:j], "")
			if len(d) > len(c) {
				panic(fmt.Sprintln("programmer error", d, c))
			} 
			for i := range d {
				w.WriteString(colorize2(c[i], d[i]))
			}
			s = s[j:]
		} else {
			w.WriteString(s[:i])
			s = s[i:]
		}
	}
	return w.String()
}

// jat1 go3 jyut6 => <span class="tone1">jat</span> <span class="tone3">go</span> <span class="tone6">jyut</span>
func colorizeJP(s string) string {
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

// wrap s in colors per the suffix on jp
func colorize2(s, jp string) string {
	b := strings.Trim(jp, "123456")
	n := jp[len(b):]
	if n == "" {
		return s
	}
	return fmt.Sprintf(`<span class="tone%s">%s</span>`, n, s)
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
