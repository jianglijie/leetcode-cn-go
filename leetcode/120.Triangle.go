package main

func min(x int, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func minimumTotal(triangle [][]int) int {
	l := len(triangle)
	res := make([]int, l)
	res = triangle[l-1]
	for i := l - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			res[j] = min(res[j], res[j+1]) + triangle[i][j]
		}
	}
	return res[0]
}

func main() {

}
