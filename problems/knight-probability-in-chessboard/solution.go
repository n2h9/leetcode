package solution

// optimized dp to store only last states
func knightProbability(N int, K int, r int, c int) float64 {
	probability := 1.0
	var movesPrev map[int]float64
	totalPrev := 1.0
	moves, total := possibleMoves(r, c, N)

	for i := 1; i <= K; i++ {
		p := total / (totalPrev * 8.0)
		probability *= p
		if probability <= 0 {
			break
		}
		movesPrev, totalPrev = moves, total
		moves, total = make(map[int]float64), 0
		for cellNum, cellQuantity := range movesPrev {
			i, j := coordsFromNum(cellNum, N)
			pm, t := possibleMoves(i, j, N)
			for cellNum, v := range pm {
				moves[cellNum] += v * cellQuantity
			}
			total += t * cellQuantity
		}
	}

	return probability
}

// return possible moves from i,j cel for board NxN
// if move is possible for i,j to cell i1, j1 =>
// map[i1*N+j1]++
// group same cells in one map item
// total is a total number of non unique moves
func possibleMoves(i, j, N int) (map[int]float64, float64) {
	moves := make(map[int]float64)
	total := 0.0
	if (i - 2) >= 0 {
		if j-1 >= 0 {
			moves[numFromCoords(i-2, j-1, N)]++
			total++
		}
		if j+1 < N {
			moves[numFromCoords(i-2, j+1, N)]++
			total++
		}
	}
	if (i + 2) < N {
		if j-1 >= 0 {
			moves[numFromCoords(i+2, j-1, N)]++
			total++
		}
		if j+1 < N {
			moves[numFromCoords(i+2, j+1, N)]++
			total++
		}
	}

	if (j - 2) >= 0 {
		if i-1 >= 0 {
			moves[numFromCoords(i-1, j-2, N)]++
			total++
		}
		if i+1 < N {
			moves[numFromCoords(i+1, j-2, N)]++
			total++
		}
	}
	if (j + 2) < N {
		if i-1 >= 0 {
			moves[numFromCoords(i-1, j+2, N)]++
			total++
		}
		if i+1 < N {
			moves[numFromCoords(i+1, j+2, N)]++
			total++
		}
	}

	return moves, total
}

func numFromCoords(i, j, N int) int {
	return i*N + j
}

func coordsFromNum(l, N int) (int, int) {
	return l / N, l % N
}
