package solution

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	line := make([]int, n)
	line[0] = ^obstacleGrid[0][0] & 1
	for i := 1; i < n; i++ {
		if obstacleGrid[0][i] != 1 {
			line[i] = line[i-1]
		}
	}

	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			line[0] = 0
		}
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				line[j] = 0
				continue
			}
			line[j] += line[j-1]
		}
	}

	return line[n-1]
}
