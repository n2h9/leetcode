package solution

func longestOnes(A []int, K int) int {
	longest := 0
	k := K
	for h, l := 0, 0; h < len(A); {
		for ; h < len(A) && (A[h] == 1 || k > 0); h++ {
			if A[h] == 0 {
				k--
			}
		}
		if val := h - l; val > longest {
			longest = val
		}
		for ; l < h && A[l] != 0; l++ {
		}
		if K > 0 {
			k++
		}
		l++
		if l > h {
			h++
		}
	}

	return longest
}
