package solution

func maximumBinaryString(binary string) string {
	res := make([]byte, 0, len(binary))
	prevZeroIndex := len(binary)
	res = append(res, binary[0])
	if res[0] == '0' {
		prevZeroIndex = 0
	}
	for i := 1; i < len(binary); i++ {
		if binary[i] == '1' {
			// if '1' just add
			res = append(res, '1')
			continue
		} // else => current char is '0'
		
		if res[len(res)-1] == '0' {
			// if last char was also zero make '00 to '10' switch
			res[len(res)-1] = '1'
			res = append(res, '0')
			prevZeroIndex = i
			continue
		}
		if prevZeroIndex > i {
			// if this is the first zero
			res = append(res, '0')
			prevZeroIndex = i
			continue
		} // else there is a zero somewhere
		
		res = append(res, '0')
		for j := len(res)-2; j >= 0 && res[j] == '1'; j-- {
			res[j] = '0'
			res[j+1] = '1'
		}
		res[prevZeroIndex] = '1'
		prevZeroIndex++
	}
	return string(res)
}

