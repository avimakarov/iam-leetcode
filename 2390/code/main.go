package main

import (
	"fmt"
	"strings"
)

type Stack struct {
	len    int
	values []string
}

func (s *Stack) Pop() {
	s.values = s.values[:len(s.values)-1]
	s.len--
}

func (s *Stack) Len() int {
	return s.len
}

func (s *Stack) Add(v string) {
	s.len++
	s.values = append(s.values, v)
}

func removeStars(s string) string {
	if s == "" {
		return ""
	}

	stack := new(Stack)
	input := strings.Split(s, "")

	for _, char := range input {
		if char == "*" {
			stack.Pop()
		} else {
			stack.Add(char)
		}
	}

	return strings.Join(stack.values, "")
}

func main() {
	tests := []struct {
		args string
		want string
	}{
		{
			args: "leet**cod*e",
			want: "lecoe",
		},
		{
			args: "erase*****",
			want: "",
		},
	}

	for i, testCase := range tests {
		fmt.Printf(
			"test case: %d :: %t \n", i, removeStars(testCase.args) == testCase.want,
		)
	}
}
