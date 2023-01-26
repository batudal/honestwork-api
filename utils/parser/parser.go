package parser

import "strings"

func ParseContent(content string) string {
	ps := strings.Split(content, "<p>")
	h4s := strings.Split(content, "<h4>")

	for i, p := range ps {
		if strings.Contains(p, "</p>") {
			ps[i] = strings.Split(ps[i], "</p>")[0]
		}
		ps[i] = strings.ReplaceAll(p, "<em>", "")
		ps[i] = strings.ReplaceAll(p, "</em>", "")
		ps[i] = strings.ReplaceAll(p, "<strong>", "")
		ps[i] = strings.ReplaceAll(p, "</strong>", "")
	}
	for j, h4 := range h4s {
		if strings.Contains(h4, "</h4>") {
			h4s[j] = strings.Split(h4s[j], "</h4>")[0]
		}
		h4s[j] = strings.ReplaceAll(h4, "<em>", "")
		h4s[j] = strings.ReplaceAll(h4, "</em>", "")
		h4s[j] = strings.ReplaceAll(h4, "<strong>", "")
		h4s[j] = strings.ReplaceAll(h4, "</strong>", "")
	}
	var flat_ps string
	var flat_h4s string
	for _, h4 := range h4s {
		flat_h4s += h4
	}
	for _, p := range ps {
		flat_ps += p
	}
	return flat_h4s + flat_ps
}
