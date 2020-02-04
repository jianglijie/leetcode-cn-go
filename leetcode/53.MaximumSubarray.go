package main

func maxSubArray(nums []int) int {
	res, sum := nums[0], 0
	for _, i := range nums{
		if sum > 0{
			sum += i
		}else{
			sum = i
		}
		if sum > res{
			res = sum
		}
	}
	return res
}

func main() {
	
}
