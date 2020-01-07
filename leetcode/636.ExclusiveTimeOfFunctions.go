package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Func struct {
	Other int
	Start int
}

func exclusiveTime(n int, logs []string) []int {
	var s []Func
	r := make([]int, n)
	for _, i := range logs{
		arr := strings.Split(i, ":")
		num, _ := strconv.Atoi(arr[0])
		sec, _ := strconv.Atoi(arr[2])
		if arr[1] == "start"{
			f := Func{Other:0, Start:sec}
			s = append(s, f)
		}else {
			execSec := sec - s[len(s)-1].Start + 1
			r[num] += execSec - s[len(s)-1].Other
			if len(s) > 1{
				s[len(s)-2].Other += execSec
			}
			s = s[:len(s)-1]
		}
	}
	return r
}

func main() {
	l := []string{"0:start:0", "1:start:2", "1:end:5", "0:end:6"}
	fmt.Println(exclusiveTime(2, l))
}
