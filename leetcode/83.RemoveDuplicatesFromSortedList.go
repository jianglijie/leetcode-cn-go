package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    p := head
    for p != nil && p.Next != nil{
        if p.Val == p.Next.Val{
            p.Next = p.Next.Next
        }else{
            p = p.Next
        }
    }
    return head
}

func main() {
	
}
