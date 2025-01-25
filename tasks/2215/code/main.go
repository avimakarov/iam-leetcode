package main

func findDifference(nums1 []int, nums2 []int) [][]int {
	elemetsOfArr1 := make(map[int]bool)
	elemetsOfArr2 := make(map[int]bool)

	for _, num := range nums1 {
		elemetsOfArr1[num] = true
	}

	for _, num := range nums2 {
		elemetsOfArr2[num] = true
	}

	res1 := make([]int, 0)
	res2 := make([]int, 0)

	for k, _ := range elemetsOfArr1 {
		if !elemetsOfArr2[k] {
			res1 = append(res1, k)
		}
	}

	for k, _ := range elemetsOfArr2 {
		if !elemetsOfArr1[k] {
			res2 = append(res2, k)
		}
	}

	return [][]int{res1, res2}
}
