package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//todo
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func clone(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	newRoot := &TreeNode{Val: root.Val}
	newRoot.Left = clone(root.Left)
	newRoot.Right = clone(root.Right)
	return newRoot
}

func generateTrees(n int) []*TreeNode {
	var res  []*TreeNode
	if n == 0{
		return res
	}
	res = append(res, nil)
    for i:=1;i<=n;i++{
    	var cur  []*TreeNode
    	//遍历之前的所有解
    	for _, item := range res{
    		//插入到根节点
    		node  := &TreeNode{Val:i}
    		node.Left = item
			cur = append(cur, node)
			//插入到右孩子，右孩子的右孩子...最多找 n 次孩子
			for j:=0;j<=n;j++{
				copy := clone(item)//复制当前的树
				right := copy//找到要插入右孩子的位置
				for k:=0;k<j;k++{//遍历 j 次找右孩子
					if right == nil{
						break
					}
					right = right.Right
				}
				if right == nil{//到达 null 提前结束
					break
				}
				//保存当前右孩子的位置的子树作为插入节点的左孩子
				rightTree := right.Right
				node1 := &TreeNode{Val:i}
				right.Right = node1//右孩子是插入的节点
				node1.Left = rightTree//插入节点的左孩子更新为插入位置之前的子树
				cur = append(cur, copy)
			}
		}
    	res = cur
	}
	return res
}

func main() {

}
