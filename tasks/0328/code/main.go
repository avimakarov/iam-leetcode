package main

import (
	"fmt"
	"reflect"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	var l = 0
	var last = head

	for {
		if last.Next == nil {
			break
		}
		l++
		last = last.Next
	}

	var i = 1
	var prev = head
	var curr = head.Next
	for {
		if i > l {
			break
		}
		// если индекс текущий эелемент - нечетный - отправляем его в конец списка
		if i%2 != 0 {
			last.Next = &ListNode{
				Val:  curr.Val,
				Next: nil,
			}
			last = last.Next
			prev.Next = curr.Next
		}
		i++
		prev = curr
		curr = curr.Next
	}

	return head
}

func main() {
	tests := []struct {
		args *ListNode
		want *ListNode
	}{
		{
			args: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val:  4,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}

	for i, testCase := range tests {
		got := oddEvenList(testCase.args)
		if !reflect.DeepEqual(got, testCase.args) {
			fmt.Printf("case %d: got %v; want %v \n", i, got, testCase.args)
		}
	}
}
