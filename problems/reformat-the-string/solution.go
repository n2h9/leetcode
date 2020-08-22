package solution

func reformat(s string) string {
	charSeq1, charSeq2 := split(s)

	if (len(charSeq1) - len(charSeq2)) > 1 {
		return ""
	}

	res := ""
	i := 0
	for ; i < len(charSeq2); i++ {
		res += string(charSeq1[i]) + string(charSeq2[i])
	}
	if i < len(charSeq1) {
		res += string(charSeq1[i])
	}

	return res
}

func split(s string) ([]byte, []byte) {
	digits := make([]byte, 0, len(s))
	letters := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			digits = append(digits, s[i])
			continue
		}
		letters = append(letters, s[i])
	}

	if len(letters) > len(digits) {
		return letters, digits
	}

	return digits, letters
}
