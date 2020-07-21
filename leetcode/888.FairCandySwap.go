package main

func fairCandySwap(A []int, B []int) []int {
	sa, sb := 0, 0
	var res []int
	for _, i := range A{
		sa += i
	}
	mapB := make(map[int]byte)
	for _, i := range B{
		sb += i
		mapB[i] = '1'
	}
	target := (sb-sa)/2
	for _, i := range A{
		if _, ok := mapB[i + target];ok{
			res = append(res, i)
			res = append(res, i+target)
			break
		}
	}
	return res
}

func main() {
	
}
