package main

import "fmt"

func uniqueOccurrences(arr []int) bool {
	hashmap1 := make(map[int]int, len(arr))
	for _, v := range arr {
		hashmap1[v]++
	}

	hashmap2 := make(map[int]bool, len(hashmap1))
	for _, v := range hashmap1 {
		if _, exists := hashmap2[v]; !exists {
			hashmap2[v] = true
		} else {
			return false
		}
	}

	return true
}

func main() {
	tests := []struct {
		args []int
		want bool
	}{
		{
			want: true,
			args: []int{1, 2, 2, 1, 1, 3},
		},
		{
			want: false,
			args: []int{1, 2},
		},
		{
			want: true,
			args: []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0},
		},
	}

	for i, testCase := range tests {
		got := uniqueOccurrences(testCase.args)
		fmt.Printf(
			"test case: %d :: %t :: got :: %t :: want %t \n", i, got == testCase.want, got, testCase.want,
		)
	}
}
