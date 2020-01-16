package main

import "fmt"

func longestWPI(hours []int) int {
	var f, s, decList []int
	f = append(f, 0)
	for _, i := range hours {
		if i > 8 {
			f = append(f, 1)
		} else {
			f = append(f, -1)
		}
	}
	var tmp = 0
	for i := 0; i < len(f); i++ {
		tmp += f[i]
		s = append(s, tmp)
		if len(decList) == 0 || s[decList[len(decList)-1]] > tmp {
			decList = append(decList, i)
		}
	}
	fmt.Println(f, s, decList)
	n, res, dsize := len(f)-1, 0, len(decList)-1
	for n >= 0 {
		for dsize >= 0 && s[decList[dsize]] < s[n] {
			if res < n-decList[dsize] {
				res = n - decList[dsize]
			}
			dsize -= 1
		}
		n -= 1
	}
	return res
}

func main() {
	l := []int{9,9,6,0,6,6,9}
	longestWPI(l)
}
