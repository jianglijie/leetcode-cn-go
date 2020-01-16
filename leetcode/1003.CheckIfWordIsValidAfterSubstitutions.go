package main

import "fmt"

func isValid(S string) bool {
	var s []rune
	for _, i := range S{
		if i == 'a' || i == 'b'{
			 s = append(s, i)
		}else if i == 'c'{
			if len(s) < 2{
				return false
			}
			if s[len(s)-1] != 'b' ||  s[len(s)-2] != 'a'{
				return false
			}
			s = s[:len(s)-2]
		}
	}
	if len(s) == 0{
		return true
	}else{
		return false
	}
}

func main() {
	fmt.Println(isValid("aabcbc"))
}
