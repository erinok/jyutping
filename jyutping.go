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

// ConvertRuby annotates chinese characters w/ color jyutping ruby (in an html table format).
// 
// Also normalizes quote types.
func ConvertRuby(s string) string {
	lines := strings.Split(s, "\n")
	for i, l := range(lines) {
		lines[i] = convertRubyLine(l)
	}
	return strings.Join(lines, "<br/>")
}

// <center><table>
// <tbody>
// <tr>
// <td><center><font size=+4>轉</font></center></td><td><center><font size=+4>職</font></center></td><td><center><font size=+4>吧</font></center></td>
// </tr><tr>
// <td><center>zyun2</center></td><td><center>zik1</center></td><td><center>baa6</center></td>
// </tr>
// </tbody></table>
// </center>
func convertRubyLine(s string) string {
	// TODO: need to collect characters in one list, ruby in another, and then join them into two table rows
	var txt []string
	var ruby []string
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
				txt = append(txt, colorizeChar(c[i], d[i]))
				ruby = append(ruby, colorizeJP1(sanitizeForHtml(d[i])))
			}
			s = s[j:]
		} else {
			txt = append(txt, sanitizeForHtml(s[:i]))
			ruby = append(ruby, "")
			s = s[i:]
		}
	}
	w := strings.Builder{}
	w.WriteString("<center><table><tbody>")
	w.WriteString("<tr>")
	for _, s := range txt {
		w.WriteString("<td><center><font size=+4>")
		w.WriteString(s)
		w.WriteString("</font></center></td>")
	}
	w.WriteString("</tr>")
	w.WriteString("<tr>")
	for _, s := range ruby {
		w.WriteString("<td><center>")
		w.WriteString(s)
		w.WriteString("</center></td>")
	}
	w.WriteString("</tr>")
	w.WriteString("</tbody></table></center>")
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
				w.WriteString(colorizeChar(c[i], d[i]))
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
		ff[i] = colorizeJP1(ff[i])
	}
	return strings.Join(ff, " ")
}

// jat1 => <span class="tone1">jat</span>
func colorizeJP1(s string) string {
	b := strings.Trim(s, "123456")
	n := s[len(b):]
	return fmt.Sprintf(`<span class="tone%s">%s</span>`, n, b)
}

// wrap s in colors per the suffix on jp
func colorizeChar(s, jp string) string {
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
