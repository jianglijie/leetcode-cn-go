package main

import (
	"fmt"
	"strings"
)

func canConstruct(ransomNote string, magazine string) bool {
	runeR, runeM := []rune(ransomNote), []rune(magazine)
	if len(runeR) > len(runeM) {
		return false
	}
	caps := make([]int, 26)
	for _, c := range runeR {
		index := strings.IndexRune(string(runeM[caps[c-97]:]), c)
		if index == -1 {
			return false
		}
		caps[c-97] += index + 1
	}
	return true
}

func main() {
	fmt.Println(canConstruct("fihjjjjei", "hjibagacbhadfaefdjaeaebgi"))
}
