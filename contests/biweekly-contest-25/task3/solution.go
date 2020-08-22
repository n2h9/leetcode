package solution

import (
	"sort"
)

func checkIfCanBreak(s1 string, s2 string) bool {
	b1 := ByteSlice([]byte(s1))
	b2 := ByteSlice([]byte(s2))
	sort.Sort(b1)
	sort.Sort(b2)

	res := true
	i := 0
	for ; i < len(b1) && b1[i] == b2[i]; i++ {
	}
	if i >= len(b1) {
		return res
	}
	if b1[i] >= b2[i] {
		for ; i < len(s1); i++ {
			if b1[i] < b2[i] {
				res = false
				break
			}
		}
	} else {
		for ; i < len(s1); i++ {
			if b1[i] > b2[i] {
				res = false
				break
			}
		}
	}

	return res
}

type ByteSlice []byte

func (p ByteSlice) Len() int           { return len(p) }
func (p ByteSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p ByteSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
