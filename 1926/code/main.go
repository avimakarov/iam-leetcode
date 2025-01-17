package main

import "fmt"

func convert(maze [][]string) [][]byte {
	res := make([][]byte, 0, len(maze))
	for _, row := range maze {
		temp := make([]byte, 0, len(row))
		for _, col := range row {
			temp = append(temp, []byte(col)...)
		}
		res = append(res, temp)
	}
	return res
}

const (
	free = 46
	wall = 43
)

type coord struct {
	i, j int
}

func isExit(maze [][]byte, c coord) bool {
	return c.i == 0 || c.i == len(maze)-1 || c.j == 0 || c.j == len(maze[0])-1
}

func canTravel(maze [][]byte, c coord) bool {
	return c.i >= 0 && c.j >= 0 && c.i < len(maze) && c.j < len(maze[0]) && maze[c.i][c.j] != wall
}

func nearestExit(maze [][]byte, entrance []int) int {
	directions := []int{-1, 0, 1, 0, -1}

	res := 1

	queue := make([]coord, 0)
	visited := make(map[coord]bool)

	visited[coord{entrance[0], entrance[1]}] = true
	queue = append(queue, coord{entrance[0], entrance[1]})

	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			next := queue[0]
			queue = queue[1:]

			for jj := 0; jj < 4; jj++ {
				newCoord := coord{
					i: next.i + directions[jj],
					j: next.j + directions[jj+1],
				}
				canMove := canTravel(maze, newCoord)
				if canMove && !visited[newCoord] {
					if isExit(maze, newCoord) {
						return res
					}
					visited[newCoord] = true
					queue = append(queue, newCoord)
				}
			}
		}
		res++
	}

	return -1
}

func main() {
	type testCaseArgs struct {
		from []int
		maze [][]string
	}
	tests := []struct {
		args testCaseArgs
		want int
	}{
		{
			want: 1,
			args: testCaseArgs{
				from: []int{1, 2},
				maze: [][]string{
					{"+", "+", ".", "+"},
					{".", ".", ".", "+"},
					{"+", "+", "+", "."},
				},
			},
		},
		{
			want: 2,
			args: testCaseArgs{
				from: []int{1, 0},
				maze: [][]string{
					{"+", "+", "+"},
					{".", ".", "."},
					{"+", "+", "+"},
				},
			},
		},
		{
			want: -1,
			args: testCaseArgs{
				from: []int{0, 0},
				maze: [][]string{
					{".", "+"},
				},
			},
		},
		{
			want: -1,
			args: testCaseArgs{
				from: []int{0, 1},
				maze: [][]string{
					{"+", "."},
				},
			},
		},
		{
			want: 1,
			args: testCaseArgs{
				from: []int{0, 1},
				maze: [][]string{
					{".", "."},
				},
			},
		},
		{
			want: 1,
			args: testCaseArgs{
				from: []int{2, 0},
				maze: [][]string{
					{"."},
					{"."},
					{"."},
					{"."},
				},
			},
		},
		{
			want: 1,
			args: testCaseArgs{
				from: []int{0, 0},
				maze: [][]string{
					{".", "+", "+", "+", "+"},
					{".", "+", ".", ".", "."},
					{".", "+", ".", "+", "."},
					{".", ".", ".", "+", "."},
					{"+", "+", "+", "+", "."},
				},
			},
		},
	}

	for i, testCase := range tests {
		got := nearestExit(
			convert(testCase.args.maze), testCase.args.from,
		)
		fmt.Printf("test case :: %d; want :: %d; got :: %d; done :: %t \n", i, testCase.want, got, testCase.want == got)
	}
}
