package solution

func leastInterval(tasks []byte, n int) int {
	occuranceMap := make(map[byte]int, 26)
	for _, task := range tasks {
		occuranceMap[task]++
	}

	result := 0
	sortedTasksList := sort(occuranceMap)
	for len(occuranceMap) > 0 {
		t := 0
		for i := 0; i < len(sortedTasksList) && t <= n; i++ {
			// execute task
			occuranceMap[sortedTasksList[i]]--
			t++
			if occuranceMap[sortedTasksList[i]] <= 0 {
				// delete from map
				delete(occuranceMap, sortedTasksList[i])
			}
		}
		if len(occuranceMap) > 0 {
			result += n + 1
			sortedTasksList = sort(occuranceMap)
		} else {
			result += t
		}
	}

	return result
}

func sort(occuranceMap map[byte]int) []byte {
	sortedTasksList := make([]byte, 0, 26)
	occuranceMapForSort := make(map[byte]int, 26)
	for k, v := range occuranceMap {
		occuranceMapForSort[k] = v
	}
	for len(occuranceMapForSort) > 0 {
		maxCount := 0
		var maxTask byte = 0
		for task, count := range occuranceMapForSort {
			if count > maxCount {
				maxCount = count
				maxTask = task
			}
		}
		sortedTasksList = append(sortedTasksList, maxTask)
		delete(occuranceMapForSort, maxTask)
	}
	return sortedTasksList
}
