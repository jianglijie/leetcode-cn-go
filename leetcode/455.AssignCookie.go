package main

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	res := 0
	for i := 0; res < len(g) && i < len(s); i++ {
		if g[res] <= s[i] {
			res++
		}
	}
	return res
}

func main() {

}
