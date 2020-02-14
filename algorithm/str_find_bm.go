package main

import "fmt"

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func getBad(p []rune) []int {
	pLen := len(p)
	bad := make([]int, 128)
	for i := 0; i < len(bad); i++ {
		bad[i] = pLen
	}
	for i := 0; i < pLen-1; i++ {
		k := p[i]
		bad[k] = pLen - 1 - i
	}
	return bad
}

func getGood(p []rune) []int {
	pLen := len(p)
	good := make([]int, pLen)
	last := pLen
	for i := pLen - 1; i >= 0; i-- {
		if isPrefix(p, i+1) {
			last = i + 1
		}
		good[pLen-1-i] = last - 1 + pLen - 1
	}
	for i := 0; i < pLen-1; i++ {
		k := suffixLength(p, i)
		good[k] = pLen - 1 - i + k
	}
	return good
}

// 前缀
func isPrefix(pattern []rune, p int) bool {
	pLen := len(pattern)
	for i, j := p, 0; i < pLen; i, j = i+1, j+1 {
		if pattern[i] != pattern[j] {
			return false
		}
	}
	return true
}

//后缀
func suffixLength(pattern []rune, p int) int {
	pLen, len := len(pattern), 0
	for i, j := p, pLen-1; i >= 0 && pattern[i] == pattern[j]; i, j = i-1, j-1 {
		len += 1
	}
	return len
}

// BM算法
// 参数：
//	p：string 查找内容
//	t: string 查找范围
// 返回值：
//  index: int 索引值，未找到-1
func BM(p string, t string) int {
	pLen := len(p)
	tLen := len(t)
	if pLen > tLen {
		return -1
	}
	tmpP, tmpT := []rune(p), []rune(t)
	bad := getBad(tmpP)
	good := getGood(tmpP)
	for i, j := pLen-1, 0; i < tLen; {
		for j = pLen - 1; tmpT[i] == tmpP[j]; i, j = i-1, j-1 {
			if j == 0 {
				return i
			}
		}
		i += max(good[pLen-j-1], bad[t[i]])
	}
	return -1
}

func main() {
	fmt.Println(BM("abc", "fsdfsdabcsdfs"))
}
