package solution

func maxDiff(num int) int {
	dsMax := digits(num)
	dsMin := make([]int, len(dsMax))
	copy(dsMin, dsMax)

	maxValToChange := dsMax[len(dsMax)-1]
	for i := len(dsMax) - 1; i >= 0; i-- {
		if dsMax[i] == 9 {
			continue
		}
		maxValToChange = dsMax[i]
		break
	}
	for i := len(dsMax) - 1; i >= 0; i-- {
		if dsMax[i] == maxValToChange {
			dsMax[i] = 9
		}
	}

	changeVal := 1
	minValToChange := dsMin[len(dsMin)-1]
	if dsMin[len(dsMin)-1] == 1 && len(dsMin) > 1 {
		changeVal = 0
		minValToChange = -1
		for i := len(dsMin) - 2; i >= 0; i-- {
			if dsMin[i] == 0 || dsMin[i] == 1 {
				continue
			}
			minValToChange = dsMin[i]
			break
		}
	}
	if minValToChange != -1 {
		for i := len(dsMin) - 1; i >= 0; i-- {
			if dsMin[i] == minValToChange {
				dsMin[i] = changeVal
			}
		}
	}

	return combine(dsMax) - combine(dsMin)
}

func digits(num int) []int {
	val := make([]int, 0, 9)
	for num > 0 {
		val = append(val, num%10)
		num /= 10
	}
	return val
}

func combine(digits []int) int {
	num := digits[0]
	power := 10
	for i := 1; i < len(digits); i++ {
		num += digits[i] * power
		power *= 10
	}
	return num
}
