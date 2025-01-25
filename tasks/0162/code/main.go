package main

import (
	"fmt"
	"sort"
)

func findPeakElement(nums []int) int {
	hm := make(map[int][]int)
	for i, num := range nums {
		hm[num] = append(hm[num], i)
	}

	sort.Ints(nums)

	return hm[nums[len(nums)-1]][0]
}

func main() {
	tests := []struct {
		nums []int
		want int
	}{
		{
			want: 2,
			nums: []int{1, 2, 3, 1},
		},
		{
			want: 5,
			nums: []int{1, 2, 1, 3, 5, 6, 4},
		},
	}

	for i, testCase := range tests {
		got := findPeakElement(testCase.nums)
		if got != testCase.want {
			fmt.Printf("case: %d; got: %d; want: %d \n", i, got, testCase.want)
		}
	}
}
