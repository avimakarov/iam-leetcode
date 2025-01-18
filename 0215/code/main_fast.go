package main

import (
	"fmt"
	"sort"
)

func findKthLargestFast(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	return nums[k-1]
}

func main() {
	type testCaseArgs struct {
		k    int
		nums []int
	}
	tests := []struct {
		want int
		args testCaseArgs
	}{
		{
			want: 5,
			args: testCaseArgs{
				k:    2,
				nums: []int{3, 2, 1, 5, 6, 4},
			},
		},
		{
			want: -1,
			args: testCaseArgs{
				k:    2,
				nums: []int{-1, -1},
			},
		},
		{
			want: 1,
			args: testCaseArgs{
				k:    9,
				nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			},
		},
	}

	for i, testCase := range tests {
		got := findKthLargestFast(testCase.args.nums, testCase.args.k)
		fmt.Printf(
			"test case: %d :: %t :: got :: %d :: want %d \n", i, got == testCase.want, got, testCase.want,
		)
	}
}
