package main

import (
	"strconv"
	"strings"
)

func dayOfYear(date string) int {
	arr := strings.Split(date, "-")
	y, _ := strconv.Atoi(arr[0])
	m, _ := strconv.Atoi(arr[1])
	d, _ := strconv.Atoi(arr[2])
	days := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if (y % 100 != 0 && y % 4 == 0) || y % 400 == 0{
		days[1] = 29
	}
	res := 0
	if m == 1{
		return d
	}else{
		for i :=0; i<m-1;i++{
			res += days[i]
		}
	}
	res += d
	return res
}

func main() {
	
}
