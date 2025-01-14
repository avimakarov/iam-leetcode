package main

import (
	"fmt"
	"sort"
)

func maxOperations(nums []int, k int) int {
	res := 0

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	l := 0
	r := len(nums) - 1

	for l < r {
		sum := nums[l] + nums[r]

		if sum == k {
			l++
			r--

			res++
		} else {
			if sum < k {
				l++
			} else {
				r--
			}
		}
	}

	return res
}

func main() {
	type testCaseArgs struct {
		nums   []int
		target int
	}

	type testCase struct {
		want int
		args testCaseArgs
	}

	tests := []testCase{
		{
			want: 2,
			args: testCaseArgs{
				nums:   []int{1, 2, 3, 4},
				target: 5,
			},
		},
		{
			want: 1,
			args: testCaseArgs{
				nums:   []int{3, 1, 3, 4, 3},
				target: 6,
			},
		},
	}

	for i, testCase := range tests {
		fmt.Printf(
			"test case: %d :: %t \n", i, maxOperations(testCase.args.nums, testCase.args.target) == testCase.want,
		)
	}
}
