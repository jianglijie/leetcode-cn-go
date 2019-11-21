package main


type BSTIterator struct {
    Val []int
}


func Constructor(root *TreeNode) BSTIterator {
    var res []int
    var stack []*TreeNode
    for root != nil || len(stack) > 0 {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        root = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        res = append(res, root.Val)
        root = root.Right
    }
    return BSTIterator{Val:res}
}


func (this *BSTIterator) Next() int {
    val := this.Val[0]
    this.Val = this.Val[1:]
    return val
}


func (this *BSTIterator) HasNext() bool {
    return len(this.Val) > 0
}


