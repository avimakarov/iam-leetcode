package main

import (
	"fmt"
)

// В каждом раунде
// сенатор обладает только 2-мя возможностями:
//
// 1. Объявить о победе
// 2. Заблокировать, тогда сл справа теряет возможность голосовать всегда
func predictPartyVictory(senate string) string {
	n := len(senate)
	r, d := make([]int, 0, len(senate)), make([]int, 0, len(senate))

	for i, char := range senate {
		if char == 'R' {
			r = append(r, i)
		}
		if char == 'D' {
			d = append(d, i)
		}
	}

	for len(r) > 0 && len(d) > 0 {
		r_idx, d_idx := r[0], d[0]
		r, d = r[1:], d[1:]

		if r_idx < d_idx {
			r = append(r, n+r_idx)
		} else {
			d = append(d, n+d_idx)
		}
	}

	if len(d) > 0 {
		return "Dire"
	}

	return "Radiant"
}

func main() {
	type test struct {
		arg string
		exp string
	}
	tests := []test{
		{
			arg: "RD",
			exp: "Radiant",
		},
		{
			arg: "RDD",
			exp: "Dire",
		},
		{
			arg: "RRD",
			exp: "Radiant",
		},
		{
			arg: "DDRRR",
			exp: "Dire",
		},
		{
			arg: "DRRDRDRDRDDRDRDR",
			exp: "Radiant",
		},
	}

	for i, testCase := range tests {
		got := predictPartyVictory(testCase.arg)
		if got == testCase.exp {
			fmt.Printf("test: %d (DONE) \n", i)
		} else {
			fmt.Printf("test: %d (FAIL) got %s, expected %s \n", i, got, testCase.exp)
		}
	}
}
