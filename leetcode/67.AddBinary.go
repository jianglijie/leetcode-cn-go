package main

import "fmt"

func reverse(s []rune) []rune {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func addBinary(a string, b string) string {
	runeA, runeB, runeRes := []rune(a), []rune(b), []rune("")
	carry, sum := 0, 0
	for i,j:=len(runeA)-1, len(runeB)-1;i>=0||j>=0;i,j=i-1,j-1{
		sum = carry
		if i>=0{
			sum += int(runeA[i]-'0')
		}
		if j>=0{
			sum += int(runeB[j]-'0')
		}
		fmt.Println(sum)
		runeRes = append(runeRes, rune(sum % 2 + '0'))
		carry = sum/2
	}
	if carry == 1{
		runeRes = append(runeRes, '1')
	}
	runeRes = reverse(runeRes)
    return string(runeRes)
}

func main() {
	fmt.Println(addBinary("11", "1"))
}
