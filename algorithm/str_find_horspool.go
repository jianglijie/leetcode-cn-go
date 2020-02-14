package main

import "fmt"

func shift_table(pattern []rune) []int {
	pLen := len(pattern)
	table := make([]int, 128)
	for i := 0; i < len(table); i++ {
		table[i] = pLen
	}
	for i := 0; i <= pLen-1; i++ {
		table[pattern[i]] = pLen - 1 - i
	}
	return table
}

// Horspool
// 参数：
//	p：string 查找内容
//	t: string 查找范围
// 返回值：
//  index: int 索引值，未找到-1
func Horspool(p string, t string) int {
	pLen := len(p)
	tLen := len(t)
	if pLen > tLen {
		return -1
	}
	tmpP, tmpT := []rune(p), []rune(t)
	table := shift_table(tmpP)
	index := pLen - 1
	for index <= tLen-1 {
		match := 0
		for match < pLen && tmpP[pLen-1-match] == tmpT[index-match] {
			match++
		}
		if match == pLen {
			return index - match + 1
		} else {
			index += table[tmpT[index]]
		}
	}
	return -1
}

func main() {
	fmt.Println(Horspool("abc", "fsdfsdabcsdfs"))
}
