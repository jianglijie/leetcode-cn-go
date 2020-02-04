package main

import "fmt"

func maxProfit(prices []int) int {
	res := 0
	for i := 0; i < len(prices)-2; i++ {
		if prices[i+1] > prices[i] {
			res += prices[i+1] - prices[i]
		}
	}
	return res
}

func main() {
	l := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(l))
}
