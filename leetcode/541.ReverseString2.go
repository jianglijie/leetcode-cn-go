package main

import "fmt"

func reverseStr(s string, k int) string {
	runeS := []rune(s)
    l := len(runeS)
    left:=0
    for i := 0 ; left < l; i++ {
        left = i*2*k
        right := left + k - 1
        if right > l - 1 {
            right = l - 1
        }
        for left < right {
            runeS[left],runeS[right] = runeS[right],runeS[left]
            left +=1
            right -=1
        }
    }
    return string(runeS)
}

func main() {
	fmt.Println(reverseStr("abcdefg", 2))
}
