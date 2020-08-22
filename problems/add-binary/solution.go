package solution

const ZERO = '0'
const UNIT = '1'

func addBinary(a string, b string) string {
	result := ""
	var overflow byte = 0

	var i, j int
	for i, j = len(a)-1, len(b)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		// if i or j is less than zero so string is finished and take 0 as a value
		var da, db byte
		if i >= 0 {
			da = a[i] - ZERO
		}
		if j >= 0 {
			db = b[j] - ZERO
		}
		sum := da + db + overflow
		if sum < 2 {
			// sum is 0 or 1, so nothing to cary
			overflow = 0
		} else {
			// sum is 10 or 11
			// save overflow
			overflow = 1
			// take low bit
			sum &= 1
		}
		result = string(sum+ZERO) + result
	}
	if overflow > 0 {
		result = string(UNIT) + result
	}

	return result
}
