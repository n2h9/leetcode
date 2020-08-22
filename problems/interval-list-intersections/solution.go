package solution

func intervalIntersection(A [][]int, B [][]int) [][]int {
	if len(A) <= 0 || len(B) <= 0 {
		return [][]int{}
	}

	res := make([][]int, 0)
	for ai, bi := 0, 0; ai < len(A) && bi < len(B); {
		if intersection := getIntersection(A[ai], B[bi]); intersection != nil {
			res = append(res, intersection)
		}
		switch {
		case A[ai][1] == B[bi][1]:
			ai++
			bi++
		case A[ai][1] > B[bi][1]:
			bi++
		case A[ai][1] < B[bi][1]:
			ai++
		}
	}

	return res
}

func getIntersection(a, b []int) []int {
	if a[1] < b[0] || b[1] < a[0] {
		return nil
	}
	lessSlice := a
	grossSlice := b
	if b[0] < a[0] {
		lessSlice = b
		grossSlice = a
	}

	intersect := []int{
		grossSlice[0],
		grossSlice[0],
	}
	for ; intersect[1] < lessSlice[1] && intersect[1] < grossSlice[1]; intersect[1]++ {
	}

	return intersect
}
