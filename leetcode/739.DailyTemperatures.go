package main

import "fmt"

func dailyTemperatures(T []int) []int {
	ret := make([]int, len(T))
	s := make([]int, len(T))
	bigger := -1
	for i := len(T) - 1; i >= 0; i-- {
		for {
			if bigger < 0 || T[i] < T[s[bigger]] {
				break
			}
			bigger--
		}
		bigger++
		s[bigger] = i
		ret[i] = 0
		if bigger == 0 {
			ret[i] = 0
		} else {
			ret[i] = s[bigger-1] - i
		}
	}
	return ret
}

func main() {
	n := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(n))
}
