package main

func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}
	sPtr, tPtr := 0, 0
	for sPtr < len(s) && tPtr < len(t) {
		if s[sPtr] == t[tPtr] {
			sPtr++
		}
		tPtr++
	}
	return sPtr == len(s)
}

func main() {

}
