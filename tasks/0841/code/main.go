package main

import "fmt"

func dfs(graph map[int][]int, visited map[int]bool, start int) {
	visited[start] = true
	for _, v := range graph[start] {
		if !visited[v] {
			dfs(graph, visited, v)
		}
	}
}

func canVisitAllRooms(rooms [][]int) bool {
	graph := make(map[int][]int)

	for i, room := range rooms {
		graph[i] = make([]int, 0)
		for _, k := range room {
			graph[i] = append(graph[i], k)
		}
	}

	visited := make(map[int]bool)
	dfs(graph, visited, 0)

	for v, _ := range graph {
		if !visited[v] {
			return false
		}
	}
	return true
}

func main() {
	tests := []struct {
		want bool
		args [][]int
	}{
		{
			want: true,
			args: [][]int{{1}, {2}, {3}, {}},
		},
		{
			want: false,
			args: [][]int{{1}, {2}, {2}, {}},
		},
		{
			want: false,
			args: [][]int{{1, 3}, {3, 0, 1}, {2}, {0}},
		},
	}

	for i, testCase := range tests {
		got := canVisitAllRooms(testCase.args)
		fmt.Printf("test case :: %d; ok :: %t; got: %t; want: %t \n", i, got == testCase.want, got, testCase.want)
	}
}
