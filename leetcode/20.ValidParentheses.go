package main


func isValid(s string) bool {
    if s == ""{
        return true
    }
    if len(s) % 2 == 1 {
        return false
    }
    temStr := ""
    myMap := map[byte]byte{
        '(': ')',
        '[': ']',
        '{': '}',
    }
    for i, _ := range s {
        if temStr == "" {
            temStr = temStr + s[i:i + 1]
            continue
        }
        if _, ok := myMap[s[i]]; !ok {
            if myMap[temStr[len(temStr) - 1]] != s[i] {
                return false
            }
            temStr = temStr[:len(temStr) - 1]
        } else {
            temStr = temStr + s[i:i + 1]
        }
    }
    return temStr == ""
}
