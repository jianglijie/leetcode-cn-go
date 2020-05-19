package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	n1 := l1
	n2 := l2
	l3 := new(ListNode)
	n3 := l3
	carry := 0
	for n1 != nil || n2 != nil || carry > 0 {
		n3.Next = new(ListNode)
		n3 = n3.Next
		x := 0
		y := 0
		if n1 != nil {
			x = n1.Val
		}
		if n2 != nil {
			y = n2.Val
		}
        t := x + y + carry
        n3.Val = t % 10
        carry = t / 10
		if n1 != nil {
			n1 = n1.Next
		}
		if n2 != nil {
			n2 = n2.Next
		}
	}
	return l3.Next
}

func main() {
	
}
