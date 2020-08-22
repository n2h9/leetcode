package solution

func maxScore(cardPoints []int, k int) int {
	return calculate(cardPoints, 0, len(cardPoints)-1, 0, k, make(map[int]map[int]int, len(cardPoints)))
}

func calculate(cardPoints []int, leftBorderIndex, rightBorderIndex, total, iterNum int, accumulator map[int]map[int]int) int {
	if iterNum <= 0 || leftBorderIndex > rightBorderIndex {
		return total
	}
	if v, ok := accumulator[leftBorderIndex]; ok {
		if v2, ok := v[rightBorderIndex]; ok {
			return v2
		}
	} else {
		accumulator[leftBorderIndex] = make(map[int]int, 1000)
	}
	takeLeftCardSum := calculate(cardPoints, leftBorderIndex+1, rightBorderIndex, cardPoints[leftBorderIndex]+total, iterNum-1, accumulator)
	takeRightCardSum := 0
	if leftBorderIndex < rightBorderIndex {
		takeRightCardSum = calculate(cardPoints, leftBorderIndex, rightBorderIndex-1, cardPoints[rightBorderIndex]+total, iterNum-1, accumulator)
	}

	result := takeLeftCardSum
	if takeRightCardSum > result {
		result = takeRightCardSum
	}
	accumulator[leftBorderIndex][rightBorderIndex] = result
	return result
}
