package main

import "fmt"

func min(x int, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func minCostClimbingStairs(cost []int) int {
	length := len(cost)
	dp := [3]int{0, 0, 0}
	dp[0], dp[1] = cost[0], cost[1]
	for i := 2; i < length; i++ {
		dp[i%3] = min(dp[(i-1)%3], dp[(i-2)%3]) + cost[i]
	}
	return min(dp[(length-1)%3], dp[(length-2)%3])
}

func main() {
	l := []int{10, 15, 20}
	fmt.Println(minCostClimbingStairs(l))
}
