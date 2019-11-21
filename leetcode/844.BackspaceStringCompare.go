package main

func backspaceCompare(S string, T string) bool {
    var s1 string
    var s2 string
    for _, i := range S{
        if i == '#' && len(s1) > 0{
            s1 = s1[:len(s1)-1]
        }else if i != '#'{
            s1 += string(i)
        }
    }
    for _, i := range T{
        if i == '#' && len(s2) > 0{
            s2 = s2[:len(s2)-1]
        }else if i != '#'{
            s2 += string(i)
        }
    }
    return s1 == s2
}


func main() {
	backspaceCompare("y#fo##f", "y#f#o##f")
}
