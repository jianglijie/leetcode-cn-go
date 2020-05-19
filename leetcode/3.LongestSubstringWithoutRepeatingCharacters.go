package main

func lengthOfLongestSubstring(s string) int {
    tmp := map[byte]int {}
    max := 0
    l := len(s)
    for i, j := 0, 0; i < l; i ++ {
        tmp[s[i]] ++
        for tmp[s[i]] > 1 {
            tmp[s[j]] --
            j ++
        }
        if max < i - j + 1{
            max = i - j + 1
        }
    }
    return max
}

func main() {
	
}
