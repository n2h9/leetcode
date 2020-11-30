package solution

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	maxLen := 0
	startIndex := 0
	charOccurrence := make(map[byte]int, n)

	for i := 0; i < n; i++ {
		if j, ok := charOccurrence[s[i]]; ok {
			if v := i-startIndex; v > maxLen {
				maxLen = v
			}
			i = j+1
			startIndex = i
			charOccurrence = make(map[byte]int, n)
		}
		charOccurrence[s[i]] = i
	}

	if v := n-startIndex; v > maxLen {
		maxLen = v
	}

	return maxLen
}