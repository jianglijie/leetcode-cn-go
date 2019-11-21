package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	var s []int
	var tmp int
	for _, i := range tokens{
		switch i {
		case "+":
			tmp = s[len(s)-2] + s[len(s)-1]
			s = s[:len(s)-2]
			s = append(s, tmp)
		case "-":
			tmp = s[len(s)-2] - s[len(s)-1]
			s = s[:len(s)-2]
			s = append(s, tmp)
		case "*":
			tmp = s[len(s)-2] * s[len(s)-1]
			s = s[:len(s)-2]
			s = append(s, tmp)
		case "/":
			tmp = s[len(s)-2] / s[len(s)-1]
			s = s[:len(s)-2]
			s = append(s, tmp)
		default:
			tmp, _ = strconv.Atoi(i)
			s = append(s, tmp)
		}
	}
	return s[0]
}

func main()  {
	tokens := []string{"0", "3", "/"}
	fmt.Println(evalRPN(tokens))
}

