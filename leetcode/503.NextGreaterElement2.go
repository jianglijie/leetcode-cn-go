package main

import "fmt"

func nextGreaterElements(nums []int) []int {
	length := len(nums)
	res := make([]int, length)
	var s []int
	for i := 2*length - 1; i >= 0; i-- {
		for {
			if len(s) == 0 || nums[s[len(s)-1]] > nums[i%length] {
				break
			}
			s = s[:len(s)-1]
		}
		if len(s) == 0 {
			res[i%length] = -1
		} else {
			res[i%length] = nums[s[len(s)-1]]
		}
		s = append(s, i%length)
	}
	return res
}

func main() {
	fmt.Println(nextGreaterElements([]int{3, 8, 4, 1, 2}))
}
