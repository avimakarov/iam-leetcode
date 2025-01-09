package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type Stack struct {
	len int
	val []int
}

func (s *Stack) pop() int {
	i := s.len - 1
	val := s.val[i]

	s.len--
	s.val = s.val[0:i]

	return val
}

func (s *Stack) add(n int) {
	s.len++
	s.val = append(s.val, n)
}

func (s *Stack) empty() bool {
	return s.len == 0
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	s := Stack{}
	curr := head

	for {
		s.add(curr.Val)

		curr = curr.Next
		if curr == nil {
			break
		}
	}

	if s.empty() {
		return nil
	}

	res := &ListNode{}
	cur := res
	for {
		v := s.pop()
		cur.Val = v

		if s.empty() {
			break
		} else {
			cur.Next = &ListNode{}
			cur = cur.Next
		}

	}

	return res
}

func main() {
	tests := []*ListNode{
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}

	for i, testCase := range tests {
		r := reverseList(testCase)
		fmt.Printf("test %d (%v): done \n", i, r)
	}
}
