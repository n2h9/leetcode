package solution

func letterCombinations(digits string) []string {
	if len(digits) <= 0 {
		return []string{}
	}
	dlm := map[byte][]byte{
		'2': []byte{'a', 'b', 'c',},
		'3': []byte{'d', 'e', 'f',},
		'4': []byte{'g', 'h', 'i',},
		'5': []byte{'j', 'k', 'l',},
		'6': []byte{'m', 'n', 'o',},
		'7': []byte{'p', 'q', 'r', 's',},
		'8': []byte{'t', 'u', 'v',},
		'9': []byte{'w', 'x', 'y', 'z',},
	}
	res := make([]string, 0, 4)
	for _, l := range dlm[digits[0]] {
		res = append(res, string(l))
	}
	for i := 1; i < len(digits); i++ {
		resPrev := res
		res = make([]string, 0)
		for _, s := range resPrev {
			for _, l := range dlm[digits[i]] {
				res = append(res, s+string(l))
			}
		}
	}
    return res
}
