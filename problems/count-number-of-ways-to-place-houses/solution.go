package solution

func countHousePlacements(n int) int {
	base := 1000*1000*1000 + 7
	if n == 1 {
		return 4
	}
	if n == 2 {
		return 9
	}
	vars := make([]int, 9)
	vars[0] = 4
	vars[1] = 2
	vars[2] = 4
	vars[3] = 2
	vars[4] = 4
	vars[5] = 1
	vars[6] = 2
	vars[7] = 2
	vars[8] = 4

	for i := 3; i < n; i++ {
		vars0135 := (vars[0] + vars[1] + vars[3] + vars[5]) % base
		vars03 := (vars[0] + vars[3]) % base
		vars01 := (vars[0] + vars[1]) % base
		vars0 := vars[0]

		vars[0] = vars0135
		vars[1] = vars03
		vars[2] = vars0135
		vars[3] = vars01
		vars[4] = vars0135
		vars[5] = vars0
		vars[6] = vars03
		vars[7] = vars01
		vars[8] = vars0135

	}

	val := 0
	for i := 0; i < 9; i++ {
		val = (val + vars[i]) % base
	}
	return val
}
