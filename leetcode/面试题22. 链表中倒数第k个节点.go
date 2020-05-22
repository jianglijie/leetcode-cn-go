package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getKthFromEnd(head *ListNode, k int) *ListNode {
    count, cur := 0, head
    for cur != nil{
        count++
        cur = cur.Next
    }
    cur = head
    for count > k {
        cur = cur.Next
        count--
    }
    return cur
}

func main() {
	
}
