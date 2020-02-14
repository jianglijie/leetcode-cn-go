package main

import "fmt"

func getNext(p string) []int {
	l := len(p)
	next := make([]int, l)
	next[0] = -1
	next[1] = 0
	i := 0
	j := 1
	for j < l-1 {
		if i == -1 || p[i] == p[j] {
			i++
			j++
			next[j] = i
		} else {
			i = next[i]
		}
	}
	return next
}

// KMP算法
// 参数：
//	p：string 查找内容
//	t: string 查找范围
// 返回值：
//  index: int 索引值，未找到-1
func KMP(p string, t string) int {
	i, j := 0, 0
	pLen := len(p)
	tLen := len(t)
	if pLen > tLen {
		return -1
	}
	next := getNext(p)
	for i < tLen && j < pLen {
		if j == -1 || t[i] == p[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == pLen {
		return i - j
	} else {
		return -1
	}
}

func main() {
	fmt.Println(KMP("abc", "fsdfsdabcsdfs"))
}
