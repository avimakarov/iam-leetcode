package main

import "fmt"

func singleNumber(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}

	res := 0
	for i := 0; i < l; i++ {
		res ^= nums[i]
	}

	return res
}

func main() {
	tests := []struct {
		want int
		args []int
	}{
		{
			want: 1,
			args: []int{2, 2, 1},
		},
		{
			want: 4,
			args: []int{4, 1, 2, 1, 2},
		},
		{
			want: 1,
			args: []int{1},
		},
	}

	for i, testCase := range tests {
		got := singleNumber(testCase.args)
		if got != testCase.want {
			fmt.Printf("case: %d; got: %d; want: %d \n", i, got, testCase.want)
		}
	}
}
