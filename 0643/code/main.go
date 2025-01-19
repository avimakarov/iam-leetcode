package main

import "fmt"

func findMaxAverage(nums []int, k int) float64 {
	var res float64 = -1e10

	if len(nums) == k {
		var curr float64
		for _, n := range nums {
			curr += float64(n)
		}
		return curr / float64(k)
	}

	var curr float64

	l := 0
	r := l + k

	for _, n := range nums[l:r] {
		curr += float64(n)
	}

	if curr/float64(k) > res {
		res = curr / float64(k)
	}

	for r < len(nums) {
		l++
		r = l + k
		curr = curr - float64(nums[l-1])
		curr = curr + float64(nums[r-1])
		if curr/float64(k) > res {
			res = curr / float64(k)
		}
	}

	return res
}

func main() {
	type testCaseArg struct {
		k    int
		nums []int
	}
	tests := []struct {
		want float64
		args testCaseArg
	}{
		{
			want: 12.75,
			args: testCaseArg{
				k:    4,
				nums: []int{1, 12, -5, -6, 50, 3},
			},
		},
	}

	for i, testCase := range tests {
		got := findMaxAverage(testCase.args.nums, testCase.args.k)
		fmt.Printf("test case :: %d; ok :: %t; got: %f; want: %f \n", i, got == testCase.want, got, testCase.want)
	}
}
