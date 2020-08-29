package solution

func lenLongestFibSubseq(A []int) int {
	if len(A) < 3 {
		return 0
	}
	dp := make([][]*dpItem, len(A))
	max := 0

	// init dp, find all susequence of len 3
	for i := 2; i < len(A); i++ {
		for j := 1; j < i; j++ {
			for k := 0; k < j; k++ {
				if A[i] == A[j]+A[k] {
					dp[i] = append(dp[i], &dpItem{len: 3, pi: j})
					max = 3
				}
			}
		}
	}

	for i := 3; i < len(A); i++ {
		for j := 2; j < i; j++ {
			for _, item := range dp[j] {
				if A[i] == A[j]+A[item.pi] {
					newLen := item.len + 1
					dp[i] = append(dp[i], &dpItem{len: newLen, pi: j})
					if newLen > max {
						max = newLen
					}
				}
			}
		}
	}

	return max
}

type dpItem struct {
	len int // len of fib subsequence
	pi  int // index of prev element of subsequence
}
