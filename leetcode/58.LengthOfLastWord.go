package main

import "fmt"

func lengthOfLastWord(s string) int {
	tmp := []rune(s)
	l := 0
	for i:=len(tmp)-1;i>=0;i--{
		if tmp[i] == ' '{
			if l == 0{
				continue
			}else{
				break
			}
		}else{
			l++
		}
	}
	return l
}

func main() {
	fmt.Println(lengthOfLastWord("sdf dddd   "))
}
