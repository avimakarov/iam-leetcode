package main

import (
	"fmt"
	"slices"
)

func equalPairs(grid [][]int) int {
	rows := make(map[int][]int)
	cols := make(map[int][]int)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			rows[i] = append(rows[i], grid[i][j])
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			cols[i] = append(cols[i], grid[j][i])
		}
	}

	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if slices.Equal(rows[i], cols[j]) {
				count++
			}
		}
	}
	return count
}

func main() {
	tests := []struct {
		args [][]int
		want int
	}{
		{
			want: 4,
			args: [][]int{
				{13, 13},
				{13, 13},
			},
		},
		{
			want: 1,
			args: [][]int{
				{3, 2, 1},
				{1, 7, 6},
				{2, 7, 7},
			},
		},
		{
			want: 3,
			args: [][]int{
				{3, 1, 2, 2},
				{1, 4, 4, 5},
				{2, 4, 2, 2},
				{2, 4, 2, 2},
			},
		},
		{
			want: 48,
			args: [][]int{
				{3, 3, 3, 6, 18, 3, 3, 3, 3, 3},
				{3, 3, 3, 3, 1, 3, 3, 3, 3, 3},
				{3, 3, 3, 3, 1, 3, 3, 3, 3, 3},
				{3, 3, 3, 3, 1, 3, 3, 3, 3, 3},
				{1, 1, 1, 11, 19, 1, 1, 1, 1, 1},
				{3, 3, 3, 18, 19, 3, 3, 3, 3, 3},
				{3, 3, 3, 3, 1, 3, 3, 3, 3, 3},
				{3, 3, 3, 3, 1, 3, 3, 3, 3, 3},
				{3, 3, 3, 1, 6, 3, 3, 3, 3, 3},
				{3, 3, 3, 3, 1, 3, 3, 3, 3, 3},
			},
		},
	}

	for i, testCase := range tests {
		got := equalPairs(testCase.args)
		if got != testCase.want {
			fmt.Printf("case: %d; got: %d; want: %d \n", i, got, testCase.want)
		}
	}
}
