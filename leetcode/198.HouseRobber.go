package main

func rob(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	dp := make([]int, length+1)
	dp[0] = 0
	dp[1] = nums[0]
	for i := 2; i <= length; i++ {
		if dp[i-1] > dp[i-2]+nums[i-1] {
			dp[i] = dp[i-1]
		} else {
			dp[i] = dp[i-2] + nums[i-1]
		}
	}
	return dp[length]
}

func main() {

}
