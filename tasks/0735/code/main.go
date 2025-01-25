package main

import (
	"fmt"
	"reflect"
)

func asteroidCollision(asteroids []int) []int {
	stack := []int{}
	for _, i := range asteroids {
		stack = append(stack, i)

		for len(stack) > 1 && stack[len(stack)-1] < 0 && stack[len(stack)-1]*stack[len(stack)-2] < 0 {
			if -stack[len(stack)-1] > stack[len(stack)-2] {
				stack[len(stack)-2] = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			} else if -stack[len(stack)-1] < stack[len(stack)-2] {
				stack = stack[:len(stack)-1]
			} else if -stack[len(stack)-1] == stack[len(stack)-2] {
				stack = stack[:len(stack)-2]
			}
		}
	}

	return stack
}

func main() {
	tests := []struct {
		args []int
		want []int
	}{
		{
			args: []int{5, 10, -5},
			want: []int{5, 10},
		},
		{
			args: []int{8, -8},
			want: []int{},
		},
		{
			args: []int{10, 2, -5},
			want: []int{10},
		},
	}

	for i, testCase := range tests {
		got := asteroidCollision(testCase.args)
		if !reflect.DeepEqual(got, testCase.want) {
			fmt.Printf("case: %d; want: %v; got: %v \n", i, testCase.want, got)
		}
	}
}
