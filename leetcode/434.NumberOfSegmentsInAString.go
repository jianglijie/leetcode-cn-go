package main

import "fmt"

func countSegments(s string) int {
	if s == ""{
		return 0
	}
    runsS, count, f := []rune(s), 0, false
    for i:=0;i<len(runsS);i++{
    	if runsS[i] == ' ' && f{
    		count++
    		f = false
		}else if runsS[i] != ' '{
			f = true
		}
	}
    if runsS[len(runsS)-1] != ' '{
    	count++
	}
    return count
}

func main() {
	fmt.Println(countSegments("                "))
}
