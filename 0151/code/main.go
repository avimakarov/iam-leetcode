package main

import (
	"fmt"
	"strings"
)

type stack struct {
	words []string
}

func (s *stack) len() int {
	return len(s.words)
}

func (s *stack) pop() string {
	res := s.words[len(s.words)-1]
	s.words = s.words[:len(s.words)-1]

	return res
}

func (s *stack) push(word string) {
	s.words = append(s.words, word)
}

func reverseWords(s string) string {
	if s == "" {
		return s
	}

	rs := []string{}
	ss := new(stack)

	for _, word := range strings.Fields(s) {
		ss.push(word)
	}

	for ss.len() > 0 {
		rs = append(rs, ss.pop())
	}

	return strings.Join(rs, " ")
}

func main() {
	tests := []struct {
		args string
		want string
	}{
		{
			args: "the sky is blue",
			want: "blue is sky the",
		},
	}

	for i, testCase := range tests {
		got := reverseWords(testCase.args)
		fmt.Printf("test case :: %d, done :: %t, want :: %s, got :: %s \n", i, got == testCase.want, testCase.want, got)
	}
}
