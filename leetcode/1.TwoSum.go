package main

import "fmt"

func twoSum(nums []int, target int) (ret []int) {
	numMap := make(map[int]int)
    for k, v := range nums{
    	if _, ok := numMap[target-v]; ok {
        	ret = append(ret, numMap[target-v])
			ret = append(ret, k)
            break
        }
        numMap[v] = k
    }
    return
}

func main() {
	nums :=[] int {2,7,11, 15}
	fmt.Println(twoSum(nums, 9))
}
