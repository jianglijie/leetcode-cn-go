package main

func firstUniqChar(s string) int {
	charDict := make(map[rune]uint)
	runeS := []rune(s)
	for i := 0; i < len(runeS); i++ {
		charDict[runeS[i]-97]++
	}
	for i := 0; i < len(runeS); i++ {
		if charDict[runeS[i]-97] == 1 {
			return i
		}
	}
	return -1
}

func main() {

}
