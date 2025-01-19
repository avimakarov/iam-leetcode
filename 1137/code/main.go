package main

import "fmt"

func tribonacci(n int) int {
	t0 := 0
	t1 := 1
	t2 := 1

	if n == 0 {
		return t0
	}

	if n == 1 {
		return t1
	}

	if n == 2 {
		return t2
	}

	for i := 3; i <= n; i++ {
		t0, t1, t2 = t1, t2, t0+t1+t2
	}

	return t2
}

func main() {
	tests := []struct {
		args int
		want int
	}{
		{
			args: 4,
			want: 4,
		},
		{
			args: 25,
			want: 1389537,
		},
	}

	for i, testCase := range tests {
		got := tribonacci(testCase.args)
		if got == testCase.want {
			fmt.Printf("test: %d (DONE) \n", i)
		} else {
			fmt.Printf("test: %d (FAIL) got %d, expected %d \n", i, got, testCase.want)
		}
	}
}
