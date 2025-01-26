package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	l := len(cost)
	d := make([]int, l)

	d[0] = cost[0]
	d[1] = cost[1]

	for i := 2; i < l; i++ {
		d[i] = cost[i] + min(d[i-1], d[i-2])
	}

	return min(d[l-1], d[l-2])
}

func main() {
	tests := []struct {
		args []int
		want int
	}{
		{
			args: []int{10, 15, 20},
			want: 15,
		},
	}

	for i, testCase := range tests {
		got := minCostClimbingStairs(testCase.args)
		if got != testCase.want {
			fmt.Printf("case %d: got %d, want %d\n", i, got, testCase.want)
		}
	}
}
