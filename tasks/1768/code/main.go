package main

import (
	"fmt"
	"strings"
)

func mergeAlternately(word1 string, word2 string) string {
	maxLen := len(word1)
	if len(word2) > maxLen {
		maxLen = len(word2)
	}

	res := make([]string, 0, len(word1)+len(word2))

	for i := 0; i < maxLen; i++ {
		if i < len(word1) {
			res = append(res, string(word1[i]))
		}
		if i < len(word2) {
			res = append(res, string(word2[i]))
		}
	}

	return strings.Join(res, "")
}

func main() {
	tests := []struct {
		arg1 string
		arg2 string
		want string
	}{
		{
			arg1: "abc",
			arg2: "pqr",
			want: "apbqcr",
		},
		{
			arg1: "ab",
			arg2: "pqrs",
			want: "apbqrs",
		},
		{
			arg1: "abcd",
			arg2: "pq",
			want: "apbqcd",
		},
	}

	for i, testCase := range tests {
		got := mergeAlternately(
			testCase.arg1, testCase.arg2,
		)
		fmt.Printf("test case :: %d; want :: %s; got :: %s; done :: %t \n", i, testCase.want, got, testCase.want == got)

	}
}
