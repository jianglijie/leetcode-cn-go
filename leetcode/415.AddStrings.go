package main

import "fmt"

func reverse(s []rune) []rune {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func addStrings(num1 string, num2 string) string {
	rune1, rune2, runeRes := []rune(num1), []rune(num2), []rune("")
	carry, sum := 0, 0
	for i,j:=len(rune1)-1, len(rune2)-1;i>=0||j>=0;i,j=i-1,j-1{
		sum = carry
		if i>=0{
			sum += int(rune1[i]-'0')
		}
		if j>=0{
			sum += int(rune2[j]-'0')
		}
		runeRes = append(runeRes, rune(sum % 10 + '0'))
		carry = sum/10
	}
	if carry > 0{
		runeRes = append(runeRes, rune(carry + '0'))
	}
	runeRes = reverse(runeRes)
    return string(runeRes)
}

func main() {
	fmt.Println(addBinary("11", "1"))
}
