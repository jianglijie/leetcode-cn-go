package main

import "fmt"

func find132pattern(nums []int) bool {
	length := len(nums)
	if length < 3 {
		return false
	}
	max := int(^uint(0) >> 1) //max_int
	min := ^max               //min_int
	s := []int{nums[length-1]}
	for i := length - 2; i >= 0; i-- {
		if nums[i] < min {
			return true
		} else {
			for {
				if len(s) == 0 || nums[i] <= s[len(s)-1] {
					break
				}
				if s[len(s)-1] > min {
					min = s[len(s)-1]
				}
				s = s[:len(s)-1]
			}
			s = append(s, nums[i])
		}
	}
	return false
}

func main() {
	fmt.Println(find132pattern([]int{3, 1, 4, 2}))
}
