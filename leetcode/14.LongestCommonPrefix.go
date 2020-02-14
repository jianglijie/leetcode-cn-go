package main

func longestCommonPrefix(strs []string) string {
	l := len(strs)
	if l == 0 {
		return ""
	}
	res := []rune(strs[0])
	for i := 1; i < len(strs); i++ {
		j := 0
		for ; j < len(res) && j < len(strs[i]); j++ {
			tmp := []rune(strs[i])
			if res[j] != tmp[j] {
				break
			}
		}
		res = res[:j]
		if len(res) == 0 {
			return ""
		}
	}
	return string(res)
}

func main() {

}
