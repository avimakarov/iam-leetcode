package main

import "fmt"

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)

	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

func main() {
	tests := []struct {
		arg1 int
		arg2 int
		want int
	}{
		{
			arg1: 3,
			arg2: 7,
			want: 28,
		},
	}

	for i, testCase := range tests {
		got := uniquePaths(testCase.arg1, testCase.arg2)
		if got != testCase.want {
			fmt.Printf("case: %d, got: %d; want: %d \n", i, got, testCase.want)
		}
	}
}
