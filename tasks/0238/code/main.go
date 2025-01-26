package main

import (
	"fmt"
	"reflect"
)

func prodArr(exclude int, nums []int) int {
	prod := 1
	for i := 0; i < len(nums); i++ {
		if exclude == i {
			continue
		}
		prod *= nums[i]
	}

	return prod
}

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res[i] = prodArr(i, nums)
	}

	return res
}

func main() {
	tests := []struct {
		args []int
		want []int
	}{
		{
			args: []int{1, 2, 3, 4},
			want: []int{24, 12, 8, 6},
		},
		{
			want: []int{0, 0, 9, 0, 0},
			args: []int{-1, 1, 0, -3, 3},
		},
	}

	for i, testCase := range tests {
		got := productExceptSelf(testCase.args)
		if !reflect.DeepEqual(got, testCase.want) {
			fmt.Printf("case: %d; got: %v; want: %v \n", i, got, testCase.want)
		}
	}
}
