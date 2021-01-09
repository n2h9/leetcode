package solution

import 	"strings"

func maximumGain(s string, x int, y int) int {
	gain := 0
	ab, ba := "ab", "ba"
	type value struct {
		gain int
		str string
	}
	moreValuable, lessValuable := value{gain: x, str: ab}, value{gain: y, str: ba}
	if y > x {
		moreValuable, lessValuable = lessValuable, moreValuable
	}
	for _, val := range []value{moreValuable, lessValuable} {
		for {
			count := strings.Count(s, val.str)
			if count <= 0 {
				break
			}
			gain += val.gain * count
			s = strings.ReplaceAll(s, val.str, "")
		}		
	}
	return gain
}

