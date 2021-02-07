package solution

func minimumLength(s string) int {
	var truncate func(b []byte) []byte
	truncate = func(b []byte) []byte {
		if len(b) <= 1 || b[0] != b[len(b)-1] {
			return b
		}

		start := 0
		for ; start < len(b) && b[start] == b[0]; start++ {
		}

		end := len(b) - 1
		for ; end >= start && b[end] == b[0]; end-- {
		}

		return truncate(b[start : end+1])
	}

	return len(truncate([]byte(s)))
}
