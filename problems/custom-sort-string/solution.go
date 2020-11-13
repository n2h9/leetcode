package solution

func customSortString(S string, T string) string {
	ls, lt := len(S), len(T)
	countCharsT := make(map[byte]byte, lt)
	for i := range T {
		countCharsT[T[i]]++
	}

	existCharS := make(map[byte]bool, ls)
	result := ""
	for i := range S {
		ch := S[i]
		existCharS[ch] = true
		var j byte
		for ; j < countCharsT[ch]; j++ {
			result += string(ch)
		}
	}

	for i := range T {
		ch := T[i]
		if existCharS[ch] {
			continue
		}
		result += string(ch)
	}

	return result
}
