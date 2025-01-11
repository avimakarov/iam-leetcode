package main

import (
	"fmt"
	"reflect"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}

	if root.Val > val {
		return searchBST(root.Left, val)
	} else {
		return searchBST(root.Right, val)
	}
}

func main() {
	tree := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 7,
		},
	}

	type testCase struct {
		args int
		tree *TreeNode
		want *TreeNode
	}

	tests := []testCase{
		{
			args: 2,
			tree: tree,
			want: tree.Left,
		},
		{
			args: 5,
			want: nil,
			tree: tree,
		},
	}

	for i, test := range tests {
		got := searchBST(test.tree, test.args)
		if reflect.DeepEqual(got, test.want) {
			fmt.Printf("test %d: DONE \n", i)
		} else {
			fmt.Printf(
				"test failed - want: %v, got: %v\n", test.want, got,
			)
		}
	}
}
