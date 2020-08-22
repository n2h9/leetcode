package solution

func shortestBridge(A [][]int) int {
	i, j := 0, 0
loop:
	for i = range A {
		for j = range A[i] {
			if A[i][j] == 0 {
				continue
			}
			break loop
		}
	}

	dsf(A, i, j)
	return bsf(A, 2) - 2
}

// find all island cells and mark them with 2
func dsf(A [][]int, i, j int) {
	if i < 0 || j < 0 || i >= len(A) || j >= len(A[i]) || A[i][j] != 1 {
		return
	}
	A[i][j] = 2

	dsf(A, i+1, j)
	dsf(A, i-1, j)
	dsf(A, i, j+1)
	dsf(A, i, j-1)
}

// process search by layer
func bsf(A [][]int, layer int) int {
	nextLayer := layer + 1
	for i := range A {
		for j := range A[i] {
			if A[i][j] != layer {
				continue
			}
			if (i + 1) < len(A) {
				if A[i+1][j] == 1 {
					return layer
				}
				if A[i+1][j] == 0 {
					A[i+1][j] = nextLayer
				}
			}
			if i > 0 {
				if A[i-1][j] == 1 {
					return layer
				}
				if A[i-1][j] == 0 {
					A[i-1][j] = nextLayer
				}
			}
			if j+1 < len(A[i]) {
				if A[i][j+1] == 1 {
					return layer
				}
				if A[i][j+1] == 0 {
					A[i][j+1] = nextLayer
				}
			}
			if j > 0 {
				if A[i][j-1] == 1 {
					return layer
				}
				if A[i][j-1] == 0 {
					A[i][j-1] = nextLayer
				}
			}
		}
	}
	return bsf(A, nextLayer)
}
