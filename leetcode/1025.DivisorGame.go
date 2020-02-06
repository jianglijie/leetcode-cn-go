package main

func divisorGame(N int) bool {
    // return N % 2 == 0
    if N <= 1{
        return false
    }
    dp := make([]bool, N+1)
    dp[1] = false
    dp[2] = true
    for i:=3;i<=N;i++{
    	for j:=1;j<i;j++{
    		if i%j==0 && !dp[i-j]{
    			dp[i] = true
			}
		}
	}
    return dp[N]
}

func main() {
	
}
