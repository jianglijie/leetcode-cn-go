package main

func missingNumber(nums []int) int {
    l := len(nums)
    for i, v := range nums{
        l = l ^ i ^ v
    }
    return l
}

func main() {
	
}
