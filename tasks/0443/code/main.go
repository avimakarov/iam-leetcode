package main

import "fmt"

type compressed []compressedChunk

type compressedChunk struct {
	char  byte
	count int
}

func compress(chars []byte) int {
	if len(chars) == 0 {
		return 0
	}
	if len(chars) == 1 {
		return 1
	}

	l := 0
	r := 1

	localCount := 0
	res := make(compressed, 0)

	for r < len(chars) {
		if chars[l] == chars[r] {
			r++
			localCount++
		} else {
			res = append(res, compressedChunk{char: chars[l], count: localCount})

			localCount = 0
			r++
			l = r - 1
		}

	}

	res = append(res, compressedChunk{char: chars[l], count: localCount})

	result := ""
	for _, r := range res {
		if r.count == 0 {
			result += string(r.char)
		}
		if r.count >= 1 {
			result += string(r.char)
			result += fmt.Sprintf("%d", r.count+1)
		}
	}

	for i, char := range result {
		chars[i] = byte(char)
	}
	chars = chars[:len(result)]

	return len(chars)
}

func main() {
	tests := []struct {
		args []string
		want int
	}{
		{
			args: []string{"a", "a", "b", "b", "c", "c", "c"},
			want: 6,
		},
		{
			args: []string{"a"},
			want: 1,
		},
		{
			args: []string{"a", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b"},
			want: 4,
		},
	}

	for i, testCase := range tests {
		byteArgs := make([]byte, 0, len(testCase.args))
		for _, arg := range testCase.args {
			byteArgs = append(byteArgs, []byte(arg)...)
		}
		fmt.Printf(
			"test case: %d :: %t :: got :: %d :: want %d \n", i, compress(byteArgs) == testCase.want, compress(byteArgs), testCase.want,
		)
	}

}
