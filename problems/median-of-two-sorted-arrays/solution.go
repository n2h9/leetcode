package solution

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1, n2 := len(nums1), len(nums2)
	n3 := n1 + n2
	merge := make([]int, 0, n3)
	for i1, i2 := 0, 0; i1 < n1 || i2 < n2; {
		for ; i1 < n1 && (i2 >= n2 || nums1[i1] <= nums2[i2]); i1++ {
			merge = append(merge, nums1[i1])
		}
		for ; i2 < n2 && (i1 >= n1 || nums2[i2] <= nums1[i1]); i2++ {
			merge = append(merge, nums2[i2])
		}
	}
	if n3 & 1 == 1 {
		// n3 is odd
		return float64(merge[n3/2])
	}
	return float64((merge[n3/2-1]+merge[n3/2]))/2.0
}
