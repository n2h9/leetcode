package solution

func checkValidGrid(grid [][]int) bool {
	if grid[0][0] != 0 {
		return false
	}

	n := len(grid)
	movesNum := n*n - 1
	var valid func(int, int, int) bool
	valid = func(num, i, j int) bool {
		if num > movesNum {
			return true
		}
		possibleMoves := [][]int{
			{i - 2, j - 1},
			{i - 1, j - 2},
			{i + 1, j - 2},
			{i + 2, j - 1},
			{i + 2, j + 1},
			{i + 1, j + 2},
			{i - 1, j + 2},
			{i - 2, j + 1},
		}
		var validMove []int
		for _, move := range possibleMoves {
			if !isExistingCell(n, move[0], move[1]) {
				continue
			}
			if grid[move[0]][move[1]] == num {
				validMove = move
				break
			}
		}

		if validMove == nil {
			return false
		}
		return valid(num+1, validMove[0], validMove[1])
	}

	return valid(1, 0, 0)
}

func isExistingCell(n, i, j int) bool {
	return i >= 0 && j >= 0 && i < n && j < n
}
