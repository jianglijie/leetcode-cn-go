package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    s := &ListNode{}
    s.Next = head
    pre, cur := s, head
    for cur != nil{
        if cur.Val == val{
            pre.Next = cur.Next
        }else{
            pre = cur
        }
        cur = cur.Next
    }
    return s.Next
}

func main() {
	
}
