package main

import "fmt"

func removeDuplicates(S string) string {
    var s string
    for _, i := range S{
        if s == ""{
            s = string(i)
        }else if i == rune(s[len(s)-1]) {
			s = s[:len(s)-1]
		}else{
			s += string(i)
		}
    }
    return s
}


func main() {
	fmt.Println(removeDuplicates("abbaca"))
}
