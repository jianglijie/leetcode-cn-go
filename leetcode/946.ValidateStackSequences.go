package main

import "fmt"

func validateStackSequences(pushed []int, popped []int) bool {
    index, length := 0, len(popped)
    var res []int
    for _, i := range pushed{
    	res = append(res, i)
    	for{
    		if len(res) == 0 || index >= length || res[len(res)-1] != popped[index]{
    			break
			}
    		index++
    		res = res[:len(res)-1]
		}
	}
    return index == length
}

func main() {
	p1 := []int{1, 2, 3, 4, 5}
	p2 := []int{4, 5, 3, 2, 1}
	fmt.Println(validateStackSequences(p1, p2))
}
