package main

import "fmt"

const (
	empty  = 0
	fresh  = 1
	rotten = 2
)

type coordinate struct {
	i, j int
}

func orangesRotting(grid [][]int) int {
	stat := map[int]int{
		empty:  0,
		fresh:  0,
		rotten: 0,
	}

	steps := 0
	start := make([]coordinate, 0)
	directions := []int{-1, 0, 1, 0, -1}

	for i, row := range grid {
		for j, cell := range row {
			stat[cell]++
			if cell == rotten {
				start = append(start, coordinate{i, j})
			}
		}
	}

	if stat[fresh] == 0 {
		return 0
	}

	if stat[rotten] == 0 && stat[fresh] != 0 {
		return -1
	}

	// bfs

	queue := make([]coordinate, 0)
	visited := make(map[coordinate]bool)

	for _, s := range start {
		visited[s] = true
		queue = append(queue, s)
	}

	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			next := queue[0]
			queue = queue[1:]
			for j := 0; j < 4; j++ {
				newCoordinate := coordinate{
					i: next.i + directions[j],
					j: next.j + directions[j+1],
				}
				if !visited[newCoordinate] && newCoordinate.i >= 0 && newCoordinate.j >= 0 && newCoordinate.i < len(grid) && newCoordinate.j < len(grid[0]) && grid[newCoordinate.i][newCoordinate.j] == fresh {
					visited[newCoordinate] = true
					queue = append(queue, newCoordinate)
					stat[fresh]--
					stat[rotten]++
				}
			}
		}
		steps++
	}

	if stat[fresh] > 0 {
		return -1
	}

	return steps - 1
}

func main() {
	tests := []struct {
		args [][]int
		want int
	}{
		{
			want: 4,
			args: [][]int{
				{2, 1, 1},
				{1, 1, 0},
				{0, 1, 1},
			},
		},
		{
			want: -1,
			args: [][]int{
				{2, 1, 1},
				{0, 1, 1},
				{1, 0, 1},
			},
		},
		{
			want: 0,
			args: [][]int{
				{0, 2},
			},
		},
		{
			want: 2,
			args: [][]int{
				{2, 1, 1},
				{1, 1, 1},
				{0, 1, 2},
			},
		},
	}

	for i, testCase := range tests {
		got := orangesRotting(testCase.args)
		fmt.Printf(
			"test case: %d :: %t :: got :: %d :: want %d \n", i, got == testCase.want, got, testCase.want,
		)
	}
}
