package main

func majorityElement(nums []int) int {
    value, count := 0, 0
    for _, i := range nums{
        if count == 0{
            value = i
        }
        if i == value{
            count += 1
        }else{
            count -= 1
        }
    }
    return value
}

func main() {
	
}
