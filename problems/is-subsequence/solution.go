package solution

func isSubsequence(s string, t string) bool {
	ls := len(s)
	if ls <= 0 {
		return true
	}

	lt := len(t)
	if ls > lt {
		return false
	}

	i, j := 0, 0
	for ; i < ls; i++ {
		for ; j < lt && s[i] != t[j]; j++ {
		}
		if j >= lt {
			break
		}
		j++
	}

	return i == ls
}
