package main

type TreeNode struct {
	Val int
	Left *TreeNode
    Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
    var stack []*TreeNode
    var res []int
    for root != nil || len(stack) > 0 {
        for root != nil {
        	res = append(res, root.Val)
            stack = append(stack, root)
            root = root.Left
        }
        root = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        root = root.Right
    }
    return res
}
