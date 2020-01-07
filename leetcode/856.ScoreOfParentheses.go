package main

import (
	"fmt"
)

func scoreOfParentheses(S string) int {
	left, right, ret := 0, 0, 0
	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			left++
		} else {
			right++
			if left == right || S[i+1] == '(' {
				ret += 1 << uint(left-1)
				left -= right
				right = 0
			}
		}
	}
	return ret
}

func main() {
	fmt.Println(scoreOfParentheses("()(())"))
}
