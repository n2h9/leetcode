package solution

func numberOfArithmeticSlices(A []int) int {
	numOfSubEndsInI, sum := 0, 0
	for i := 2; i < len(A); i++ {
		if A[i]-A[i-1] == A[i-1]-A[i-2] {
			numOfSubEndsInI++
			sum += numOfSubEndsInI
			continue
		}
		numOfSubEndsInI = 0
	}
	return sum
}
