package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	strNum := strconv.Itoa(x)
	ret := ""
	for _, v := range strNum {
		ret = string(v) + ret
	}
	if ret == strNum{
		return true
	}else{
		return false
	}
}

func main() {
	fmt.Println(isPalindrome(12321))
}
