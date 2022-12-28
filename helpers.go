package wgconfig

import (
	"strings"
)

func trimComment(s string) string {
	if len(s) > 0 && (s[0] == ';' || s[0] == '#') {
		s = s[1:]
	}
	return strings.TrimLeft(s, " ")
}
