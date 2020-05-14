package main

func checkRecord(s string) bool {
    runeS := []rune(s)
    c := 0
    for i:=0;i<len(runeS);i++{
    	if runeS[i] == 'A'{
    		c++
    		if c > 1{
				return false
			}
		}else if runeS[i] == 'L'{
			if i<len(runeS)-2 && runeS[i+1] == 'L' && runeS[i+2] == 'L'{
				return false
			}
		}
	}
    return true
}

func main() {
	
}
