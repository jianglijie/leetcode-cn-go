package main

//官方
import (
	"fmt"
	"strconv"
	"unicode"
)

func decodeAtIndex(S string, K int) string {
	size, length := 0, len(S)
	for _, i := range S{
		if unicode.IsDigit(i) {
			num, _ := strconv.Atoi(string(i))
			size *= num
		} else {
			size++
		}
	}
	for i := length - 1; i >= 0; i-- {
		K %= size
		if K == 0 && unicode.IsLetter(rune(S[i])) {
			return string(S[i])
		}
		if unicode.IsDigit(rune(S[i])) {
			num, _ := strconv.Atoi(string(S[i]))
			size /= num
		} else {
			size--
		}
	}
	return ""
}

func main() {
	fmt.Println(decodeAtIndex("leet2code3", 10))
}
