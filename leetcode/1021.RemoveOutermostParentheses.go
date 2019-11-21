package main

func removeOuterParentheses(S string) string {
    var s string
	flag := false
	n := 0
	for _, i := range S {
		if i == '(' && !flag {
			flag = true
			n++
			continue
		}
		if flag {
			if i == '(' {
				n++
				s += string(i)
			}
			if i == ')' {
				n--
				if n == 0 {
					flag = false
					continue
				}
				s += string(i)
			}
		}
	}
	return s
}
