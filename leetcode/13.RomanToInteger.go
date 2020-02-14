package main

func romanToInt(s string) int {
	romanDict := map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	res := 0
	tmp := []rune(s)
	for i := 0; i < len(tmp)-1; i++ {
		if romanDict[tmp[i]] < romanDict[tmp[i+1]] {
			res -= romanDict[tmp[i]]
		} else {
			res += romanDict[tmp[i]]
		}
	}
	return res + romanDict[tmp[len(s)-1]]
}

func main() {
}
