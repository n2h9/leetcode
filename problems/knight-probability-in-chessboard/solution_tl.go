package solution

// // DP which stores all states
// func knightProbabilityDPSimple(N int, K int, r int, c int) float64 {
// 	probabilities := make([]float64, K+1)
// 	moves := make([][][]int, K+2)

// 	probabilities[0] = 1
// 	moves[0] = [][]int{[]int{r, c}}
// 	moves[1] = possibleMoves(r, c, N)

// 	for i := 1; i <= K; i++ {
// 		p := float64(len(moves[i])) / (float64(len(moves[i-1])) * 8.0)
// 		probabilities[i] = probabilities[i-1] * p
// 		if probabilities[i] <= 0 {
// 			return probabilities[i]
// 		}
// 		moves[i+1] = make([][]int, 0)
// 		for _, v := range moves[i] {
// 			moves[i+1] = append(moves[i+1], possibleMoves(v[0], v[1], N)...)
// 		}
// 	}

// 	return probabilities[K]
// }

// func possibleMoves(i, j, N int) [][]int {
// 	moves := make([][]int, 0, 8)
// 	if (i - 2) >= 0 {
// 		if j-1 >= 0 {
// 			moves = append(moves, []int{i - 2, j - 1})
// 		}
// 		if j+1 < N {
// 			moves = append(moves, []int{i - 2, j + 1})
// 		}
// 	}
// 	if (i + 2) < N {
// 		if j-1 >= 0 {
// 			moves = append(moves, []int{i + 2, j - 1})
// 		}
// 		if j+1 < N {
// 			moves = append(moves, []int{i + 2, j + 1})
// 		}
// 	}

// 	if (j - 2) >= 0 {
// 		if i-1 >= 0 {
// 			moves = append(moves, []int{i - 1, j - 2})
// 		}
// 		if i+1 < N {
// 			moves = append(moves, []int{i + 1, j - 2})
// 		}
// 	}
// 	if (j + 2) < N {
// 		if i-1 >= 0 {
// 			moves = append(moves, []int{i - 1, j + 2})
// 		}
// 		if i+1 < N {
// 			moves = append(moves, []int{i + 1, j + 2})
// 		}
// 	}

// 	return moves
// }

// // optimized dp to store only last states
// func knightProbability(N int, K int, r int, c int) float64 {
// 	probability := 1.0
// 	movesPrev := [][]int{[]int{r, c}}
// 	moves := possibleMoves(r, c, N)

// 	for i := 1; i <= K; i++ {
// 		p := float64(len(moves)) / (float64(len(movesPrev)) * 8.0)
// 		probability *= p
// 		if probability <= 0 {
// 			break
// 		}
// 		movesPrev = moves
// 		moves = make([][]int, 0)
// 		for _, v := range movesPrev {
// 			moves = append(moves, possibleMoves(v[0], v[1], N)...)
// 		}
// 	}

// 	return probability
// }
