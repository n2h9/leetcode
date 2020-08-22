package solution

func numMagicSquaresInside(grid [][]int) int {
	res := 0
	for i := 0; i <= len(grid)-3; i++ {
		for j := 0; j <= len(grid[i])-3; j++ {
			if isMagic(grid, i, j) {
				res++
			}
		}
	}
	return res
}

func isMagic(grid [][]int, i, j int) bool {
	distinctMap := make(map[int]bool, 9)
	for ii := i; ii < i+3; ii++ {
		for jj := j; jj < j+3; jj++ {
			if grid[ii][jj] < 1 || grid[ii][jj] > 9 {
				return false
			}
			distinctMap[grid[ii][jj]] = true
		}
	}

	if len(distinctMap) != 9 {
		return false
	}

	row0 := grid[i][j] + grid[i][j+1] + grid[i][j+2]
	row1 := grid[i+1][j] + grid[i+1][j+1] + grid[i+1][j+2]
	row2 := grid[i+2][j] + grid[i+2][j+1] + grid[i+2][j+2]

	col0 := grid[i][j] + grid[i+1][j] + grid[i+2][j]
	col1 := grid[i][j+1] + grid[i+1][j+1] + grid[i+2][j+1]
	col2 := grid[i][j+2] + grid[i+1][j+2] + grid[i+2][j+2]

	d1 := grid[i][j] + grid[i+1][j+1] + grid[i+2][j+2]
	d2 := grid[i+2][j] + grid[i+1][j+1] + grid[i][j+2]
	return d2 == d1 &&
		col2 == d1 &&
		col1 == d1 &&
		col0 == d1 &&
		row2 == d1 &&
		row1 == d1 &&
		row0 == d1
}
