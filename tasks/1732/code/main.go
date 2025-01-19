package main

import "fmt"

func largestAltitude(gain []int) int {
	cur := 0
	max := 0

	for _, g := range gain {
		cur += g
		if cur > max {
			max = cur
		}
	}

	return max
}

func main() {
	tests := []struct {
		want int
		args []int
	}{
		{
			want: 1,
			args: []int{-5, 1, 5, 0, -7},
		},
		{
			want: 0,
			args: []int{-4, -3, -2, -1, 4, 3, 2},
		},
	}

	for i, testCase := range tests {
		got := largestAltitude(testCase.args)
		if got == testCase.want {
			fmt.Printf("test: %d (DONE) \n", i)
		} else {
			fmt.Printf("test: %d (FAIL) got %d, expected %d \n", i, got, testCase.want)
		}
	}
}
