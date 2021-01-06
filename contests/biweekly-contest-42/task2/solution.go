package solution

func averageWaitingTime(customers [][]int) float64 {
	time, waitingTime := 0, 0
	for _, customer := range customers {
		if customer[0] > time {
			time = customer[0]
		}
		time += customer[1]
		waitingTime += time - customer[0]
	}
	return float64(waitingTime) / float64(len(customers))
}

