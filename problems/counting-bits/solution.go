package solution

func countBits(num int) []int {
	if num <= 0 {
		return []int{0}
	}
	bitsNum := make([]int, num+1)
	bitsNum[0] = 0
	bitsNum[1] = 1

	for x := 2; x <= num; x++ {
		lowBit := x & 1
		bitsNum[x] = bitsNum[x>>1] + lowBit
	}
	return bitsNum
}
