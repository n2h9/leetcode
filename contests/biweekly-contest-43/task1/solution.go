package solution

func totalMoney(n int) int {
	total := 0
	weekStartAmount := 1
	dayAmount := weekStartAmount
	dayNumber := 1
	for ;n > 0; n-- {
		total += dayAmount
		dayAmount++
		dayNumber++
		if dayNumber > 7 {
			dayNumber = 1
			weekStartAmount++
			dayAmount = weekStartAmount
		}
	}
	return total
}
