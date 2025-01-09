package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func length(head *ListNode) int {
	l := 0
	h := head

	for {
		l += 1
		h = h.Next

		if h == nil {
			break
		}
	}

	return l
}

func deleteMiddle(head *ListNode) *ListNode {
	var (
		d int

		c = head

		l = length(head)
	)

	if c == nil || l == 1 {
		return nil
	}

	for i := 0; i < l/2-1; i++ {
		c = c.Next
	}

	other := c.Next
	c.Next = other.Next

	return head
}

func main() {
	tests := []*ListNode{
		// nil
		nil,

		// one node list
		&ListNode{
			Val:  2,
			Next: nil,
		},

		// delete the last
		// element from the node list
		&ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
		// delete the element from the
		// node list with the even length
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
		},
		// delete the element from the
		// node list with the odd  length
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val: 6,
								Next: &ListNode{
									Val:  7,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
	}

	for i, testCase := range tests {
		deleteMiddle(testCase)
		fmt.Printf("test %d: done \n", i)
	}
}
