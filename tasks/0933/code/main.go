package main

import "fmt"

type RecentCounter struct {
	requests []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		requests: make([]int, 0),
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.requests = append(this.requests, t)

	timeStart, timeEnd := t-3000, t

	for _, timestamp := range this.requests {
		if !(timestamp >= timeStart && timestamp <= timeEnd) {
			this.requests = this.requests[1:]
		}
	}

	return len(this.requests)
}

func main() {
	c := Constructor()
	fmt.Println(c.Ping(1))
	fmt.Println(c.Ping(100))
	fmt.Println(c.Ping(3001))
	fmt.Println(c.Ping(3002))
}
