package solution

func numPairsDivisibleBy60(time []int) int {
	divisibles := make([]int, 0, 16)
	for number := 60; number < 1000; number += 60 {
		divisibles = append(divisibles, number)
	}

	timeMap := make(map[int]int, len(time))
	for _, duration := range time {
		timeMap[duration]++
	}

	count := 0
	for _, duration := range time {
		for _, divisible := range divisibles {
			number := divisible - duration
			if v, ok := timeMap[number]; ok {
				count += v
				if number == duration {
					count--
				}
			}
		}
	}

	// because the order is not matter
	return count / 2
}
