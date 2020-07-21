package main

import "sort"

func largestPerimeter(A []int) int {
	if len(A) < 3 {
		return 0
	}
	sort.Ints(A)
	for i := len(A) - 1; i >= 2;i--{
		if A[i-1] + A[i-2] > A[i]{
			return A[i-1] + A[i-2] + A[i]
		}
	}
	return 0
}


func main() {
	
}
