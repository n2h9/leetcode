package solution

func wordBreak(s string, wordDict []string) bool {
	if len(s) <= 0 || len(wordDict) <= 0 {
		return false
	}
	dp := make([]bool, len(s))
	for i := 0; i < len(s); i++ {
		for _, w := range wordDict {
			l := len(w)
			if (i - l) < -1 {
				continue
			}
			if i < l {
				if s[i-l+1:i+1] == w {
					dp[i] = true
					break
				}
				continue
			}
			if dp[i-l] && s[i-l+1:i+1] == w {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)-1]
}
