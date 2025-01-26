package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode, depth int) int {
	if node == nil {
		return 0
	}

	return 1 + max(dfs(node.Left, depth+1), dfs(node.Right, depth+1))
}

func maxDepth(root *TreeNode) int {
	return dfs(root, 0)
}

func main() {
	tests := []struct {
		args *TreeNode
		want int
	}{
		{
			want: 3,
			args: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
		},
	}

	for i, testCase := range tests {
		got := maxDepth(testCase.args)
		if got != testCase.want {
			fmt.Printf("case: %d, got: %d, want: %d", i, got, testCase.want)
		}
	}
}
