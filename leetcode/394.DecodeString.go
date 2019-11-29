package main

// 借用
import (
	"bytes"
	"fmt"
    "strconv"
    "unicode"
)

type Pair struct {
    nextNum int
    thisStr string
}

func multiplyStr(n int, s string) string {
    var buffer bytes.Buffer
    for i := 0; i < n; i ++ {
        buffer.WriteString(s)
    }
    return buffer.String()
}

func decodeString(s string) string {
    num := 0
    curStr := ""
    stack := []Pair{}

    for _, c := range s {
        if unicode.IsDigit(c) { // end of string
            n, _ := strconv.Atoi(string(c))
            num = num * 10 + n
        } else if c  == '[' { // end of number
            stack = append(stack, Pair{num, curStr})
            num, curStr = 0, ""
        } else if c == ']' {
            pair := stack[len(stack) - 1]
            stack = stack[:len(stack) - 1]
            curStr = pair.thisStr + multiplyStr(pair.nextNum, curStr)
        } else {
            curStr = curStr + string(c)
        }
    }

    return curStr
}


func main()  {
	fmt.Println(decodeString("3[a2[c]]"))
}