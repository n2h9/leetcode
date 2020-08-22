package solution

func countSubstrings(s string) int {
	n := len(s)
	if n <= 0 {
		return 0
	}

	// palindromMap[i*n+j] == true => subs[i,j] is palindrom
	palindromMap := make(map[int]bool)
	// oneCharPalindromMap[i*n+j] == true => subs[i,j] is palindrom of one char: f.e. aaaa
	oneCharPalindromMap := make(map[int]bool)

	for i := 0; i < n; i++ {
		palindromMap[i*n+i] = true
		oneCharPalindromMap[i*n+i] = true
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n-k; i++ {
			j := i + k
			if !palindromMap[i*n+j] {
				continue
			}
			if oneCharPalindromMap[i*n+j] {
				// eq chars also check that index is correct
				if eqChars(s, i-1, j) {
					palindromMap[(i-1)*n+j] = true
					oneCharPalindromMap[(i-1)*n+j] = true
				}
				// eq chars also check that index is correct
				if eqChars(s, i, j+1) {
					palindromMap[i*n+j+1] = true
					oneCharPalindromMap[i*n+j+1] = true
				}
				if eqChars(s, i-1, j+1) {
					palindromMap[(i-1)*n+j+1] = true
					// AXXXA is not one char palindrom
					// but
					// XXXXX is one char palindrom
					// thats why additional check eqChars(s, i-1, i)
					if eqChars(s, i-1, i) {
						oneCharPalindromMap[(i-1)*n+j+1] = true
					}
				}
				continue
			}
			if eqChars(s, i-1, j+1) {
				palindromMap[(i-1)*n+j+1] = true
			}
		}
	}

	return len(palindromMap)
}

// compare chars
// if index out of bound => false
func eqChars(s string, i, j int) bool {
	if i < 0 || j < 0 {
		return false
	}
	if i >= len(s) || j >= len(s) {
		return false
	}

	return s[i] == s[j]
}
