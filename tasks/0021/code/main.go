package main

import (
	"fmt"
	"reflect"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

//	1 -> 2 -> 4 || list1
//	1 -> 3 -> 4 || list2
//	create new empty linked list as "result"
//	while list1 != nil && list2 != nil
//		list1.Val <= list2.Val, then
//			result.Next = list1 => result = 0 -> 1 -> 2 -> 4
//			list1 = 2 -> 4
//		---
//		list1.Val >= list2.Val, then
//			result.Next = list2 => result = 0 -> 1 -> 1 -> 3 -> 4
//			list2 = 3 -> 4
//		---
//		list1.Val < list2.Val, then
//			result.Next = list1 => result = 0 -> 1 -> 1 -> 2 -> 4
//			list1 = 4
//		---
//		list1.Val > list2.Val, then
//			result.Next = list2 -> result = 0 -> 1 -> 1 -> 2 -> 3
//			list2 = 4
//		---
//		list1.Val <= list2.Val, then
//			result.Next = list1 => result = 0 -> 1 -> 1 -> 2 -> 3 -> 4
//			list1 = nil
//		---
//	list1 is nil, then
//		result.Next = list2 => result = 0 -> 1 -> 1 -> 2 -> 3 -> 4 -> 4
//

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 != nil {
		return list2
	}
	if list1 != nil && list2 == nil {
		return list1
	}

	result := &ListNode{}
	currResult := result

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			currResult.Next = list1
			list1 = list1.Next
		} else {
			currResult.Next = list2
			list2 = list2.Next
		}
		currResult = currResult.Next
	}

	if list1 != nil {
		currResult.Next = list1
	} else {
		currResult.Next = list2
	}

	return result.Next
}

func main() {
	testCases := []struct {
		arg1     ListNode
		arg2     ListNode
		expected ListNode
	}{
		{
			arg1: ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
			arg2: ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
			expected: ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val:  4,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
	}

	for _, testCase := range testCases {
		got := mergeTwoLists(&testCase.arg1, &testCase.arg2)
		if !reflect.DeepEqual(got, testCase.expected) {
			fmt.Printf("expected %v, got %v \n", testCase.expected, got)
		}
	}
}
