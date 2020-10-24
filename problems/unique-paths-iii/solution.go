package solution

const EMPTY_CELL = 0
const START_CELL = 1
const END_CELL = 2
const PROCESSING_CELL = 3

func uniquePathsIII(grid [][]int) int {
	emptyCellNum := 0
	endI, endJ := 0, 0
	for i := len(grid) - 1; i >= 0; i-- {
		for j := len(grid[i]) - 1; j >= 0; j-- {
			if grid[i][j] == END_CELL {
				endI, endJ = i, j
				continue
			}
			if grid[i][j] == EMPTY_CELL {
				emptyCellNum++
			}
		}
	}
	return numberOfPaths(grid, endI, endJ-1, emptyCellNum) +
		numberOfPaths(grid, endI, endJ+1, emptyCellNum) +
		numberOfPaths(grid, endI-1, endJ, emptyCellNum) +
		numberOfPaths(grid, endI+1, endJ, emptyCellNum)
}

func numberOfPaths(grid [][]int, i, j, emptyCellsNum int) int {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return 0
	}
	if grid[i][j] == START_CELL && emptyCellsNum == 0 {
		return 1
	}
	if emptyCellsNum <= 0 || grid[i][j] != EMPTY_CELL {
		return 0
	}
	grid[i][j] = PROCESSING_CELL
	emptyCellsNum--
	res := numberOfPaths(grid, i, j-1, emptyCellsNum)
	res += numberOfPaths(grid, i, j+1, emptyCellsNum)
	res += numberOfPaths(grid, i-1, j, emptyCellsNum)
	res += numberOfPaths(grid, i+1, j, emptyCellsNum)
	grid[i][j] = EMPTY_CELL
	return res
}
