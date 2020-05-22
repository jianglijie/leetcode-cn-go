package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    arr := []int{}
    if head == nil || head.Next == nil {
        return true
    }
    for head != nil {
        arr = append(arr,head.Val)
        head = head.Next
    }
    left, right := 0, len(arr)-1
    for left < right {
        if arr[left] == arr[right]{
            left++
            right--
        }else{
            return false
        }
    }
    return true
}

func main() {
	
}
