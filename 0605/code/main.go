package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}

	if len(flowerbed) == 1 {
		if n == 1 && flowerbed[0] == 0 {
			return true
		}
		if n == 1 && flowerbed[0] == 1 {
			return false
		}
	}

	if flowerbed[0] == 0 && flowerbed[1] == 0 {
		n--
		flowerbed[0] = 1
	}

	for i := 1; i < len(flowerbed); i++ {
		if i == len(flowerbed)-1 {
			if flowerbed[i] == 0 && flowerbed[i-1] == 0 {
				n--
				flowerbed[i] = 1
			}
		}
		if flowerbed[i] == 0 && flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
			n--
			flowerbed[i] = 1
		}
	}

	return n <= 0
}

func main() {
	type testCaseArg struct {
		n   int
		arr []int
	}

	tests := []struct {
		want bool
		args testCaseArg
	}{
		{
			want: true,
			args: testCaseArg{
				n:   1,
				arr: []int{1, 0, 0, 0, 1},
			},
		},
		{
			want: false,
			args: testCaseArg{
				n:   2,
				arr: []int{1, 0, 0, 0, 1},
			},
		},
		{
			want: true,
			args: testCaseArg{
				n:   1,
				arr: []int{1, 0, 0, 0, 1, 0, 1},
			},
		},
		{
			want: false,
			args: testCaseArg{
				n:   1,
				arr: []int{1, 0, 1, 0, 1, 0, 1},
			},
		},
		{
			want: true,
			args: testCaseArg{
				n:   2,
				arr: []int{1, 0, 0, 0, 1, 0, 0},
			},
		},
	}

	for i, testCase := range tests {
		got := canPlaceFlowers(testCase.args.arr, testCase.args.n)
		if got == testCase.want {
			fmt.Printf("test: %d (DONE) \n", i)
		} else {
			fmt.Printf("test: %d (FAIL) got %t, expected %t \n", i, got, testCase.want)
		}
	}
}
