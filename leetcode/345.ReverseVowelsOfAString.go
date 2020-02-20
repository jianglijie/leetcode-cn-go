package main

import "fmt"

func isVowels(c rune) bool {
	if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' || c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U' {
		return true
	} else {
		return false
	}
}

func reverseVowels(s string) string {
	runeS := []rune(s)
	i,j := 0, len(runeS)-1
	for i<j{
		if !isVowels(runeS[i]){
			i++
			continue
		}
		if !isVowels(runeS[j]){
			j--
			continue
		}
		runeS[i], runeS[j] = runeS[j], runeS[i]
		i++
		j--
	}
	return string(runeS)
}

func main() {
	fmt.Println(reverseVowels("hello"))
}
