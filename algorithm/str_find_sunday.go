package main

import "fmt"

// Sunday算法
// 参数：
//	p：string 查找内容
//	t: string 查找范围
// 返回值：
//  index: int 索引值，未找到-1
func Sunday(p string, t string) int {
	pLen := len(p)
	tLen := len(t)
	if pLen > tLen {
		return -1
	}
	if pLen == 0{
        return 0
    }
	tmpP, tmpT := []rune(p), []rune(t)
	table := make([]int, 128)
	for i := 0; i < 128; i++ {
		table[i] = pLen + 1
	}
	for i := 0; i < pLen; i++ {
		table[tmpP[i]] = pLen - i
	}
	start, l := 0, 0
	for start <= tLen-pLen {
		l = 0
		for tmpT[start+l] == tmpP[l] {
			l++
			if l >= pLen {
				return start
			}
			if start+l == tLen {
				return -1
			}
		}
		if start+pLen == tLen{
			return -1
		}
		start += table[tmpT[start+pLen]]
	}
	return -1
}

func main() {
	fmt.Println(Sunday("abc", "fsdfsdabcsdfs"))
}
