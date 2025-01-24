package main

import (
	"fmt"
	"math/bits"
	"slices"
)

func countBits(n int) []int {
	res := make([]int, 0, n)
	for i := 0; i < n+1; i++ {
		res = append(res, bits.OnesCount(uint(i)))

	}
	return res
}

func main() {
	tests := []struct {
		args int
		want []int
	}{
		{
			args: 2,
			want: []int{0, 1, 1},
		},
		{
			args: 5,
			want: []int{0, 1, 1, 2, 1, 2},
		},
	}

	for i, testCase := range tests {
		got := countBits(testCase.args)
		if !slices.Equal(got, testCase.want) {
			fmt.Printf("case: %d; got: %d; want: %d \n", i, got, testCase.want)
		}
	}
}
