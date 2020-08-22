package solution

import (
	"sort"
)

func topKFrequent(words []string, k int) []string {
	if len(words) <= 0 || k <= 0 {
		return []string{}
	}
	appearanceMap := make(map[string]int)
	for _, word := range words {
		appearanceMap[word]++
	}
	buckets := make([][]string, len(words)+1)
	for word, appearanceNum := range appearanceMap {
		if buckets[appearanceNum] == nil {
			buckets[appearanceNum] = make([]string, 0)
		}
		buckets[appearanceNum] = append(buckets[appearanceNum], word)
	}

	res := []string{}
	bucketsIndex := len(buckets) - 1
	for k > 0 && bucketsIndex >= 0 {
		for ; bucketsIndex >= 0 && buckets[bucketsIndex] == nil; bucketsIndex-- {
		}
		if bucketsIndex < 0 || buckets[bucketsIndex] == nil {
			continue
		}
		sort.Strings(buckets[bucketsIndex])
		for i := 0; i < len(buckets[bucketsIndex]) && k > 0; i, k = i+1, k-1 {
			res = append(res, buckets[bucketsIndex][i])
		}
		bucketsIndex--
	}
	return res
}
