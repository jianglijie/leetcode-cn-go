package main

import "fmt"

func isPalindrome(s string) bool {
	runeS, runeRes := []rune(s), []rune("")
	for i := 0; i <= len(runeS)-1; i++ {
		if (runeS[i] >= 97 && runeS[i] <= 122) || (runeS[i] >= 48 && runeS[i] <= 57) {
			runeRes = append(runeRes, runeS[i])
		} else if runeS[i] >= 65 && runeS[i] <= 90 {
			runeRes = append(runeRes, runeS[i]+32)
		}
	}
	for i, j := 0, len(runeRes)-1; i <= j; i, j = i+1, j-1 {
		if runeRes[i] != runeRes[j] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("0P"))
}
