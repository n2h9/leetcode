package solution

func orderOfLargestPlusSign(N int, mines [][]int) int {
	m := len(mines)
	minesMap := make(map[int]map[int]bool, m)
	for _, mine := range mines {
		if minesMap[mine[0]+1] == nil {
			minesMap[mine[0]+1] = make(map[int]bool)
		}
		minesMap[mine[0]+1][mine[1]+1] = true
	}

	left := make([][]int, N+1)
	right := make([][]int, N+2)
	up := make([][]int, N+1)
	down := make([][]int, N+2)

	left[0] = make([]int, N+1)
	up[0] = make([]int, N+1)
	for i := 1; i <= N; i++ {
		left[i] = make([]int, N+1)
		up[i] = make([]int, N+1)
		minesMapCurr := minesMap[i]
		for j := 1; j <= N; j++ {
			if minesMapCurr != nil && minesMapCurr[j] {
				// this is  a mine cell
				continue
			}
			left[i][j] = 1 + left[i][j-1]
			up[i][j] = 1 + up[i-1][j]
		}
	}

	right[N+1] = make([]int, N+2)
	down[N+1] = make([]int, N+2)
	for i := N; i >= 1; i-- {
		right[i] = make([]int, N+2)
		down[i] = make([]int, N+2)
		minesMapCurr := minesMap[i]
		for j := N; j >= 1; j-- {
			if minesMapCurr != nil && minesMapCurr[j] {
				// this is  a mine cell
				continue
			}
			right[i][j] = 1 + right[i][j+1]
			down[i][j] = 1 + down[i+1][j]
		}
	}

	maxVal := 0

	for i := 1; i <= N; i++ {
		minesMapCurr := minesMap[i]
		for j := 1; j <= N; j++ {
			if minesMapCurr != nil && minesMapCurr[j] {
				// this is  a mine cell
				continue
			}
			possibleMaxVal := 1 + min(
				left[i][j-1],
				right[i][j+1],
				up[i-1][j],
				down[i+1][j],
			)
			if possibleMaxVal > maxVal {
				maxVal = possibleMaxVal
			}
		}
	}

	return maxVal
}

func min(a ...int) int {
	m := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] < m {
			m = a[i]
		}
	}
	return m
}
