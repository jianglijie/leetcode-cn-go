package main

func longestPalindrome(s string) string {
	l := len(s)
	if l < 2 {
		return s
	}
	dp := make([][]bool, l)
    for i := 0; i < len(dp); i++ {
        dp[i] = make([]bool, l)
    }
	max, start := 1, 0
	for j := 1; j < l; j++ {
		for i := 0; i < j; i++ {
			if s[i] == s[j] {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			} else {
				dp[i][j] = false
			}
			if dp[i][j] {
				tmp := j - i + 1
				if tmp > max {
					max = tmp
					start = i
				}
			}
		}
	}
	return s[start : start+max]
}

func main() {

}
