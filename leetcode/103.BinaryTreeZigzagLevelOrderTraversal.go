package main

type Node struct {
	data  string
	left  *Node
	right *Node
}

func zigzagLevelOrder(root *TreeNode) [][]int {
    if root == nil { return [][]int{} }
    res := make([][]int, 0)
    appendNode(root, &res, 0)
    for i := range res {
        if i % 2 == 1 {
            reverseSlice(res[i])
        }
    }
    return res
}

func appendNode(node *TreeNode, res *[][]int, h int) {
    if node == nil { return }
    if len(*res) < h + 1 {
        *res = append(*res, []int{})
    }
    (*res)[h] = append((*res)[h], node.Val)
    appendNode(node.Left, res, h + 1)
    appendNode(node.Right, res, h + 1)
}

func reverseSlice(s []int) {
    i, j := 0, len(s) - 1
    for i < j {
        s[i], s[j] = s[j], s[i]
        i++
        j--
    }
}


