package solution

func solveNQueens(n int) [][]string {
	if n <= 0 {
		return [][]string{[]string{}}
	}

	const MAX_INT = 1<<31 - 1
	// hust because by default byte array is filled by zero values use it as empty char
	const CHAR_EMPTY byte = 0
	// and replace to this char before return result
	const CHAR_RESULT_EMPTY byte = '.'
	const CHAR_QUEEN byte = 'Q'
	const CHAR_ATACK byte = 'X'

	// board will stay valid if a Queen will be putted in cell y (vertical) x (horizontal)
	boardStayValid := func(board [][]byte, y, x int) bool {
		return board[y][x] == CHAR_EMPTY
	}

	// y - vertical coord from top = 0 to bottom = n -1
	// x - horizontatl coord from left = 0 to right = n - 1
	boardAddQueen := func(board [][]byte, y, x int) {
		board[y][x] = CHAR_QUEEN
		for i := 0; i < n; i++ {
			if i != x {
				// mark horizontal cells as cell by ATTACK
				board[y][i] = CHAR_ATACK
			}
			if i != y {
				// mark vertical cells as cell by ATTACK
				board[i][x] = CHAR_ATACK
			}
		}
		// diagonal 1: from current cell to left bottom
		for y2, x2 := y+1, x-1; y2 < n && x2 >= 0; y2, x2 = y2+1, x2-1 {
			board[y2][x2] = CHAR_ATACK
		}
		// diagonal 1: from current cell to right top
		for y2, x2 := y-1, x+1; y2 >= 0 && x2 < n; y2, x2 = y2-1, x2+1 {
			board[y2][x2] = CHAR_ATACK
		}
		// diagonal 2: from current cell to left top
		for y2, x2 := y-1, x-1; y2 >= 0 && x2 >= 0; y2, x2 = y2-1, x2-1 {
			board[y2][x2] = CHAR_ATACK
		}
		// diagonal 1: from current cell to right bottom
		for y2, x2 := y+1, x+1; y2 < n && x2 < n; y2, x2 = y2+1, x2+1 {
			board[y2][x2] = CHAR_ATACK
		}
	}

	// replace all non QUEEN chars with EMPTY result chars
	// returns "hash" of board
	boardCleanCells := func(board [][]byte) {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if board[i][j] != CHAR_QUEEN {
					board[i][j] = CHAR_RESULT_EMPTY
				}
			}
		}
	}

	boardCopy := func(board [][]byte) [][]byte {
		newBoard := make([][]byte, n)
		for i := 0; i < n; i++ {
			newBoard[i] = make([]byte, n)
			copy(newBoard[i], board[i])
		}
		return newBoard
	}

	bytesToStringSlice := func(b [][]byte) []string {
		slice := make([]string, n)
		for i := 0; i < n; i++ {
			slice[i] = string(b[i])
		}
		return slice
	}

	result := [][]string{}

	var processLine func(y int, board [][]byte)
	processLine = func(y int, board [][]byte) {
		if y >= n {
			// means we achieve one of the solutions solution
			boardCleanCells(board)
			result = append(result, bytesToStringSlice(board))
			return
		}

		for x := 0; x < n; x++ {
			if !boardStayValid(board, y, x) {
				// invalid queen placement => discard solution
				continue
			}
			currentBoard := boardCopy(board)
			boardAddQueen(currentBoard, y, x)
			processLine(y+1, currentBoard)
		}
	}

	currentBoard := make([][]byte, n)
	for i := 0; i < n; i++ {
		currentBoard[i] = make([]byte, n)
	}
	processLine(0, currentBoard)

	return result
}
