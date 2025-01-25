package main

func moveZeroes(nums []int) {
	nonZeroIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[nonZeroIndex] = nums[nonZeroIndex], nums[i]
			nonZeroIndex++
		}
	}
}
