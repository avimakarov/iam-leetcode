package main

import (
	"fmt"
	"sort"
)

type pQueue struct {
	k    int
	nums []int
}

func newQueue(n int, k int) *pQueue {
	return &pQueue{
		k:    k,
		nums: make([]int, 0, n),
	}
}

func (p *pQueue) nth(n int) int {
	return p.nums[n-1]
}

func (p *pQueue) enqueue(num int) {
	if len(p.nums) == 0 {
		p.nums =
			append(p.nums, num)
		return
	}

	if p.nums[0] < num {
		p.nums = append(
			[]int{num}, p.nums...,
		)
		return
	}

	if p.nums[len(p.nums)-1] > num {
		p.nums = append(p.nums, num)
		return
	}

	for i := 0; i < len(p.nums); i++ {
		if i+1 < len(p.nums) && p.nums[i+1] < num {
			p.nums = append(
				p.nums[:i+1], append([]int{num}, p.nums[i+1:]...)...,
			)
			return
		}
	}

	p.nums = append(p.nums, num)
}

func findKthLargestSlow(nums []int, k int) int {
	q := newQueue(len(nums), k)
	for _, num := range nums {
		q.enqueue(num)
	}

	return q.nth(k)
}

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
		got := findKthLargestSlow(testCase.args.nums, testCase.args.k)
		fmt.Printf(
			"test case: %d :: %t :: got :: %d :: want %d \n", i, got == testCase.want, got, testCase.want,
		)
	}
}
