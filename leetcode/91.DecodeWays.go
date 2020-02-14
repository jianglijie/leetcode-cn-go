package main

import "fmt"

func numDecodings(s string) int {
	l := len(s)
	dp := make([]int, l+1)
	dp[l%3] = 1
	if s[l-1] != '0' {
		dp[(l-1)%3] = 1
	}
	for i := l - 2; i >= 0; i-- {
		if s[i] == '0' {
			dp[i%3] = 0
			continue
		}
		tmp1, tmp2 := dp[(i+1)%3], 0
		x, y := int(s[i]-'0')*10, int(s[i+1]-'0')
		if x+y < 26 {
			tmp2 = dp[(i+2)%3]
		}
		dp[i%3] = tmp1 + tmp2
	}
	return dp[0]
}

func main() {
	fmt.Println(numDecodings("12"))
}
