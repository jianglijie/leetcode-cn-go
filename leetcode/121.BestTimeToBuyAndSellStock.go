package main

import "fmt"

func maxProfit(prices []int) int {
	max, min := 0, int(^uint(0)>>1) //max_int
	for _, i := range prices {
		if i < min {
			min = i
		} else if i-min > max {
			max = i - min
		}
	}
	return max
}

func main() {
	l := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(l))
}
