package solution

func camelMatch(queries []string, pattern string) []bool {
	res := make([]bool, 0, len(queries))

	for _, query := range queries {
		res = append(res, match(query, pattern))
	}

	return res
}

func match(query, pattern string) bool {
	pi := 0

	for i := range query {
		if isUpper(query[i]) {
			if pi >= len(pattern) || query[i] != pattern[pi] {
				return false
			}
			pi++
			continue
		}

		if pi < len(pattern) && query[i] == pattern[pi] {
			pi++
		}
	}

	return pi >= len(pattern)
}

func isUpper(ch byte) bool {
	return ch >= 'A' && ch <= 'Z'
}
