package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	res := 0

	// idea::important
	// tag::stack,binaryTree

	// traceTree means display all
	// paths from the root to any vertex
	//
	// if we sum the values in reverse order
	// we can find a subset of vertices whose
	// sum of values corresponds to the desired one
	traceTree(root, []int{}, targetSum, &res)

	return res
}

func traceTree(root *TreeNode, path []int, target int, res *int) {
	if root == nil {
		return
	}

	path = append(path, root.Val)

	localSum := 0
	for i := len(path) - 1; i >= 0; i-- {
		localSum += path[i]
		if localSum == target {
			*res++
		}
	}
	traceTree(root.Left, path, target, res)
	traceTree(root.Right, path, target, res)
}

func main() {
	type test struct {
		want   int
		tree   *TreeNode
		target int
	}

	tests := []test{
		{
			want:   3,
			target: 8,
			tree: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val:   3,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   -2,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &TreeNode{
						Val:  2,
						Left: nil,
						Right: &TreeNode{
							Val:   1,
							Left:  nil,
							Right: nil,
						},
					},
				},
				Right: &TreeNode{
					Val:  -3,
					Left: nil,
					Right: &TreeNode{
						Val:   11,
						Left:  nil,
						Right: nil,
					},
				},
			},
		},
	}

	for i, test := range tests {
		fmt.Printf(
			"test %d: %t \n",
			i, test.want == pathSum(test.tree, test.target),
		)
	}
}
