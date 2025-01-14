package main

import (
	"fmt"
	"reflect"
	"strings"
)

var phones = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func bt(combinations []string, digits []string) []string {
	for {
		if len(digits) == 0 {
			return combinations
		}

		newCombinations := make([]string, 0, len(combinations))

		for i, _ := range combinations {
			for _, d := range phones[digits[0]] {
				newCombinations = append(newCombinations, combinations[i]+d)
			}
		}

		digits = digits[1:]
		combinations = newCombinations
	}
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return make([]string, 0)
	}

	d := strings.Split(digits, "")
	result := make([]string, 0, len(d))

	for _, char := range phones[d[0]] {
		result = append(result, bt([]string{char}, d[1:])...)
	}

	return result
}

func main() {
	type testCase struct {
		args string
		want []string
	}
	tests := []testCase{
		{
			args: "2",
			want: []string{
				"a", "b", "c",
			},
		},
		{
			args: "23",
			want: []string{
				"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf",
			},
		},
	}

	for i, testCase := range tests {
		fmt.Printf(
			"test case: %d; %t \n",
			i, reflect.DeepEqual(letterCombinations(testCase.args), testCase.want),
		)
	}
}
