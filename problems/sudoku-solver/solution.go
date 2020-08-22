package solution

const CELL_EMPTY = '.'
const CELL_ONE = '1'
const CELL_NINE = '9'

func solveSudoku(board [][]byte) {
	solveSudokuRec(board)
}

// false if sudoku solution is incorrect
func solveSudokuRec(board [][]byte) bool {
	emptyCellsNumber, cellAvailableDigitsMap := solveSudokuSimple(board)
	// means that exist empty cells where no digits available
	if len(cellAvailableDigitsMap) != emptyCellsNumber {
		// incorrect sol
		return false
	}

	// no empty cells => all done
	if emptyCellsNumber == 0 {
		return true
	}

	// find cell with small possible values and try to gues :)
	// if solution exists => one of a digit is a true solution
	cellNumber := findMinAvailableDigitsCellNumber(cellAvailableDigitsMap)
	cellLine, cellCol := cellCoordsByNumber(cellNumber)
	// get possible digits
	cellAvailableDigits := keysToSlice(cellAvailableDigitsMap[cellNumber])
	for _, digit := range cellAvailableDigits {
		// set digit in a board
		board[cellLine][cellCol] = digit
		if solveSudokuRec(board) {
			// all done :)
			return true
		}
		// return current board state
		for k := range cellAvailableDigitsMap {
			i, j := cellCoordsByNumber(k)
			board[i][j] = CELL_EMPTY
		}
	}

	return false
}

func solveSudokuSimple(board [][]byte) (
	int,
	map[int]map[byte]bool,
) {
	areaUnavailableDigits := calculateUnavailableDigits(board)

	cellAvailableDigitsMap, emptyCellsNumber := calculateAllCellsAvaialbleDigitsMap(board, areaUnavailableDigits)

	return emptyCellsNumber, cellAvailableDigitsMap
}

// Groups three maps of unavailable digits
// for line, col, and square
// digit is unavailable if it is already used in same area
type UnavailableDigits struct {
	// index 0 to 9 contains numbers which are not available  for particular line
	lines []map[byte]bool
	// index 0 to 9 contains numbers which are not available  for particular column
	cols []map[byte]bool
	// index 0 to 9 contains numbers which are not available  for particular square
	squares []map[byte]bool
}

func calculateUnavailableDigits(board [][]byte) *UnavailableDigits {
	digits := &UnavailableDigits{
		lines:   make([]map[byte]bool, 9),
		cols:    make([]map[byte]bool, 9),
		squares: make([]map[byte]bool, 9),
	}
	for i := 0; i < 9; i++ {
		digits.lines[i] = make(map[byte]bool, 9)
		digits.cols[i] = make(map[byte]bool, 9)
		digits.squares[i] = make(map[byte]bool, 9)
	}
	for i, line := range board {
		for j, b := range line {
			if b == CELL_EMPTY {
				continue
			}
			digits.lines[i][b] = true
			digits.cols[j][b] = true
			digits.squares[squareNumberByCoords(i, j)][b] = true
		}
	}
	return digits
}

// returns map and number of empty cells
func calculateAllCellsAvaialbleDigitsMap(
	board [][]byte,
	areaUnavailableDigits *UnavailableDigits,
) (map[int]map[byte]bool, int) {
	var cellAvailableDigits map[int]map[byte]bool
	prevStepEmptyCellsNumber := -1
	emptyCellsNumber := 0
	// meand that one of the cells became not empty cell
	for emptyCellsNumber != prevStepEmptyCellsNumber {
		prevStepEmptyCellsNumber = emptyCellsNumber
		emptyCellsNumber = 0
		// important to recreate map evety step for futher correctnes check based on map len
		cellAvailableDigits = make(map[int]map[byte]bool)
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if board[i][j] != CELL_EMPTY {
					continue
				}
				availableDigitsMap := createAvailableDigitsMap(
					areaUnavailableDigits.lines[i],
					areaUnavailableDigits.cols[j],
					areaUnavailableDigits.squares[squareNumberByCoords(i, j)],
				)
				if len(availableDigitsMap) == 1 {
					// only one num available, ok we found it
					num := keysToSlice(availableDigitsMap)[0]
					// put it on board
					board[i][j] = num
					// update unavailable digits map
					areaUnavailableDigits.lines[i][num] = true
					areaUnavailableDigits.cols[j][num] = true
					areaUnavailableDigits.squares[squareNumberByCoords(i, j)][num] = true
					continue
				}
				// the cell is not filled yet
				emptyCellsNumber++
				// save
				if len(availableDigitsMap) > 0 {
					// important to no add empty maps for futher correctness check
					cellAvailableDigits[cellNumberByCoords(i, j)] = availableDigitsMap
				}
			}
		}
	}

	return cellAvailableDigits, emptyCellsNumber
}

func createAvailableDigitsMap(
	unavaialbleDigitsForLineMap map[byte]bool,
	unavaialbleDigitsForColMap map[byte]bool,
	unavaialbleDigitsForSquareMap map[byte]bool,
) map[byte]bool {
	avaialbleDigitsMap := make(map[byte]bool, 9)
	var num byte
	for num = CELL_ONE; num <= CELL_NINE; num++ {
		if _, ok := unavaialbleDigitsForLineMap[num]; ok {
			continue
		}
		if _, ok := unavaialbleDigitsForColMap[num]; ok {
			continue
		}
		if _, ok := unavaialbleDigitsForSquareMap[num]; ok {
			continue
		}
		avaialbleDigitsMap[num] = true
	}
	return avaialbleDigitsMap
}

func findMinAvailableDigitsCellNumber(cellAvailableDigitsMap map[int]map[byte]bool) int {
	cellNumber := 0
	// save first random key to num
	for k := range cellAvailableDigitsMap {
		cellNumber = k
		break
	}
	digitsNumber := len(keysToSlice(cellAvailableDigitsMap[cellNumber]))

	for k, v := range cellAvailableDigitsMap {
		dn := len(keysToSlice(v))
		if dn < digitsNumber {
			cellNumber = k
			digitsNumber = dn
		}
	}

	return cellNumber
}

func cellNumberByCoords(i, j int) int {
	return i*9 + j
}

func cellCoordsByNumber(num int) (int, int) {
	return num / 9, num % 9
}

func squareNumberByCoords(i, j int) int {
	switch {
	case i <= 2 && j <= 2:
		return 0
	case i <= 2 && j <= 5:
		return 1
	case i <= 2:
		return 2
	case i <= 5 && j <= 2:
		return 3
	case i <= 5 && j <= 5:
		return 4
	case i <= 5:
		return 5
	case j <= 2:
		return 6
	case j <= 5:
		return 7
	}
	return 8
}

func keysToSlice(m map[byte]bool) []byte {
	bytes := make([]byte, 0, len(m))
	for k := range m {
		bytes = append(bytes, k)
	}
	return bytes
}
