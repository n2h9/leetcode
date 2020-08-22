package solution

func matrixBlockSum(mat [][]int, K int) [][]int {
	// m and n >= 1 by condition, do not check it
	m := len(mat)
	n := len(mat[0])
	res := make([][]int, m)

	// calculate (0,0) value
	res[0] = make([]int, n)
	for i := 0; i < m && i <= K; i++ {
		for j := 0; j < n && j <= K; j++ {
			res[0][0] += mat[i][j]
		}
	}
	// --

	// calculate vertical (i,0) i=0,m-1;
	// ths - is top horizontal sum, used to caluclate sum of cur cell bassed on sum of prev cell
	ths := 0
	for i := 1; i < m; i++ {
		res[i] = make([]int, n)
		res[i][0] = res[i-1][0] + sumHor(mat, i+K, 0, K) - ths
		if i < m-1 {
			ths = sumHor(mat, i-K, 0, K)
		}
	}
	//--

	// calculate other cells
	// lvs - is left vertical sum, used to caluclate sum of cur cell bassed on sum of prev cell
	//--
	for i := 0; i < m; i++ {
		lvs := 0
		for j := 1; j < n; j++ {
			res[i][j] = res[i][j-1] + sumVer(mat, j+K, i-K, i+K) - lvs
			if j < n-1 {
				lvs = sumVer(mat, j-K, i-K, i+K)
			}
		}
	}

	return res
}

func sumHor(a [][]int, vi, start, end int) int {
	if vi < 0 || vi >= len(a) {
		return 0
	}
	if end >= len(a[vi]) {
		end = len(a[vi]) - 1
	}
	s := 0
	for i := start; i <= end; i++ {
		s += a[vi][i]
	}
	return s
}

func sumVer(mat [][]int, hi, start, end int) int {
	if hi < 0 || hi >= len(mat[0]) {
		return 0
	}
	if start < 0 {
		start = 0
	}
	if end >= len(mat) {
		end = len(mat) - 1
	}
	s := 0
	for i := start; i <= end; i++ {
		s += mat[i][hi]
	}
	return s
}
