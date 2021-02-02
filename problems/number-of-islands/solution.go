package solution

func numIslands(grid [][]byte) int {
	var (
		islandMark          byte = '1'
		visitedIslandMark   byte = 'x'
		count               int  = 0
		markIslandAsVisited func(g [][]byte, i, j int)
	)

	markIslandAsVisited = func(g [][]byte, i, j int) {
		if i >= len(g) || i < 0 || j >= len(g[i]) || j < 0 {
			return
		}
		if g[i][j] != islandMark {
			return
		}
		g[i][j] = visitedIslandMark

		coords := [][]int{
			[]int{i - 1, j},
			[]int{i + 1, j},
			[]int{i, j - 1},
			[]int{i, j + 1},
		}
		for _, coord := range coords {
			markIslandAsVisited(g, coord[0], coord[1])
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != islandMark {
				continue
			}
			count++
			markIslandAsVisited(grid, i, j)
		}
	}

	return count
}
