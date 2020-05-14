package main

func singleNumber(nums []int) int {
    ret := 0
    for _, i := range nums{
        ret ^= i
    }
    return ret
}

func main() {
	
}
