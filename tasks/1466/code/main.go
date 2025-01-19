package main

import "fmt"

type counter int

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func dfs(graph map[int][]int, counter *counter, visited map[int]bool, start int) {
	visited[abs(start)] = true
	for _, v := range graph[abs(start)] {
		if !visited[abs(v)] {
			if v > 0 {
				*counter++
			}
			dfs(graph, counter, visited, v)
		}
	}
}

func minReorder(n int, connections [][]int) int {
	graph := make(map[int][]int)

	for _, connection := range connections {
		graph[connection[0]] = append(graph[connection[0]], +connection[1])
		graph[connection[1]] = append(graph[connection[1]], -connection[0])
	}

	counter := counter(0)
	visited := make(map[int]bool)
	dfs(graph, &counter, visited, 0)

	return int(counter)
}

func main() {
	type testCaseArg struct {
		vertexCount int
		connections [][]int
	}
	tests := []struct {
		args testCaseArg
		want int
	}{
		{
			args: testCaseArg{
				vertexCount: 6,
				connections: [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}},
			},
			want: 3,
		},
		{
			want: 3,
			args: testCaseArg{
				vertexCount: 6,
				connections: [][]int{{4, 5}, {0, 1}, {1, 3}, {2, 3}, {4, 0}},
			},
		},
	}

	for i, testCase := range tests {
		got := minReorder(testCase.args.vertexCount, testCase.args.connections)
		fmt.Printf(
			"test case: %d :: %t :: got :: %d :: want %d \n", i, got == testCase.want, got, testCase.want,
		)
	}
}
