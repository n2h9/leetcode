package solution

func countStudents(students []int, sandwiches []int) int {
	studTopIndex, sandwichTopIndex := 0, 0
	lastFedStudIndex := -1
	studLeft := len(students)

	for sandwichTopIndex < len(sandwiches) && studLeft > 0 && lastFedStudIndex != studTopIndex {
		if students[studTopIndex] == sandwiches[sandwichTopIndex] {
			studLeft--
			lastFedStudIndex = studTopIndex
			students[lastFedStudIndex] = -1
			sandwichTopIndex++
		}
		studTopIndex++
		if studTopIndex >= len(students) {
			// if all students pass and nobody was fed
			if lastFedStudIndex == -1 {
				break
			}
			studTopIndex = 0
		}
	}

	return studLeft
}

