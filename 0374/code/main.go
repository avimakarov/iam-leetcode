package main

func guessNumber(n int) int {
	if guess(n) == 0 {
		return n
	}

	l := 1
	h := n

	for l <= h {

		m := l + ((h - l) / 2)

		switch guess(m) {
		case 0:
			{
				return m
			}
		case 1:
			{
				l = m + 1
			}
		case -1:
			{
				h = m - 1
			}
		}
	}

	return -1
}
