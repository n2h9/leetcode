package solution

import (
	"strings"
)

func maxScore(s string) int {
	max := 0

	n := len(s)
	runes := []rune(s)
	for i := 1; i < n; i++ {
		l := strings.ReplaceAll(
			string(runes[0:i]),
			"1",
			"",
		)
		r := strings.ReplaceAll(
			string(runes[i:n]),
			"0",
			"",
		)
		m := len(l) + len(r)
		if m > max {
			max = m
		}
	}

	return max
}
