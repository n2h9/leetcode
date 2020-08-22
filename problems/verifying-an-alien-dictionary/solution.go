package solution

func isAlienSorted(words []string, order string) bool {
	orderWeight := make(map[byte]byte, 26)
	var i byte
	for i = 0; i < 26; i++ {
		orderWeight[order[i]] = i
	}
	for i := 1; i < len(words); i++ {
		if !leq(words[i-1], words[i], orderWeight) {
			return false
		}
	}
	return true
}

func leq(a, b string, orderWeight map[byte]byte) bool {
	i := 0
	for ; i < len(a) && i < len(b); i++ {
		if orderWeight[a[i]] < orderWeight[b[i]] {
			return true
		}
		if orderWeight[a[i]] > orderWeight[b[i]] {
			return false
		}
	}

	return i >= len(a)
}
