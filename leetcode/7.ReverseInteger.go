package main

import (
	"fmt"
	"strconv"
)

func reverse(x int) int {
	if x == 0 {
        return 0
    }
	strNum := strconv.Itoa(x)
	ret := ""
    pos := len(strNum) - 1

    flag := 1
    for pos >= 0 {
        if pos == 0 && strNum[0] == '-' {
            flag = -1
        } else {
            ret = fmt.Sprintf("%s%s", ret, string(strNum[pos]))
        }
        pos -= 1
    }
    val, _ := strconv.Atoi(ret)
	if val > 1<<31 - 1 {
		return 0
	}
    return val * flag
}

func main() {
	fmt.Println(reverse(123900))
}
