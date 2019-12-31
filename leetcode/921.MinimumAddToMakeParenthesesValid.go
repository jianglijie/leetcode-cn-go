package main

import "fmt"

func minAddToMakeValid(S string) int {
	left := 0
	right := 0
	for _, i := range S {
		if i == '(' {
			left++
		} else if i == ')' {
			if left > 0 {
				left--
			} else {
				right++
			}
		}
	}
	return left + right
}

func main() {
	fmt.Println(minAddToMakeValid(""))
}
